package aptos

func NewAptosClient(nodeURL string) *AptosClient {
	return &AptosClient{
		nodeURL: nodeURL,
	}
}

type AptosClient struct {
	nodeURL string
}
