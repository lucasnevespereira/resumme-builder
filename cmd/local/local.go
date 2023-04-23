package local

import (
	"github.com/spf13/cobra"
	"resumme-builder/internal/models"
	"resumme-builder/internal/pkg/pdf"
	"resumme-builder/internal/pkg/resume"
	"resumme-builder/internal/utils/logger"
)

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Generates output locally",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Log.Info("Generating output")

		resumeData, err := resume.ReadLocalData()
		if err != nil {
			logger.Log.Error(err)
		}

		htmlFile, err := resume.ParseToHtml(resumeData)
		if err != nil {
			return err
		}

		pdfData, err := pdf.GenerateFromHtml(htmlFile)
		if err != nil {
			return err
		}

		if err := pdf.Write(models.OutputPdfFile, pdfData); err != nil {
			return err
		}

		return nil
	},
}

func Cmd() *cobra.Command {
	return localCmd
}
