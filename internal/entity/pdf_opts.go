package entity

import "github.com/unidoc/unipdf/v3/model"

type PdfOpts struct {
	isEncrypt   bool
	encryptOpts *model.ReaderOpts
	outputPath  *string
}

func NewPdfOpts(isEncrypt bool, encryptOpts *model.ReaderOpts, outputPath *string) *PdfOpts {
	return &PdfOpts{
		isEncrypt:   isEncrypt,
		encryptOpts: encryptOpts,
		outputPath:  outputPath,
	}
}

func (p *PdfOpts) SetIsEncrypt(isEncrypt bool) {
	p.isEncrypt = isEncrypt
}

func (p *PdfOpts) IsEncrypt() bool {
	return p.isEncrypt
}

func (p *PdfOpts) SetEncryptOpts(opts *model.ReaderOpts) {
	p.encryptOpts = opts
}

func (p *PdfOpts) EncryptOpts() *model.ReaderOpts {
	return p.encryptOpts
}

func (p *PdfOpts) SetOutputPath(outputPath *string) {
	p.outputPath = outputPath
}

func (p *PdfOpts) OutputPath() *string {
	return p.outputPath
}
