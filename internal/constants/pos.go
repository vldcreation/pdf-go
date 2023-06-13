package constants

const (
	// PosX is the X position of qr image
	PosX = 500
	// PosY is the Y position of qr image
	PosY = 790
	// Width is the width of qr image
	Width = 50
)

var (
	DEFAULT_QR_OPTS = struct {
		PosX  float64
		PosY  float64
		Width float64
	}{
		PosX:  PosX,
		PosY:  PosY,
		Width: Width,
	}
)
