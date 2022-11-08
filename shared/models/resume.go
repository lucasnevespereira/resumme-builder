package models

import (
	"resume-builder/utils/lang"
	"resume-builder/utils/lang/en"
	"resume-builder/utils/lang/fr"
	"strings"
)

type Resume struct {
	Data     ResumeData `json:"data"`
	Template string     `json:"template"`
	Lang     string     `json:"lang"`
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
	Labels     ResumeLabels `json:"labels"`
}

type ResumeLabels struct {
	Education   string `json:"educationLabel"`
	Experiences string `json:"experiencesLabel"`
	Projects    string `json:"projectsLabel"`
	Skills      string `json:"skillsLabel"`
	Languages   string `json:"languagesLabel"`
	Present     string `json:"presentLabel"`
}

type Experience struct {
	Role     string   `json:"role"`
	Company  string   `json:"company"`
	Started  string   `json:"started"`
	Stopped  string   `json:"stopped"`
	Location string   `json:"location"`
	Details  []string `json:"details"`
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

func (r *Resume) GetProjectsLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.ProjectsLabel
	case lang.FR:
		return fr.ProjectsLabel
	default:
		return en.PresentLabel
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

func (r *Resume) GetPresentLabel() string {
	switch strings.ToLower(r.Lang) {
	case lang.EN:
		return en.PresentLabel
	case lang.FR:
		return fr.PresentLabel
	default:
		return en.PresentLabel
	}
}
