package usecase

import "github.com/vldcreation/privy-pdf-go/internal/entity"

func GenerateQRCode(qr *entity.Qrcode) error {
	defer qr.W().Close()
	return qr.QR().Save(qr.W())
}
