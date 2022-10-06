package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestExecutionPayloadHeader(t *testing.T) {
	baseFeePerGas := U256Str{}
	baseFeePerGas[0] = 0x08

	h := ExecutionPayloadHeader{
		ParentHash:       Hash{0x01},
		FeeRecipient:     Address{0x02},
		StateRoot:        Root{0x03},
		ReceiptsRoot:     Root{0x04},
		LogsBloom:        Bloom{0x05},
		Random:           Hash{0x06},
		BlockNumber:      5001,
		GasLimit:         5002,
		GasUsed:          5003,
		Timestamp:        5004,
		ExtraData:        []byte{0x07},
		BaseFeePerGas:    baseFeePerGas,
		BlockHash:        Hash{0x09},
		TransactionsRoot: Root{0x0a},
	}
	b, err := json.Marshal(h)
	require.NoError(t, err)

	expectedJSON := `{
        "parent_hash": "0x0100000000000000000000000000000000000000000000000000000000000000",
        "fee_recipient": "0x0200000000000000000000000000000000000000",
        "state_root": "0x0300000000000000000000000000000000000000000000000000000000000000",
        "receipts_root": "0x0400000000000000000000000000000000000000000000000000000000000000",
        "logs_bloom": "0x05000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
        "prev_randao": "0x0600000000000000000000000000000000000000000000000000000000000000",
        "block_number": "5001",
        "gas_limit": "5002",
        "gas_used": "5003",
        "timestamp": "5004",
        "extra_data": "0x07",
        "base_fee_per_gas": "8",
        "block_hash": "0x0900000000000000000000000000000000000000000000000000000000000000",
        "transactions_root": "0x0a00000000000000000000000000000000000000000000000000000000000000"
    }`
	require.JSONEq(t, expectedJSON, string(b))

	// Now unmarshal it back and compare to original
	h2 := new(ExecutionPayloadHeader)
	require.NoError(t, DecodeJSON(bytes.NewReader(b), h2))
	require.Equal(t, h.ParentHash, h2.ParentHash)

	p, err := h2.HashTreeRoot()
	require.NoError(t, err)
	rootHex := fmt.Sprintf("%x", p)
	require.Equal(t, "31ffc5e97d80143da2f96bbc831a11b444c393d5b0c9a43f799ab2b3cbe29be4", rootHex)
}

func TestUnmarshalGetHeaderResponse(t *testing.T) {
	body := []byte(`{"version":"bellatrix","data":{"message":{"header":{"parent_hash":"0xe28385e7bd68df656cd0042b74b69c3104b5356ed1f20eb69f1f925df47a3ab7","fee_recipient":"0x0000000000000000000000000000000000000000","state_root":"0x0000000000000000000000000000000000000000000000000000000000000000","receipts_root":"0x0000000000000000000000000000000000000000000000000000000000000000","logs_bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","prev_randao":"0x0000000000000000000000000000000000000000000000000000000000000000","block_number":"0","gas_limit":"0","gas_used":"0","timestamp":"0","extra_data":"0x","base_fee_per_gas":"0","block_hash":"0xe28385e7bd68df656cd0042b74b69c3104b5356ed1f20eb69f1f925df47a3ab7","transactions_root":"0x0000000000000000000000000000000000000000000000000000000000000000"},"value":"12345","pubkey":"0x8a1d7b8dd64e0aafe7ea7b6c95065c9364cf99d38470c12ee807d55f7de1529ad29ce2c422e0b65e3d5a05c02caca249"},"signature":"0x88387c476a4792649e4fead76032ab0ae3c2108933b2aae32a3d90e01d472790f707ba49caad43bc65b169efebe4e6541085f53594ff01afcbfda79f6cae82de21b5dbee5668c6b5346f8a70cd18118c6119c42c335a8078101d195c0c9a0aa9"}}`)
	dst := &GetHeaderResponse{}
	require.NoError(t, DecodeJSON(bytes.NewReader(body), dst))
}

