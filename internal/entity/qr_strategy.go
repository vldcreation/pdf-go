package entity

// PagePos to define the position of the QR code in the page
// @default LAST_PAGE
type PagePos int

// Do convertion from int to PagePos
// @param p int - the position of the QR code in the page
// e.g. 1 - first page, 0 - last page, -1 all page, otherwise should be specific page.
func ToPagePos(p int) PagePos {
	return PagePos(p)
}

func (p PagePos) Pointer() *PagePos {
	return &p
}

func (p PagePos) Int() int {
	return int(p)
}

func (p PagePos) Int64() int64 {
	return int64(p)
}

type QrStrategy struct {
	// applyPos to define the position of the QR code in the page
	applyPos PagePos
}

// NewQrStrategy to create a new QrStrategy
func NewQrStrategy(applyPos PagePos) *QrStrategy {
	return &QrStrategy{applyPos: applyPos}
}

func (s *QrStrategy) ApplyPos() PagePos {
	return s.applyPos
}
