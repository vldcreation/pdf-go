package usecase

import (
	"image"
	"os"

	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
	"github.com/vldcreation/privy-pdf-go/internal/constants"
	"github.com/vldcreation/privy-pdf-go/internal/entity"
)

// Add image to a specific page of a PDF.  xPos and yPos define the lower left corner of the image location, and iwidth
// is the width of the image in PDF coordinates (height/width ratio is maintained).
func AddQrCodeToPdf(inputPath string, outputPath string, qrPath string, qSt *entity.QrStrategy, qOpts *entity.QrOpts) error {
	btR, err := os.Open(qrPath)
	if err != nil {
		panic(err)
	}

	defer btR.Close()

	qrCode, _, err := image.Decode(btR)

	// Read the input pdf file.
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	// Make a new PDF creator.
	c := creator.New()

	// Load the pages and add to creator.  Apply the QR code to the specified page.
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return err
		}

		err = c.AddPage(page)
		if err != nil {
			return err
		}

		if isSpecificPage(i, qSt.ApplyPos().Int()) || isAllPage(qSt.ApplyPos().Int()) || isLastPage(i, numPages, qSt.ApplyPos().Int()) {
			// Apply the QR code to the specified page or all pages if -1.
			img, err := c.NewImageFromGoImage(qrCode)
			if err != nil {
				return err
			}
			img.SetWidth(qOpts.Width())
			img.SetHeight(qOpts.Width())
			img.SetPos(qOpts.XPos(), qOpts.YPos())
			img.Scale(1.0, 1.0)
			err = c.Draw(img)
			if err != nil {
				return err
			}
		}
	}

	err = c.WriteToFile(outputPath)
	if err != nil {
		return err
	}

	return nil
}

func isSpecificPage(i, applyPos int) bool {
	return i+1 == applyPos
}

func isAllPage(applyPos int) bool {
	return applyPos == constants.ALL_PAGE
}

func isLastPage(i, numPages, applyPos int) bool {
	return i+1 == numPages && applyPos == constants.LAST_PAGE
}