func TestBlindedBeaconBlock(t *testing.T) {
	parentHash := Hash{0xa1}
	blockHash := Hash{0xa1}
	feeRecipient := Address{0xb1}

	msg := &BlindedBeaconBlock{
		Slot:          1,
		ProposerIndex: 2,
		ParentRoot:    Root{0x03},
		StateRoot:     Root{0x04},
		Body: &BlindedBeaconBlockBody{
			Eth1Data: &Eth1Data{
				DepositRoot:  Root{0x05},
				DepositCount: 5,
				BlockHash:    Hash{0x06},
			},
			ProposerSlashings: []*ProposerSlashing{},
			AttesterSlashings: []*AttesterSlashing{},
			Attestations:      []*Attestation{},
			Deposits:          []*Deposit{},
			VoluntaryExits:    []*SignedVoluntaryExit{},
			SyncAggregate:     &SyncAggregate{CommitteeBits{0x07}, Signature{0x08}},
			ExecutionPayloadHeader: &ExecutionPayloadHeader{
				ParentHash:       parentHash,
				FeeRecipient:     feeRecipient,
				StateRoot:        Root{0x09},
				ReceiptsRoot:     Root{0x0a},
				LogsBloom:        Bloom{0x0b},
				Random:           Hash{0x0c},
				BlockNumber:      5001,
				GasLimit:         5002,
				GasUsed:          5003,
				Timestamp:        5004,
				ExtraData:        []byte{0x0d},
				BaseFeePerGas:    IntToU256(123456789),
				BlockHash:        blockHash,
				TransactionsRoot: Root{0x0e},
			},
		},
	}

	// Get HashTreeRoot
	root, err := msg.HashTreeRoot()
	require.NoError(t, err)
	expected := "9bcf3fc3b2b600ed054d5f08953b62c87e8870982fefb6314c1f7860a902090d"
	require.Equal(t, expected, common.Bytes2Hex(root[:]))

	// Marshalling
	b, err := json.Marshal(msg)
	require.NoError(t, err)

	expectedJSON := `{
        "slot": "1",
        "proposer_index": "2",
        "parent_root": "0x0300000000000000000000000000000000000000000000000000000000000000",
        "state_root": "0x0400000000000000000000000000000000000000000000000000000000000000",
        "body": {
            "randao_reveal": "0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
            "eth1_data": {
                "deposit_root": "0x0500000000000000000000000000000000000000000000000000000000000000",
                "deposit_count": "5",
                "block_hash": "0x0600000000000000000000000000000000000000000000000000000000000000"
            },
            "graffiti": "0x0000000000000000000000000000000000000000000000000000000000000000",
            "proposer_slashings": [],
            "attester_slashings": [],
            "attestations": [],
            "deposits": [],
            "voluntary_exits": [],
            "sync_aggregate": {
                "sync_committee_bits": "0x07000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
                "sync_committee_signature": "0x080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
            },
            "execution_payload_header": {
                "parent_hash": "0xa100000000000000000000000000000000000000000000000000000000000000",
                "fee_recipient": "0xb100000000000000000000000000000000000000",
                "state_root": "0x0900000000000000000000000000000000000000000000000000000000000000",
                "receipts_root": "0x0a00000000000000000000000000000000000000000000000000000000000000",
                "logs_bloom": "0x0b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
                "prev_randao": "0x0c00000000000000000000000000000000000000000000000000000000000000",
                "block_number": "5001",
                "gas_limit": "5002",
                "gas_used": "5003",
                "timestamp": "5004",
                "extra_data": "0x0d",
                "base_fee_per_gas": "123456789",
                "block_hash": "0xa100000000000000000000000000000000000000000000000000000000000000",
                "transactions_root": "0x0e00000000000000000000000000000000000000000000000000000000000000"
            }
        }
    }`
	require.JSONEq(t, expectedJSON, string(b))

	// Now unmarshal it back and compare to original
	msg2 := new(BlindedBeaconBlock)
	require.NoError(t, DecodeJSON(bytes.NewReader(b), msg2))
	require.Equal(t, msg, msg2)

	// HashTreeRoot
	root, err = msg2.HashTreeRoot()
	require.NoError(t, err)
	require.Equal(t, expected, common.Bytes2Hex(root[:]))
}

