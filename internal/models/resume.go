package models

type Resume struct {
	Basics       Basics         `json:"basics"`
	Work         []Work         `json:"work"`
	Volunteer    []Volunteer    `json:"volunteer"`
	Projects     []Project      `json:"projects"`
	Publications []Publications `json:"publications"`
	Education    []Education    `json:"education"`
	Certificates []Certificate  `json:"certificates"`
	Skills       []Skill        `json:"skills"`
	SoftSkills   []Skill        `json:"softSkills"`
	Languages    []Language     `json:"languages"`
	Interests    []Interest     `json:"interests"`
	Meta         Meta           `json:"meta"`
}

type Basics struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Image    string    `json:"image"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Summary  string    `json:"summary"`
	Location Location  `json:"location"`
	URL      string    `json:"url"`
	Profiles []Profile `json:"profiles"`
}

type Location struct {
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
	MapURL      string `json:"mapUrl"`
}

type Profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type Work struct {
	Position     string   `json:"position"`
	Company      string   `json:"company"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Location     string   `json:"location"`
	Highlights   []string `json:"highlights"`
	ContractType string   `json:"contractType"`
	CompanyLogo  *string  `json:"companyLogo,omitempty"`
	TeamDetails  *string  `json:"teamDetails,omitempty"`
	StackDetails *string  `json:"stackDetails,omitempty"`
	CompanyURL   *string  `json:"companyURL,omitempty"`
}

type Volunteer struct {
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	URL          string   `json:"url"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}

type Project struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Highlights   []string `json:"highlights"`
	URL          string   `json:"url"`
	StackDetails *string  `json:"stackDetails,omitempty"`
}

type Publications struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"releaseDate"`
	Summary     string `json:"summary"`
	URL         string `json:"url"`
}

type Education struct {
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	Location    string   `json:"location"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Score       string   `json:"score"`
	Courses     []string `json:"courses"`
}

type Certificate struct {
	Title  string `json:"title"`
	Date   string `json:"date"`
	Issuer string `json:"issuer"`
	Score  string `json:"score"`
	URL    string `json:"url"`
}

type Skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}

type Language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type Interest struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}

type Meta struct {
	Template string `json:"template"`
	Lang     string `json:"lang"`
}
