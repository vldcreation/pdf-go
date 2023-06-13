package entity

import (
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type Qrcode struct {
	// Content of the qrcode
	text string
	// Name of the qrcode
	// Example: "qrcode.png"
	QrPath string
	// Logo of the qrcode
	// Example: "logo.png"
	// If empty, the qrcode will not have a logo
	QrLogoPath *string
	// qrcode.QRCode
	// If nil, the qrcode will be generated from default
	qr *qrcode.QRCode
	// standard.Writer
	// If nil, the qrcode will be generated from default
	w *standard.Writer
}

// NewQrcode create a new Qrcode
func NewQrcode(text string, qrPath string, qrLogoPath *string, qr *qrcode.QRCode, w *standard.Writer) *Qrcode {
	qrc := applyDefaultQr(text, qrPath, qrLogoPath)

	if qr != nil {
		qrc.setQR(qr)
	}

	if w != nil {
		qrc.setW(w)
	}

	return qrc
}

func (q *Qrcode) setQR(qr *qrcode.QRCode) {
	if qr == nil {
		qr, err := qrcode.NewWith(q.text)

		if err != nil {
			panic(err)
		}

		q.qr = qr

		return

	}

	q.qr = qr
}

func (q *Qrcode) QR() *qrcode.QRCode {
	// ensure q.qr doesn't nil to avoid nil pointer dereferences
	if q.qr == nil {
		q.setQR(nil)
	}

	return q.qr
}

func (q *Qrcode) setW(w *standard.Writer) {
	// Validates logo size as follows qrWidth >= 2*logoWidth && qrHeight >= 2*logoHeight
	// Instead of default expression qrWidth >= 5*logoWidth && qrHeight >= 5*logoHeight
	if w == nil {
		w, err := standard.New(
			q.QrPath,
		)
		if err != nil {
			panic(err)
		}

		if q.QrLogoPath != nil {
			// set logo
			w, _ = standard.New(
				q.QrPath,
				applyDefaultImageOpt(*q.QrLogoPath)...,
			)
		}

		q.w = w

		return
	}

	q.w = w
}

func applyDefaultImageOpt(qrLogoPath string) []standard.ImageOption {
	return []standard.ImageOption{
		standard.WithLogoImageFilePNG(qrLogoPath),
		standard.WithQRWidth(8),
		standard.WithBorderWidth(2),
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
	}
}

func (q *Qrcode) W() *standard.Writer {
	// ensure q.w not nil
	if q.w == nil {
		q.setW(nil)
	}
	return q.w
}

// applyDefault apply default options to the qrcode
// for more information, see https://github.com/yeqown/go-qrcode/blob/main/example/main.go
func applyDefaultQr(text string, qrPath string, qrLogoPath *string) *Qrcode {
	qr := &Qrcode{
		text:       text,
		QrPath:     qrPath,
		QrLogoPath: qrLogoPath,
	}

	qr.setW(nil)

	qr.setQR(nil)

	return qr
}
