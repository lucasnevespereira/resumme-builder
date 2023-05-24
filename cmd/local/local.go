package local

import (
	"github.com/spf13/cobra"
	"resumme-builder/internal/models"
	"resumme-builder/internal/pkg/parser"
	"resumme-builder/internal/pkg/pdf"
	"resumme-builder/internal/pkg/template"
	"resumme-builder/internal/services"
	"resumme-builder/internal/utils/fs"
	"resumme-builder/internal/utils/logger"
)

var resumeDataFile string

func init() {
	localCmd.Flags().StringVarP(&resumeDataFile, "file", "f", "", "Resume data file")
	err := localCmd.MarkFlagRequired("file")
	if err != nil {
		logger.Log.Error("Failed to mark 'file' flag as required:", err)
	}
}

var localCmd = &cobra.Command{
	Use:     "local",
	Short:   "Generates output locally",
	PreRunE: preRunLocalCommand,
	RunE:    runLocalCommand,
}

func preRunLocalCommand(cmd *cobra.Command, args []string) error {
	filePath := cmd.Flag("file").Value.String()
	err := fs.EnsureNonEmptyFile(filePath)
	if err != nil {
		return err
	}

	return nil
}

func runLocalCommand(cmd *cobra.Command, args []string) error {
	logger.Log.Info("Generating output")

	templateManager := template.NewTemplateManager("ui")
	parser := parser.NewHTMLParser(models.OutputDir, models.OutputHtmlFile, templateManager)
	pdfGenerator := pdf.NewPDFGenerator()
	resumeService := services.NewResumeService(parser, pdfGenerator)

	fileData, err := fs.ReadFile(resumeDataFile)
	if err != nil {
		return err
	}
	resumeData, err := resumeService.UnmarshalResume(fileData)
	if err != nil {
		return err
	}

	if err := resumeService.GeneratePDF(resumeData); err != nil {
		return err
	}

	return nil
}

func Cmd() *cobra.Command {
	return localCmd
}
