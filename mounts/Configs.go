package mounts

type Configs struct {
}

type Keys struct {
	PublicKey  string `json:"public_key" toml:"public_key"`
	PrivateKey string `json:"private_key" toml:"private_key"`
	SigKey     string `json:"sig_key" toml:"sig_key"`
	AesKey     string `json:"aes_key" toml:"aes_key"`
	IV         string `json:"iv" toml:"iv"`
}
