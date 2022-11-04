package models

type Resume struct {
	Data     ResumeData `json:"data"`
	Template string     `json:"template"`
}

type ResumeData struct {
	Name       string       `json:"name"`
	Job        string       `json:"job"`
	Email      string       `json:"email"`
	Website    string       `json:"website"`
	Location   string       `json:"location"`
	Mission    string       `json:"mission"`
	Skills     []string     `json:"skills"`
	Experience []Experience `json:"experience"`
	Projects   []string     `json:"projects"`
	Education  []string     `json:"education"`
	Languages  []string     `json:"languages"`
}

type Experience struct {
	Role     string   `json:"role"`
	Company  string   `json:"company"`
	Started  string   `json:"started"`
	Stopped  string   `json:"stopped"`
	Location string   `json:"location"`
	Details  []string `json:"details"`
}
