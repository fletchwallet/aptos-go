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

type Transaction struct {
	Type                string           `json:"type"`
	Events              []TxEvents       `json:"events"`
	Payload             *WriteSetPayload `json:"payload"`
	Version             uint64           `json:"version,string"`
	SequenceNumber      uint64           `json:"sequence_number,string"`
	MaxGasAmount        uint64           `json:"max_gas_amount,string"`
	GasUnitPrice        uint64           `json:"gas_unit_price,string"`
	GasCurrencyCode     string           `json:"gas_currency_code"`
	ExpirationTime      uint64           `json:"expiration_timestamp_secs,string"`
	Sender              string           `json:"sender"`
	Hash                string           `json:"hash"`
	StateRootHash       string           `json:"state_root_hash"`
	EventRootHash       string           `json:"event_root_hash"`
	GasUsed             uint64           `json:"gas_used,string"`
	Success             bool             `json:"success"`
	VMStatus            string           `json:"vm_status"`
	AccumulatorRootHash string           `json:"accumulator_root_hash"`
	Changes             interface{}      `json:"changes"`
	Signature           interface{}      `json:"signature"`
}

type TxEvents struct {
	Key            string      `json:"key"`
	SequenceNumber uint64      `json:"sequence_number"`
	Type           string      `json:"type"`
	Data           interface{} `json:"data"`
}

type WriteSetPayload struct {
	Type     string      `json:"type"`
	WriteSet interface{} `json:"write_set"`
}
