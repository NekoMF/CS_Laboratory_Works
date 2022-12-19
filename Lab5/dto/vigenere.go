package dto

type Vigenere struct {
	Key      string `json:"key"`
	Text     string `json:"text"`
	Alphabet string `json:"alphabet"`
}
