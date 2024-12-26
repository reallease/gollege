package questions

type Alternative struct {
	Statement string `json:"statement"`
	Correct bool `json:"correct"`
	Files []string `json:"files"`
}

type Question struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`

	Files []string `json:"files"`

	Question  string `json:"question"`
	Statement string `json:"statement"`
	Subject string `json:"subject"`

	Copyright string `json:"copyright"`

	Alternatives []Alternative 

	Credits string `json:"credits"`
}