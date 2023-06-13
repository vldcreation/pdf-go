package usecase

import (
	"log"

	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/document/convert"
	"github.com/vldcreation/privy-pdf-go/internal/entity"
)

func ExtractFromDoc(inputPath, outputPath string, opt *convert.Options) (*entity.PDF, error) {
	doc, err := document.Open(inputPath)
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	defer doc.Close()

	// opt @options
	// co := &convert.Options{
	// 	ProcessFields: true,
	// }
	c := convert.ConvertToPdfWithOptions(doc, opt)

	err = c.WriteToFile(outputPath)
	if err != nil {
		log.Fatalf("error converting document: %s", err)
	}

	return LoadPDF(outputPath, nil)
}
