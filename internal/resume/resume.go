package resume

import (
	"github.com/pkg/errors"
	"html/template"
	"os"
	"resume-builder/shared/models"
	"resume-builder/utils"
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
	resume.Data.Labels.Languages = resume.GetLanguagesLabel()
	resume.Data.Labels.Present = resume.GetPresentLabel()

	if resume.Template == "" {
		resume.Template = utils.BasicTemplate
	}
	t, err := template.ParseGlob("ui/" + resume.Template + "/*")
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
