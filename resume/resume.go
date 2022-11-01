package resume

import (
	"encoding/json"
	"errors"
	"html/template"
	"os"
	"resume-builder/utils"
)

func ParseToHtml() error {
	htmlOut, err := os.Create(utils.OutputHtmlFile)
	if err != nil {
		return err
	}

	var data map[string]interface{}

	file, err := os.ReadFile(utils.ResumeDataFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}
	t, err := template.ParseGlob(utils.TemplateFiles)
	if err != nil {
		return err
	}
	err = t.Execute(htmlOut, data)
	if err != nil {
		return err
	}

	return nil
}
