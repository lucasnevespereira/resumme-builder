package translator

import (
	"github.com/lucasnevespereira/deeplgo/deepl"
	"github.com/pkg/errors"
	"resumme-builder/internal/models"
	"resumme-builder/internal/utils/logger"
)

type Translator struct {
	client *deepl.Client
}

func New(authKey string) *Translator {
	return &Translator{client: deepl.NewClient(authKey)}
}

func (c *Translator) Work(works []models.Work, lang string) []models.Work {
	var summaries []string
	var positions []string
	var originalOrder []int

	for i, work := range works {
		summaries = append(summaries, work.Summary)
		positions = append(positions, work.Position)
		originalOrder = append(originalOrder, i)
	}

	fields := append(summaries, positions...)

	resp, err := c.client.Translate(fields, "", lang)
	if err != nil {
		logger.Log.Error(errors.Wrap(err, "Translator client translate"))
		return works
	}

	// split the translated fields into summaries and positions
	translatedSummaries := resp.Translations[:len(summaries)]
	translatedPositions := resp.Translations[len(summaries):]

	// loop through the works in their original order and update the translated summaries and positions
	for i, originalIndex := range originalOrder {
		works[originalIndex].Summary = translatedSummaries[i].Text
		works[originalIndex].Position = translatedPositions[i].Text
	}

	return works
}

func (c *Translator) Basics(basics models.Basics, lang string) models.Basics {
	fields := []string{basics.Label, basics.Summary}
	resp, err := c.client.Translate(fields, "", lang)
	if err != nil {
		logger.Log.Error(errors.Wrap(err, "Translator client translate"))
		return basics
	}

	basics.Label = resp.Translations[0].Text
	basics.Summary = resp.Translations[1].Text

	return basics
}

func (c *Translator) Projects(projects []models.Project, lang string) []models.Project {
	var descriptions []string
	var originalOrder []int

	for i, project := range projects {
		descriptions = append(descriptions, project.Description)
		originalOrder = append(originalOrder, i)
	}

	resp, err := c.client.Translate(descriptions, "", lang)
	if err != nil {
		logger.Log.Error(errors.Wrap(err, "Translator client translate"))
		return projects
	}

	for i, originalIndex := range originalOrder {
		projects[originalIndex].Description = resp.Translations[i].Text
	}

	return projects
}

func (c *Translator) Interests(interests []models.Interest, lang string) []models.Interest {
	var names []string
	var originalOrder []int

	for i, interest := range interests {
		names = append(names, interest.Name)
		originalOrder = append(originalOrder, i)
	}

	resp, err := c.client.Translate(names, "", lang)
	if err != nil {
		logger.Log.Error(errors.Wrap(err, "Translator client translate"))
		return interests
	}

	for i, originalIndex := range originalOrder {
		interests[originalIndex].Name = resp.Translations[i].Text
	}

	return interests
}

func (c *Translator) SoftSkills(softSkills []models.Skill, lang string) []models.Skill {
	var names []string
	var originalOrder []int

	for i, skill := range softSkills {
		names = append(names, skill.Name)
		originalOrder = append(originalOrder, i)
	}

	resp, err := c.client.Translate(names, "", lang)
	if err != nil {
		logger.Log.Error(errors.Wrap(err, "Translator client translate"))
		return softSkills
	}

	for i, originalIndex := range originalOrder {
		softSkills[originalIndex].Name = resp.Translations[i].Text
	}

	return softSkills
}

func (c *Translator) Languages(langs []models.Language, lang string) []models.Language {
	var languages []string
	var fluencies []string
	var originalOrder []int

	for i, lang := range langs {
		languages = append(languages, lang.Language)
		fluencies = append(fluencies, lang.Fluency)
		originalOrder = append(originalOrder, i)
	}

	fields := append(languages, fluencies...)

	resp, err := c.client.Translate(fields, "", lang)
	if err != nil {
		logger.Log.Error(errors.Wrap(err, "Translator client translate"))
		return langs
	}

	// split the translated fields
	translatedLanguages := resp.Translations[:len(languages)]
	translatedFluencies := resp.Translations[len(languages):]

	// loop through original order and update the translated fields
	for i, originalIndex := range originalOrder {
		langs[originalIndex].Language = translatedLanguages[i].Text
		langs[originalIndex].Fluency = translatedFluencies[i].Text
	}

	return langs
}

func (c *Translator) Education(education []models.Education, lang string) []models.Education {
	var studies []string
	var areas []string
	var originalOrder []int

	for i, edu := range education {
		studies = append(studies, edu.StudyType)
		areas = append(areas, edu.Area)
		originalOrder = append(originalOrder, i)
	}

	fields := append(studies, areas...)

	resp, err := c.client.Translate(fields, "", lang)
	if err != nil {
		logger.Log.Error(errors.Wrap(err, "Translator client translate"))
		return education
	}

	// split the translated fields
	translatedStudies := resp.Translations[:len(studies)]
	translatedAreas := resp.Translations[len(studies):]

	// loop through original order and update the translated fields
	for i, originalIndex := range originalOrder {
		education[originalIndex].StudyType = translatedStudies[i].Text
		education[originalIndex].Area = translatedAreas[i].Text
	}

	return education
}

func (c *Translator) Resume(resume models.Resume, lang string) models.Resume {
	resume.Basics = c.Basics(resume.Basics, lang)
	resume.Work = c.Work(resume.Work, lang)
	resume.Education = c.Education(resume.Education, lang)
	resume.Projects = c.Projects(resume.Projects, lang)
	resume.SoftSkills = c.SoftSkills(resume.SoftSkills, lang)
	resume.Languages = c.Languages(resume.Languages, lang)
	resume.Interests = c.Interests(resume.Interests, lang)
	return resume
}
