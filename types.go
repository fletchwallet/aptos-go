package aptos

type LedgerInfo struct {
	ChainID         int32  `json:"chain_id"`
	LedgerVersion   uint   `json:"ledger_version,string"`
	LedgerTimestamp uint64 `json:"ledger_timestamp,string"`
}

type Account struct {
	SequenceNumber uint64 `json:"sequence_number,string"`
	AuthKey        string `json:"authentication_key"`
}

type AccountResource struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type AccountModule struct {
	ByteCode string         `json:"bytecode"`
	ABI      *MoveModuleABI `json:"abi"`
}

type MoveModuleABI struct {
	Address string   `json:"address"`
	Name    string   `json:"string"`
	Friends []string `json:"friends"`
}

type MoveFunction struct {
	Name              string        `json:"name"`
	Visibility        string        `json:"visibility"`
	GenericTypeParams []interface{} `json:"generic_type_params"`
	Params            []string      `json:"params"`
	Returns           []string      `json:"returns"`
}

type MoveStruct struct {
	Name              string        `json:"name"`
	IsNative          bool          `json:"is_native"`
	Abilities         []string      `json:"abilities"`
	GenericTypeParams []interface{} `json:"generic_type_params"`
}

type MoveStructField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
