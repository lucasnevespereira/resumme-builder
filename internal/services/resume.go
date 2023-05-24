package services

import (
	"github.com/pkg/errors"
	"resumme-builder/internal/models"
	"resumme-builder/internal/pkg/parser"
	"resumme-builder/internal/pkg/pdf"
	"resumme-builder/internal/utils/fs"
	"resumme-builder/internal/utils/json"
)

type ResumeService struct {
	Parser *parser.HTMLParser
	Pdf    *pdf.Generator
}

func NewResumeService(parser *parser.HTMLParser, pdf *pdf.Generator) *ResumeService {
	return &ResumeService{
		Parser: parser,
		Pdf:    pdf,
	}
}

func (s *ResumeService) GeneratePDF(resumeData models.Resume) error {
	htmlFile, err := s.Parser.ParseToHtml(resumeData)
	if err != nil {
		return err
	}

	pdfData, err := s.Pdf.GenerateFromHTML(htmlFile)
	if err != nil {
		return err
	}

	if err := fs.WriteFile(models.OutputPdfFile, pdfData); err != nil {
		return err
	}

	return nil
}

func (s *ResumeService) UnmarshalResume(data []byte) (models.Resume, error) {
	var resumeData models.Resume
	err := json.Unmarshal(data, &resumeData)
	if err != nil {
		return models.Resume{}, errors.Wrap(err, "failed to unmarshal JSON")
	}

	return resumeData, nil
}
