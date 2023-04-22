package models

import (
	"resumme-builder/internal/utils/lang"
	"resumme-builder/internal/utils/lang/en"
	"resumme-builder/internal/utils/lang/fr"
	"strings"
)

type Resume struct {
	Basics       Basics        `json:"basics"`
	Work         []Work        `json:"work"`
	Projects     []Project     `json:"projects"`
	Education    []Education   `json:"education"`
	Certificates []Certificate `json:"certificates"`
	Skills       []Skill       `json:"skills"`
	SoftSkills   []Skill       `json:"softSkills"`
	Languages    []Language    `json:"languages"`
	Interests    []Interest    `json:"interests"`
	Meta         Meta          `json:"meta"`
	Labels       ResumeLabels
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
}

type Profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type Work struct {
	Position     string   `json:"position"`
	Company      string   `json:"company"`
	ContractType string   `json:"contractType"`
	StartDate    string   `json:"startDate"`
	EndDate      any      `json:"endDate"`
	Summary      string   `json:"summary"`
	Location     string   `json:"location"`
	Highlights   []string `json:"highlights"`
}

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
	URL         string   `json:"url"`
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
	Score  string `json:"score"`
	Issuer string `json:"issuer"`
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

type ResumeLabels struct {
	Education   string
	Experiences string
	Projects    string
	Skills      string
	SoftSkills  string
	Languages   string
	Interests   string
	Profile     string
	Since       string
}

func (r *Resume) GetEducationLabel() string {
	switch strings.ToLower(r.Meta.Lang) {
	case lang.EN:
		return en.EducationLabel
	case lang.FR:
		return fr.EducationLabel
	default:
		return en.EducationLabel
	}
}

func (r *Resume) GetExperiencesLabel() string {
	switch strings.ToLower(r.Meta.Lang) {
	case lang.EN:
		return en.ExperiencesLabel
	case lang.FR:
		return fr.ExperiencesLabel
	default:
		return en.ExperiencesLabel
	}
}

func (r *Resume) GetSkillsLabel() string {
	switch strings.ToLower(r.Meta.Lang) {
	case lang.EN:
		return en.SkillsLabel
	case lang.FR:
		return fr.SkillsLabel
	default:
		return en.SkillsLabel
	}
}

func (r *Resume) GetSoftSkillsLabel() string {
	switch strings.ToLower(r.Meta.Lang) {
	case lang.EN:
		return en.SoftSkillsLabel
	case lang.FR:
		return fr.SoftSkillsLabel
	default:
		return en.SoftSkillsLabel
	}
}

func (r *Resume) GetProjectsLabel() string {
	switch strings.ToLower(r.Meta.Lang) {
	case lang.EN:
		return en.ProjectsLabel
	case lang.FR:
		return fr.ProjectsLabel
	default:
		return en.ProjectsLabel
	}
}

func (r *Resume) GetLanguagesLabel() string {
	switch strings.ToLower(r.Meta.Lang) {
	case lang.EN:
		return en.LanguagesLabel
	case lang.FR:
		return fr.LanguagesLabel
	default:
		return en.LanguagesLabel
	}
}

func (r *Resume) GetInterestsLabel() string {
	switch strings.ToLower(r.Meta.Lang) {
	case lang.EN:
		return en.InterestsLabel
	case lang.FR:
		return fr.InterestsLabel
	default:
		return en.InterestsLabel
	}
}

func (r *Resume) GetProfileLabel() string {
	switch strings.ToLower(r.Meta.Lang) {
	case lang.EN:
		return en.ProfileLabel
	case lang.FR:
		return fr.ProfileLabel
	default:
		return en.ProfileLabel
	}
}

func (r *Resume) GetSinceLabel() string {
	switch strings.ToLower(r.Meta.Lang) {
	case lang.EN:
		return en.SinceLabel
	case lang.FR:
		return fr.SinceLabel
	default:
		return en.SinceLabel
	}
}
