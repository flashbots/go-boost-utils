GIT_VER := $(shell git describe --tags --always --dirty="-dev")
# ECR_URI := 223847889945.dkr.ecr.us-east-2.amazonaws.com/your-project-name

all: clean build

v:
	@echo "Version: ${GIT_VER}"

clean:
	rm -rf your-project build/

build:
	go build -ldflags "-X main.version=${GIT_VER}" -v -o your-project main.go

test:
	go test ./...

lint:
	gofmt -d ./
	go vet ./...
	staticcheck ./...

generate-ssz:
	rm -f types/builder_encoding.go types/signing_encoding.go types/common_encoding.go
	sszgen --path types --include ../go-ethereum/common/hexutil --objs Eth1Data,BeaconBlockHeader,SignedBeaconBlockHeader,ProposerSlashing,Checkpoint,AttestationData,IndexedAttestation,AttesterSlashing,Attestation,Deposit,VoluntaryExit,SyncAggregate,ExecutionPayloadHeader,VersionedExecutionPayloadHeader,BlindedBeaconBlockBody,BlindedBeaconBlock,RegisterValidatorRequestMessage,BuilderBid,SignedBuilderBid,SigningData,forkData,Transactions

cover:
	go test -coverprofile=/tmp/go-sim-lb.cover.tmp ./...
	go tool cover -func /tmp/go-sim-lb.cover.tmp
	unlink /tmp/go-sim-lb.cover.tmp

cover-html:
	go test -coverprofile=/tmp/go-sim-lb.cover.tmp ./...
	go tool cover -html=/tmp/go-sim-lb.cover.tmp
	unlink /tmp/go-sim-lb.cover.tmp

build-for-docker:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=${GIT_VER}" -v -o your-project main.go

docker-image:
	DOCKER_BUILDKIT=1 docker build . -t your-project
	docker tag your-project:latest ${ECR_URI}:${GIT_VER}
	docker tag your-project:latest ${ECR_URI}:latest

docker-push:
	docker push ${ECR_URI}:${GIT_VER}
	docker push ${ECR_URI}:latest

k8s-deploy:
	@echo "Checking if Docker image ${ECR_URI}:${GIT_VER} exists..."
	@docker manifest inspect ${ECR_URI}:${GIT_VER} > /dev/null || (echo "Docker image not found" && exit 1)
	kubectl set image deploy/deployment-your-project app-your-project=${ECR_URI}:${GIT_VER}
	kubectl rollout status deploy/deployment-your-project
