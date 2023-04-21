package models

import (
	"resumme-builder/internal/utils/lang"
	"resumme-builder/internal/utils/lang/en"
	"resumme-builder/internal/utils/lang/fr"
	"strings"
)

type Resume struct {
	Data     ResumeData `json:"data"`
	Template string     `json:"template"`
	Lang     string     `json:"lang"`
}

type ResumeData struct {
	FirstName      string          `json:"firstName"`
	LastName       string          `json:"LastName"`
	JobPosition    string          `json:"jobPosition"`
	Profile        string          `json:"profile"`
	Photo          string          `json:"photo"`
	Email          string          `json:"email"`
	Phone          string          `json:"phone"`
	Socials        Socials         `json:"socials"`
	Location       string          `json:"location"`
	SoftSkills     []string        `json:"softSkills"`
	Skills         []string        `json:"skills"`
	Experiences    []Experience    `json:"experiences"`
	Projects       []Project       `json:"projects"`
	Education      []Education     `json:"education"`
	Languages      []Language      `json:"languages"`
	Certifications []Certification `json:"certifications"`
	Hobbies        []string        `json:"hobbies"`
	Labels         ResumeLabels    `json:"labels"`
}

type ResumeLabels struct {
	Education   string `json:"educationLabel"`
	Experiences string `json:"experiencesLabel"`
	Projects    string `json:"projectsLabel"`
	Skills      string `json:"skillsLabel"`
	SoftSkills  string `json:"softSkillsLabel"`
	Languages   string `json:"languagesLabel"`
	Hobbies     string `json:"hobbiesLabel"`
	Profile     string `json:"profileLabel"`
	Since       string `json:"sinceLabel"`
}

type Socials struct {
	Website  string `json:"website"`
	Linkedin string `json:"linkedin"`
	Github   string `json:"github"`
}

type Language struct {
	Label string `json:"label"`
	Level string `json:"level"`
}

type Certification struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type Project struct {
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

type Experience struct {
	Position    string `json:"position"`
	Company     string `json:"company"`
	Started     string `json:"started"`
	Stopped     string `json:"stopped"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

type Education struct {
	School      string `json:"school"`
	Degree      string `json:"degree"`
	Location    string `json:"location"`
	Started     string `json:"started"`
	Stopped     string `json:"stopped"`
	Description string `json:"description"`
}

func (r *Resume) GetEducationLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.EducationLabel
	case lang.FR:
		return fr.EducationLabel
	default:
		return en.EducationLabel
	}
}

func (r *Resume) GetExperiencesLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.ExperiencesLabel
	case lang.FR:
		return fr.ExperiencesLabel
	default:
		return en.ExperiencesLabel
	}
}

func (r *Resume) GetSkillsLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.SkillsLabel
	case lang.FR:
		return fr.SkillsLabel
	default:
		return en.SkillsLabel
	}
}

func (r *Resume) GetSoftSkillsLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.SoftSkillsLabel
	case lang.FR:
		return fr.SoftSkillsLabel
	default:
		return en.SoftSkillsLabel
	}
}

func (r *Resume) GetProjectsLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.ProjectsLabel
	case lang.FR:
		return fr.ProjectsLabel
	default:
		return en.ProjectsLabel
	}
}

func (r *Resume) GetLanguagesLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.LanguagesLabel
	case lang.FR:
		return fr.LanguagesLabel
	default:
		return en.LanguagesLabel
	}
}

func (r *Resume) GetHobbiesLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.HobbiesLabel
	case lang.FR:
		return fr.HobbiesLabel
	default:
		return en.HobbiesLabel
	}
}

func (r *Resume) GetProfileLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.ProfileLabel
	case lang.FR:
		return fr.ProfileLabel
	default:
		return en.ProfileLabel
	}
}

func (r *Resume) GetSinceLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.SinceLabel
	case lang.FR:
		return fr.SinceLabel
	default:
		return en.SinceLabel
	}
}