func TestMerkelizeTxs(t *testing.T) {
	txs := Transactions{}
	root, err := txs.HashTreeRoot()
	require.NoError(t, err)
	expected := "7ffe241ea60187fdb0187bfa22de35d1f9bed7ab061d9401fd47e34a54fbede1"
	require.Equal(t, expected, common.Bytes2Hex(root[:]))
}

func TestMerkelizePayload(t *testing.T) {
	input := `{"slot":"1","proposer_index":"7","parent_root":"0x7c1018e636481b7813e68a00af9f52f0d344f89eed431bb8a50618e2bc212dc6","state_root":"0xbaa15a02568c3e0442652c616f50cb60e8e11e86e2858fa7994e67a4017d6d3e","body":{"randao_reveal":"0xb6ea50c6ab03f159a893414161b2fb6d2ec61dc82868b13520acc180fc2d9b0d2d841d467295dbbae0e81bee7d3022060750f64879e5a3f0755380aa97710893d3e8cf2edac09e684c893999e3ef94f19231edf5b4fa46afe90ea1fb6b6c9e64","eth1_data":{"deposit_root":"0x23090150015e4c9d0c7ba87f97087375cdf19d6e2caeedc994d7c445b3460119","deposit_count":"32","block_hash":"0xccaf66b50e791f95d4b50bae4de28af9396824e7c29f99aeba19414fdf72673f"},"graffiti":"0x0000000000000000000000000000000000000000000000000000000000000000","proposer_slashings":[],"attester_slashings":[],"attestations":[{"aggregation_bits":"0x03","data":{"slot":"0","index":"0","beacon_block_root":"0x7c1018e636481b7813e68a00af9f52f0d344f89eed431bb8a50618e2bc212dc6","source":{"epoch":"0","root":"0x0000000000000000000000000000000000000000000000000000000000000000"},"target":{"epoch":"0","root":"0x7c1018e636481b7813e68a00af9f52f0d344f89eed431bb8a50618e2bc212dc6"}},"signature":"0xae9ec2c1bf76ec5a5d78c2a252dfb66a00f2828b3000d5b052f189064a836a864379afd3ce82f45517ff3a3b15b1c38d1551edde6352c07948e59596bdc97abd0be2cf27c6562bfb20cbacde37fab37eda7e5d1f73622e7e7fe1472a2bbd158a"}],"deposits":[],"voluntary_exits":[],"sync_aggregate":{"sync_committee_bits":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","sync_committee_signature":"0xc00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},"execution_payload_header":{"parent_hash":"0xccaf66b50e791f95d4b50bae4de28af9396824e7c29f99aeba19414fdf72673f","fee_recipient":"0x0000000000000000000000000000000000000000","state_root":"0xca3149fa9e37db08d1cd49c9061db1002ef1cd58db2210f2115c8c989b2bdf45","receipts_root":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","logs_bloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","prev_randao":"0xccaf66b50e791f95d4b50bae4de28af9396824e7c29f99aeba19414fdf72673f","block_number":"1","gas_limit":"30000000","gas_used":"0","timestamp":"1652735778","extra_data":"0x","base_fee_per_gas":"7","block_hash":"0x2244ab321090e7f53b51328d64d2a02f03ff9aa65f37208ec404cac8867a9dc3","transactions_root":"0x7ffe241ea60187fdb0187bfa22de35d1f9bed7ab061d9401fd47e34a54fbede1"}}}`
	var block BlindedBeaconBlock
	require.NoError(t, DecodeJSON(strings.NewReader(input), &block))

	// Verify block ssz
	out, err := block.MarshalSSZ()
	require.NoError(t, err)
	expected := common.Hex2Bytes("010000000000000007000000000000007c1018e636481b7813e68a00af9f52f0d344f89eed431bb8a50618e2bc212dc6baa15a02568c3e0442652c616f50cb60e8e11e86e2858fa7994e67a4017d6d3e54000000b6ea50c6ab03f159a893414161b2fb6d2ec61dc82868b13520acc180fc2d9b0d2d841d467295dbbae0e81bee7d3022060750f64879e5a3f0755380aa97710893d3e8cf2edac09e684c893999e3ef94f19231edf5b4fa46afe90ea1fb6b6c9e6423090150015e4c9d0c7ba87f97087375cdf19d6e2caeedc994d7c445b34601192000000000000000ccaf66b50e791f95d4b50bae4de28af9396824e7c29f99aeba19414fdf72673f0000000000000000000000000000000000000000000000000000000000000000800100008001000080010000690200006902000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006902000004000000e4000000000000000000000000000000000000007c1018e636481b7813e68a00af9f52f0d344f89eed431bb8a50618e2bc212dc60000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007c1018e636481b7813e68a00af9f52f0d344f89eed431bb8a50618e2bc212dc6ae9ec2c1bf76ec5a5d78c2a252dfb66a00f2828b3000d5b052f189064a836a864379afd3ce82f45517ff3a3b15b1c38d1551edde6352c07948e59596bdc97abd0be2cf27c6562bfb20cbacde37fab37eda7e5d1f73622e7e7fe1472a2bbd158a03ccaf66b50e791f95d4b50bae4de28af9396824e7c29f99aeba19414fdf72673f0000000000000000000000000000000000000000ca3149fa9e37db08d1cd49c9061db1002ef1cd58db2210f2115c8c989b2bdf4556e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b42100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ccaf66b50e791f95d4b50bae4de28af9396824e7c29f99aeba19414fdf72673f010000000000000080c3c90100000000000000000000000022bf8262000000001802000007000000000000000000000000000000000000000000000000000000000000002244ab321090e7f53b51328d64d2a02f03ff9aa65f37208ec404cac8867a9dc37ffe241ea60187fdb0187bfa22de35d1f9bed7ab061d9401fd47e34a54fbede1")
	require.Equal(t, out, expected)

	// Verify execution payload header htr
	root, err := block.Body.ExecutionPayloadHeader.HashTreeRoot()
	require.NoError(t, err)
	require.Equal(t, "75987a9d1b5e16c71bca93abd542c764b67db7e2512060733bd51e07e21e75ce", common.Bytes2Hex(root[:]))

	// Verify block body htr
	root, err = block.Body.HashTreeRoot()
	require.NoError(t, err)
	expected = common.Hex2Bytes("cf5f87f9645b481d25dc7b8783c3d7914b5e1b45fffda5b7c505b872189eb1d0")
	require.Equal(t, expected, root[:])

	// Verify block htr
	root, err = block.HashTreeRoot()
	require.NoError(t, err)
	require.Equal(t, "87b57a69321ec21e8a83a39f2f0f885a3be9bbddb80794b3b2700c3cf8230aa1", common.Bytes2Hex(root[:]))
}

