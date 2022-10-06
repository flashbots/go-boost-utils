package types

// PayloadStatusV1 https://github.com/ethereum/execution-apis/blob/main/src/engine/specification.md#payloadstatusv1
type PayloadStatusV1 struct {
	Status          string  `json:"status"`
	LatestValidHash *Hash   `json:"latestValidHash"`
	ValidationError *string `json:"validationError"`
}

// ForkChoiceResponse https://github.com/ethereum/execution-apis/blob/main/src/engine/specification.md#response-1
type ForkChoiceResponse struct {
	PayloadStatus PayloadStatusV1 `json:"payloadStatus"`
	PayloadID     *PayloadID      `json:"payloadId"`
}
