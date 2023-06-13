package entity

type QrOpts struct {
	// XPos is the x position of the QR code
	// relative to the left edge of the page.
	xPos float64
	// YPos is the y position of the QR code
	// relative to the bottom edge of the page.
	yPos float64
	// Width is the width of the QR code.
	// The height is same with width.
	width float64
}

func NewQrOpts(xPos, yPos, width float64) *QrOpts {
	return &QrOpts{
		xPos:  xPos,
		yPos:  yPos,
		width: width,
	}
}

func (q *QrOpts) XPos() float64 {
	return q.xPos
}

func (q *QrOpts) YPos() float64 {
	return q.yPos
}

func (q *QrOpts) Width() float64 {
	return q.width
}