func TestIndexedAttestation(t *testing.T) {
	input := `{
      "attesting_indices": [
        "1", "2", "3"
      ],
      "signature": "0x1c66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505cc411d61252fb6cb3fa0017b679f8bb2305b26a285fa2737f175668d0dff91cc1b66ac1fb663c9bc59509846d6ec05345bd908eda73e670af888da41af171505",
      "data": {
        "slot": "1",
        "index": "1",
        "beacon_block_root": "0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2",
        "source": {
          "epoch": "1",
          "root": "0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"
        },
        "target": {
          "epoch": "1",
          "root": "0xcf8e0d4e9587369b2301d0790347320302cc0943d5a1884560367e8208d920f2"
        }
      }
    }
    `
	var result IndexedAttestation
	require.NoError(t, DecodeJSON(strings.NewReader(input), &result))
}

func TestExecutionPayloadCase0(t *testing.T) {
	jsonFile, err := os.Open("../testdata/executionpayload/case0.json")
	require.NoError(t, err)
	defer jsonFile.Close()
	executionPayload := new(ExecutionPayload)
	require.NoError(t, DecodeJSON(jsonFile, &executionPayload))
}

func TestDepositCase0(t *testing.T) {
	jsonFile, err := os.Open("../testdata/deposit/case0.json")
	require.NoError(t, err)
	defer jsonFile.Close()
	deposit := new(Deposit)
	require.NoError(t, DecodeJSON(jsonFile, &deposit))

	_, err = deposit.HashTreeRoot()
	require.NoError(t, err)
	// require.Equal(t, "0x47eb8c0bd8d867c4854dbdbf5068e66c7d55129378339cf2c4c557f4266e9fb4", fmt.Sprintf("%#x", htr))
}

