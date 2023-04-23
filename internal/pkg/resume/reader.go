package resume

import (
	"os"
	"resumme-builder/internal/models"
	"resumme-builder/internal/utils/json"
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
