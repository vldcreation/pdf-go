# github.com/vldcreation/privy-pdf-go

## How To Use
```
1. Install all depdencies
-- make clean
-- make install
2. See main.go for more detail example
3. All Use Case
- Unlock Pdf
  - func UnlockPDF(inputPath, outputPath string, password []byte)
- Extract Document
  - func ExtractFromDoc(inputPath, outputPath string, opt *convert.Options) (*entity.PDF, error)
- Generate Qr Code
  - func GenerateQRCode(qr *entity.Qrcode) error
- Load Pdf
  - func LoadPDF(inputPath string, opts *entity.PdfOpts) (*entity.PDF, error)
- Add Qr to Pdf
  - func AddQrCodeToPdf(inputPath string, outputPath string, qrPath string, qSt *entity.QrStrategy, qOpts *entity.QrOpts) error
- Extract Thumbnail
  - 
- Get PDF Info
  - 
- Set new Metadata (replacing)
  - github.com/vldcreation/privy-pdf-go/internal/usecase
see more detail in 
```