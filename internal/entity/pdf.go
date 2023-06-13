package entity

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/unidoc/unipdf/v3/core"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
	pdf "github.com/unidoc/unipdf/v3/model"
)

type PDF struct {
	reader *pdf.PdfReader
}

func NewPDF(reader *pdf.PdfReader) *PDF {
	return &PDF{reader: reader}
}

func (p *PDF) ThumbNail(path string) error {
	log.Info().Msg("Extracting thumbnail from instance PDF")
	page, err := p.reader.GetPage(1)

	c := creator.New()

	err = c.AddPage(page)
	if err != nil {
		panic(err)
	}
	log.Info().Msg("Finish Extract thumbnail from instance PDF")

	log.Info().Msg("Writing thumbnail to disk: " + path)
	return c.WriteToFile(path)
}

func (p *PDF) GetPdfInfo() (*PdfInfo, error) {
	log.Info().Msg("Start getting pdf info from instance PDF")
	pdfInfo := PdfInfo{}

	_source, err := p.reader.GetPdfInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}
	log.Info().Msg("Finish getting pdf info from instance PDF")

	if _source.Author != nil {
		pdfInfo.Author = _source.Author.String()
	}

	if _source.Subject != nil {
		pdfInfo.Subject = _source.Subject.String()
	}

	if _source.Title != nil {
		pdfInfo.Title = _source.Title.String()
	}

	if _source.Keywords != nil {
		pdfInfo.Keywords = _source.Keywords.String()
	}

	pdfInfo.Extras = _source

	return &pdfInfo, nil
}

func (p *PDF) SetPdfInfo(outpathPath string, info *PdfInfo, custom map[string]string) error {
	log.Info().Msg("Start set a new metadata to instance PDF")
	author := info.Author
	model.SetPdfAuthor(author)

	pdfReader := p.GetPdfReader()

	// Don't copy document info into the new PDF.
	opt := &model.ReaderToWriterOpts{
		SkipInfo: true,
	}

	// Generate a PdfWriter instance from existing PdfReader.
	defaultPdfWriter, err := pdfReader.ToWriter(opt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	customPdfWriter := defaultPdfWriter

	// Write new PDF with custom information dictionary.
	// Get the instance from existing information dictionary.
	pdfInfo := info.Extras
	pdfInfo.Author = core.MakeString(info.Author)
	pdfInfo.Subject = core.MakeString(info.Subject)
	pdfInfo.Creator = core.MakeString("Vldcreation's pdf @c UniDoc")
	pdfInfo.Producer = core.MakeString("UniDoc v3.37.0 (Metered subscription) - http://unidoc.io")

	key, val := custom["key"], custom["value"]
	pdfInfo.AddCustomInfo(key, val)

	customPdfWriter.SetDocInfo(pdfInfo)
	log.Info().Msg("Finish set a new metadata to instance PDF")

	log.Info().Msg("Write a new PDF with custom information dictionary to disk: " + outpathPath)
	err = customPdfWriter.WriteToFile(outpathPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	return nil
}

func (p *PDF) GetPdfReader() *pdf.PdfReader {
	return p.reader
}
