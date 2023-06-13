package entity

import "github.com/unidoc/unipdf/v3/model"

type PdfInfo struct {
	Title    string
	Author   string
	Subject  string
	Keywords string

	Extras *model.PdfInfo
}

func NewPdfInfo() *PdfInfo {
	return &PdfInfo{}
}
