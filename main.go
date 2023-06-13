package main

import (
	"fmt"
	"log"

	unipdflicense "github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/model"
	"github.com/vldcreation/privy-pdf-go/internal/config"
	"github.com/vldcreation/privy-pdf-go/internal/constants"
	"github.com/vldcreation/privy-pdf-go/internal/entity"
	"github.com/vldcreation/privy-pdf-go/internal/usecase"
	"github.com/vldcreation/privy-pdf-go/internal/util"
)

var c = config.NewConfig(".")

func init() {
	err := unipdflicense.SetMeteredKey(c.GetUniDOcLicenseKey())
	if err != nil {
		fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
		fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
		fmt.Printf("If you don't have one - Grab one in the Free Tier at https://cloud.unidoc.io\n")
		panic(err)
	}

	// This example requires both for unioffice and unipdf.
	// err = license.SetMeteredKey(c.GetUniDOcLicenseKey())
	// if err != nil {
	// 	fmt.Printf("ERROR: Failed to set metered key: %v\n", err)
	// 	fmt.Printf("Make sure to get a valid key from https://cloud.unidoc.io\n")
	// 	fmt.Printf("If you don't have one - Grab one in the Free Tier at https://cloud.unidoc.io\n")
	// 	panic(err)
	// }
}

func main() {
	fileInput := "sample_protected.pdf"
	fileOutput := "sample_protected.pdf"

	// Unlock protected pdf
	util.SetUnlockFilename(&fileOutput)
	fileInputPath, fileOutputPath := getInputOutputPath(fileInput, fileOutput)

	readerOpts := model.NewReaderOpts()

	readerOpts.Password = "test"

	pdf, err := usecase.LoadPDF(fileInputPath, entity.NewPdfOpts(
		true,
		readerOpts,
		&fileOutputPath,
	))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pdf.GetPdfReader().GetNumPages())

	qrPath := constants.OUTPUT_QRCODE_TEST_DATA + "simple.png"
	qrLogoPath := constants.INPUT_QRCODE_TEST_DATA + "logo.png"

	if err := usecase.GenerateQRCode(
		entity.NewQrcode(
			c.AppPrivyVerifyDomain+util.GenerateRandoomUUID(),
			qrPath,
			&qrLogoPath,
			nil,
			nil)); err != nil {
		panic(err)
	}

	qrPdfOutputPath := constants.OUTPUT_TEST_DATA + "simple_qr.pdf"

	err = usecase.AddQrCodeToPdf(
		fileOutputPath,
		qrPdfOutputPath,
		qrPath,
		entity.NewQrStrategy(entity.ToPagePos(constants.LAST_PAGE)),
		entity.NewQrOpts(
			constants.DEFAULT_QR_OPTS.PosX,
			constants.DEFAULT_QR_OPTS.PosY,
			constants.DEFAULT_QR_OPTS.Width),
	)

	if err != nil {
		panic(err)
	}

	// Extract Thumbnail
	pdf, err = usecase.LoadPDF(qrPdfOutputPath, nil)
	if err != nil {
		panic(err)
	}

	if err := pdf.ThumbNail(constants.OUTPUT_TEST_DATA + "thumb.pdf"); err != nil {
		panic(err)
	}

	pdfInfo, err := pdf.GetPdfInfo()
	if err != nil {
		panic(err)
	}

	fmt.Printf("pdf info: %+v\n", pdfInfo)

	// Replace Metadata

	// init instance pdfInfo
	newPdfInfo := entity.NewPdfInfo()
	newPdfInfo.Author = "vldcreation"
	newPdfInfo.Title = "vldcreation"
	newPdfInfo.Subject = "vldcreation"
	newPdfInfo.Keywords = "vldcreation"

	newPdfInfo.Extras = pdfInfo.Extras

	if err := pdf.SetPdfInfo(
		constants.OUTPUT_TEST_DATA+"new_simple_qr.pdf",
		pdfInfo,
		map[string]string{
			"Producer": "vldcreation",
		},
	); err != nil {
		panic(err)
	}

}

func getInputOutputPath(inputpath, outpath string) (string, string) {
	return constants.INPUT_TEST_DATA + inputpath, constants.OUTPUT_TEST_DATA + outpath
}
