package modelos

// Representa o formato da alteracao de senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
