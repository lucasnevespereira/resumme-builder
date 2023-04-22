package resume

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"resumme-builder/internal/models"
	"resumme-builder/internal/utils/logger"
	"time"
)

func ParseToHtml(resume models.Resume) (string, error) {
	startedAt := time.Now()

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

	logger.Log.Infof("Html %s generated in %f seconds", htmlOut.Name(), time.Since(startedAt).Seconds())

	return htmlOut.Name(), nil
}

func ensureOutputDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			logger.Log.Errorf("EnsureOutputDir: %v", err)
		}
	}
}
