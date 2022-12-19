package dto

type Caesar struct {
	Shift    int    `json:"shift"`
	Text     string `json:"text"`
	Alphabet string `json:"alphabet"`
}
