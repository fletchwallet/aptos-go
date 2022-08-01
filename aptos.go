package aptos

type AptosClient struct {
	nodeURL string
}

func NewAptosClient(nodeURL string) *AptosClient {
	return &AptosClient{
		nodeURL: nodeURL,
	}
}

