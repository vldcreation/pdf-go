package usecase

import (
	"bytes"
	"fmt"
	"log"

	"github.com/unidoc/unipdf/v3/model"
	"github.com/vldcreation/privy-pdf-go/internal/util"
)

func UnlockPDF(inputPath, outputPath string, password []byte) error {
	bt, err := util.LoadFile(inputPath)
	if err != nil {
		log.Fatalf("error while load file %+v\n", err)
		return err
	}

	// read pdf using no opts
	// we assume that the pdf is encrypted
	// so just pass password as parameter instead option parameter
	rs := bytes.NewReader(bt)
	pdf, err := model.NewPdfReader(rs)
	if err != nil {
		log.Fatalf("error while read pdf %+v\n", err)
		return err
	}

	isEncrypted, err := pdf.IsEncrypted()
	if err != nil {
		log.Fatalf("error while check pdf is encrypted %+v\n", err)
		return err
	}

	if isEncrypted {
		auth, err := pdf.Decrypt(password)
		if err != nil {
			return err
		}
		if !auth {
			return fmt.Errorf("Wrong password")
		}
	}

	pdfWriter, err := pdf.ToWriter(nil)
	if err != nil {
		return err
	}

	// Write to file.
	return pdfWriter.WriteToFile(outputPath)
}
