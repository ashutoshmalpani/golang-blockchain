package blockchain

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	Id  []byte
	Out int
	Sig string
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
