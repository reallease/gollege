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

	// AnswerOne   string `json:"answer1"`
	// AnswerTwo   string `json:"answer2"`
	// AnswerThree string `json:"answer3"`
	// AnswerFour  string `json:"answer4"`
	// AnswerFive  string `json:"answer5"`

	// CorrectAlternative string `json:"correctAlternative"`
}
