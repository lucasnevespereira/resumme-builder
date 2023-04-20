package resume

import (
	"encoding/json"
	"github.com/pkg/errors"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"resumme-builder/internal/models"
	"resumme-builder/internal/utils"
	"resumme-builder/internal/utils/logger"
)

func ReadLocalData() (models.Resume, error) {
	fileData, err := os.ReadFile(models.ResumeDataFile)
	if err != nil {
		return models.Resume{}, err
	}

	var resumeData models.Resume
	err = json.Unmarshal(fileData, &resumeData)
	if err != nil {
		return models.Resume{}, err
	}

	return resumeData, nil
}

func ParseToHtml(resume models.Resume) (string, error) {
	utils.EnsureOutputDir(models.OutputDir)

	htmlOut, err := os.Create(models.OutputHtmlFile)
	if err != nil {
		return "", errors.Wrap(err, "ParseToHtml - Create")
	}

	resume.Data.Labels.Education = resume.GetEducationLabel()
	resume.Data.Labels.Experiences = resume.GetExperiencesLabel()
	resume.Data.Labels.Projects = resume.GetProjectsLabel()
	resume.Data.Labels.Skills = resume.GetSkillsLabel()
	resume.Data.Labels.SoftSkills = resume.GetSoftSkillsLabel()
	resume.Data.Labels.Languages = resume.GetLanguagesLabel()
	resume.Data.Labels.Hobbies = resume.GetHobbiesLabel()
	resume.Data.Labels.Since = resume.GetSinceLabel()

	if resume.Template == "" {
		resume.Template = models.ClassicTemplate
	}

	t, err := getTemplate(resume.Template)
	if err != nil {
		return "", err
	}

	err = t.Execute(htmlOut, resume)
	if err != nil {
		return "", errors.Wrap(err, "ParseToHtml - Execute")
	}

	logger.Log.Infof("Html parsed in %s", htmlOut.Name())

	return htmlOut.Name(), nil
}

func getTemplate(name string) (*template.Template, error) {
	templateFiles, err := filepath.Glob("ui/" + name + "/*")
	if err != nil {
		return nil, errors.Wrap(err, "getTemplate filepath.Glob")
	}

	templateFuncs := template.FuncMap{
		"isLast": func(index, length int) bool {
			return index == length-1
		},
	}

	t := template.New(path.Base(templateFiles[0])).Funcs(templateFuncs)
	t, err = t.ParseFiles(templateFiles...)
	if err != nil {
		return nil, errors.Wrap(err, "getTemplate template.New")
	}

	return t, nil
}
