package resume

import (
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"resumme-builder/internal/models"
	"resumme-builder/internal/utils/json"
	"resumme-builder/internal/utils/logger"
	"strings"
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
	ensureOutputDir(models.OutputDir)

	htmlOut, err := os.Create(models.OutputHtmlFile)
	if err != nil {
		return "", errors.Wrap(err, "ParseToHtml - Create")
	}

	resume.Labels.Education = resume.GetEducationLabel()
	resume.Labels.Experiences = resume.GetExperiencesLabel()
	resume.Labels.Projects = resume.GetProjectsLabel()
	resume.Labels.Skills = resume.GetSkillsLabel()
	resume.Labels.SoftSkills = resume.GetSoftSkillsLabel()
	resume.Labels.Languages = resume.GetLanguagesLabel()
	resume.Labels.Interests = resume.GetInterestsLabel()
	resume.Labels.Profile = resume.GetProfileLabel()
	resume.Labels.Since = resume.GetSinceLabel()

	if resume.Meta.Template == "" {
		resume.Meta.Template = models.ClassicTemplate
	}

	t, err := getTemplate(resume.Meta.Template)
	if err != nil {
		return "", err
	}

	err = t.Execute(htmlOut, resume)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("ParseToHtml %s - Execute", resume.Meta.Template))
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
		"displayLocation": func(location models.Location) string {
			var parts []string

			if location.City != "" {
				parts = append(parts, location.City)
			}
			if location.Region != "" {
				parts = append(parts, location.Region)
			}
			if location.CountryCode != "" {
				parts = append(parts, location.CountryCode)
			}

			locationStr := strings.Join(parts, ", ")
			locationStr = strings.TrimSuffix(locationStr, ", ")

			return locationStr
		},
		"trimURLPrefix": func(url string) string {
			prefixes := []string{"http://", "https://", "https://www.", "www."}
			for _, prefix := range prefixes {
				if strings.HasPrefix(url, prefix) {
					return strings.TrimPrefix(url, prefix)
				}
			}
			return url
		},
		"getFirstName": func(name string) string {
			parts := strings.Split(name, " ")
			if len(parts) > 0 {
				return parts[0]
			}
			return name
		},
		"getLastName": func(name string) string {
			parts := strings.Split(name, " ")
			if len(parts) > 1 {
				return strings.Join(parts[1:], " ")
			}
			return name
		},
	}

	t := template.New(path.Base(templateFiles[0])).Funcs(templateFuncs)
	t, err = t.ParseFiles(templateFiles...)
	if err != nil {
		return nil, errors.Wrap(err, "getTemplate template.New")
	}

	return t, nil
}

func ensureOutputDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			logger.Log.Errorf("EnsureOutputDir: %v", err)
		}
	}
}
