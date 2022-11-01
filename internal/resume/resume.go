package resume

import (
	"encoding/json"
	"github.com/pkg/errors"
	"html/template"
	"os"
	"resume-builder/utils"
)

func ParseToHtml(templateFiles string) error {
	utils.EnsureOutputDir(utils.OutputDir)

	htmlOut, err := os.Create(utils.OutputHtmlFile)
	if err != nil {
		return errors.Wrap(err, "ParseToHtml - Create")
	}

	var data map[string]interface{}

	file, err := os.ReadFile(utils.ResumeDataFile)
	if err != nil {
		return errors.Wrap(err, "ParseToHtml - ReadFile")
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return errors.Wrap(err, "ParseToHtml - Unmarshal")
	}
	t, err := template.ParseGlob(templateFiles)
	if err != nil {
		return errors.Wrap(err, "ParseToHtml - ParseGlob")
	}
	err = t.Execute(htmlOut, data)
	if err != nil {
		return errors.Wrap(err, "ParseToHtml - Execute")
	}

	return nil
}
