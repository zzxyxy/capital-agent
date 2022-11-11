package connection

type LoginJson struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
