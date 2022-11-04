package resume

import (
	"github.com/pkg/errors"
	"html/template"
	"os"
	"resume-builder/shared/models"
	"resume-builder/utils"
)

func ParseToHtml(resumeData models.Resume) (string, error) {
	utils.EnsureOutputDir(utils.OutputDir)

	htmlOut, err := os.Create(utils.OutputHtmlFile)
	if err != nil {
		return "", errors.Wrap(err, "ParseToHtml - Create")
	}

	if resumeData.Template == "" {
		resumeData.Template = utils.BasicTemplate
	}
	t, err := template.ParseGlob("ui/" + resumeData.Template + "/*")
	if err != nil {
		return "", errors.Wrap(err, "ParseToHtml - ParseGlob")
	}
	err = t.Execute(htmlOut, resumeData)
	if err != nil {
		return "", errors.Wrap(err, "ParseToHtml - Execute")
	}

	utils.Logger.Infof("Html parsed in %s", htmlOut.Name())

	return htmlOut.Name(), nil
}
