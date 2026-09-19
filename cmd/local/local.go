package local

import (
	"os"
	"path/filepath"
	"strings"

	"resumme-builder/internal/models"
	"resumme-builder/internal/pkg/pdf"
	"resumme-builder/internal/render"
	"resumme-builder/internal/utils/fs"
	"resumme-builder/internal/utils/json"
	"resumme-builder/internal/utils/logger"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var resumeDataFile, resumeUIDir, outputPdfFilename string

func init() {
	localCmd.Flags().StringVarP(&resumeDataFile, "file", "f", "", "Resume data file")
	localCmd.Flags().StringVarP(&outputPdfFilename, "name", "n", "", "Output PDF file name")
	localCmd.Flags().StringVarP(&resumeUIDir, "ui", "u", "ui", "UI data directory")
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

	// Set the global outputPdfFilename variable
	outputPdfFilename = cmd.Flag("name").Value.String()
	if outputPdfFilename == "" {
		baseFilename := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
		outputPdfFilename = filepath.Join(models.OutputDir, baseFilename+".pdf")
	} else {
		// Extract only the filename (no path) and place it in output directory
		filename := filepath.Base(outputPdfFilename)
		outputPdfFilename = filepath.Join(models.OutputDir, filename)
	}
	logger.Log.Info("Output PDF file name:", outputPdfFilename)

	uiDir := cmd.Flag("ui").Value.String()
	return fs.EnsureDir(uiDir)
}

func runLocalCommand(cmd *cobra.Command, args []string) error {
	logger.Log.Info("Generating output")

	printer := pdf.NewPDFGenerator()
	renderer, err := render.New(os.DirFS(resumeUIDir), printer)
	if err != nil {
		return err
	}

	fileData, err := fs.ReadFile(resumeDataFile)
	if err != nil {
		return err
	}
	var resumeData models.Resume
	if err := json.Unmarshal(fileData, &resumeData); err != nil {
		return errors.Wrap(err, "failed to unmarshal JSON")
	}

	html, err := renderer.HTML(resumeData)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(models.OutputDir, 0o755); err != nil {
		return err
	}
	if err := fs.WriteFile(models.OutputHtmlFile, html); err != nil {
		return err
	}

	pdfData, err := printer.Print(cmd.Context(), html)
	if err != nil {
		return err
	}
	return fs.WriteFile(outputPdfFilename, pdfData)
}

func Cmd() *cobra.Command {
	return localCmd
}