func TestDepositDataCase0(t *testing.T) {
	jsonFile, err := os.Open("../testdata/depositdata_case0.json")
	require.NoError(t, err)
	defer jsonFile.Close()
	depositData := new(DepositData)
	require.NoError(t, DecodeJSON(jsonFile, &depositData))

	_, err = depositData.HashTreeRoot()
	require.NoError(t, err)
	// require.Equal(t, "0xdcd7002d3d4804047bf559a5f642072467379f2db4b4b4b598a1e5f5d3a8269e", fmt.Sprintf("%#x", htr))
}

func TestEth1DataCase0(t *testing.T) {
	jsonFile, err := os.Open("../testdata/eth1data_case0.json")
	require.NoError(t, err)
	defer jsonFile.Close()
	eth1Data := new(Eth1Data)
	require.NoError(t, DecodeJSON(jsonFile, &eth1Data))

	_, err = eth1Data.HashTreeRoot()
	require.NoError(t, err)
	// require.Equal(t, "0x8fc4fbbff19ab83ac236352b82e3941a3aa87a784bc7c210c7bda8f1ab4cf854", fmt.Sprintf("%#x", htr))
}

func TestSignedBeaconBlock(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		txRoot    string
	}{
		{
			// https://github.com/ethereum/consensus-spec-tests/tree/master/tests/mainnet/bellatrix/ssz_static/SignedBeaconBlock/ssz_random/case_0
			name:      "case0",
			inputFile: "../testdata/signed-beacon-block-case0.json",
			txRoot:    "0xdf88d2d8cc3602bcc9b949c65308225bf336802ddd0ac4452a60950346be1b4b",
		},
		{
			// https://github.com/ethereum/consensus-spec-tests/tree/master/tests/mainnet/bellatrix/ssz_static/SignedBeaconBlock/ssz_random/case_1
			name:      "case1",
			inputFile: "../testdata/signed-beacon-block-case1.json",
			txRoot:    "0xb20e2a4663df68310fd2e4b5434018ad88757f5a6a0530936b759bca003228f9",
		},
		{
			// https://github.com/ethereum/consensus-spec-tests/tree/master/tests/mainnet/bellatrix/ssz_static/SignedBeaconBlock/ssz_random/case_2
			name:      "case2",
			inputFile: "../testdata/signed-beacon-block-case2.json",
			txRoot:    "0x846167c4733d3d0111523d6a1d05af04cdf94d4b00ee90595f820d1082833847",
		},
		{
			// https://github.com/ethereum/consensus-spec-tests/tree/master/tests/mainnet/bellatrix/ssz_static/SignedBeaconBlock/ssz_random/case_3
			name:      "case3",
			inputFile: "../testdata/signed-beacon-block-case3.json",
			txRoot:    "0xb1d1b761ebc76aa2bd90bed51eab6311200480dc4b6c5c363c50617241fd6841",
		},
		{
			// https://github.com/ethereum/consensus-spec-tests/tree/master/tests/mainnet/bellatrix/ssz_static/SignedBeaconBlock/ssz_random/case_4
			name:      "case4",
			inputFile: "../testdata/signed-beacon-block-case4.json",
			txRoot:    "0xbb892813737d34a771c525639aadc64c5f317ce4da0c3b1b800c4741fdc6b34f",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jsonFile, err := os.Open(test.inputFile)
			require.NoError(t, err)
			defer jsonFile.Close()
			signedBeaconBlock := new(SignedBeaconBlock)
			require.NoError(t, DecodeJSON(jsonFile, &signedBeaconBlock))

			header, err := PayloadToPayloadHeader(signedBeaconBlock.Message.Body.ExecutionPayload)
			require.NoError(t, err)
			require.Equal(t, test.txRoot, header.TransactionsRoot.String())
		})
	}
}

func TestSignedBlindedBeaconBlockWithDeposit(t *testing.T) {
	jsonFile, err := os.Open("../testdata/signed-blinded-beacon-block-with-deposit.json")
	require.NoError(t, err)
	defer jsonFile.Close()
	signedBlindedBeaconBlock := new(SignedBlindedBeaconBlock)
	require.NoError(t, DecodeJSON(jsonFile, &signedBlindedBeaconBlock))
}
