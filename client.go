package ecpay

type ECPayClient struct {
	HashIV  string
	HashKey string
}

func NewECPayClient(hashIV string, hashKey string) *ECPayClient {
	return &ECPayClient{HashIV: hashIV, HashKey: hashKey}
}

func NewStageECPayClient() *ECPayClient {
	return &ECPayClient{HashIV: "5294y06JbISpM5x9", HashKey: "v77hoKGq4kWxNNIS"}
}