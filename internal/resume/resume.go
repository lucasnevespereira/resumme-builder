package resume

import (
	"github.com/pkg/errors"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"resumme-builder/internal/models"
	"resumme-builder/internal/utils"
)

func ParseToHtml(resume models.Resume) (string, error) {
	utils.EnsureOutputDir(utils.OutputDir)

	htmlOut, err := os.Create(utils.OutputHtmlFile)
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
		resume.Template = utils.ClassicTemplate
	}

	templateFiles, err := filepath.Glob("ui/" + resume.Template + "/*")
	if err != nil {
		return "", errors.Wrap(err, "ParseToHtml - filepath.Glob")
	}
	templateFuncs := template.FuncMap{
		"isLast": func(index, length int) bool {
			return index == length-1
		},
	}
	t := template.New(path.Base(templateFiles[0])).Funcs(templateFuncs)
	t, err = t.ParseFiles(templateFiles...)
	if err != nil {
		return "", errors.Wrap(err, "ParseToHtml - ParseGlob")
	}

	err = t.Execute(htmlOut, resume)
	if err != nil {
		return "", errors.Wrap(err, "ParseToHtml - Execute")
	}

	utils.Logger.Infof("Html parsed in %s", htmlOut.Name())

	return htmlOut.Name(), nil
}
