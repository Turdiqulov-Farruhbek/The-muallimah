package pdfmaker

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/jung-kurt/gofpdf/v2"
	"github.com/skip2/go-qrcode"
	"gitlab.com/muallimah/course_service/internal/pkg/config"
)

func GenerateCertificate(name, course, issuedBy string, cf *config.Config) (*string, error) {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFillColor(255, 252, 240) // Light cream background
	pdf.Rect(0, 0, 297, 210, "F")

	drawDecorativeBorder(pdf)

	pdf.SetFont("Times", "B", 40)
	pdf.SetTextColor(139, 69, 19) // Saddle brown
	pdf.SetY(30)
	pdf.CellFormat(297, 20, "Certificate of Excellence", "", 0, "C", false, 0, "")

	pdf.SetFont("Caveat", "", 40)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetY(60)
	pdf.CellFormat(297, 20, name, "", 0, "C", false, 0, "")

	pdf.SetFont("Times", "I", 18)
	pdf.SetY(80)
	pdf.CellFormat(297, 15, "has successfully completed the course", "", 0, "C", false, 0, "")
	pdf.SetY(95)
	pdf.SetFont("Times", "B", 22)
	pdf.CellFormat(297, 15, course, "", 0, "C", false, 0, "")

	pdf.SetFont("Times", "I", 14)
	pdf.SetY(120)
	pdf.CellFormat(297, 10, "We commend your dedication, hard work, and outstanding achievement.", "", 0, "C", false, 0, "")
	pdf.SetY(130)
	pdf.CellFormat(297, 10, "May this accomplishment be the beginning of continued success in your future endeavors.", "", 0, "C", false, 0, "")

	pdf.SetFont("Times", "", 12)
	pdf.SetY(160)
	pdf.CellFormat(99, 10, fmt.Sprintf("Awarded on: %s", time.Now().Format("January 2, 2006")), "", 0, "R", false, 0, "")
	pdf.CellFormat(99, 10, "", "", 0, "C", false, 0, "")
	pdf.CellFormat(99, 10, issuedBy, "", 0, "L", false, 0, "")

	pdf.SetY(170)
	pdf.CellFormat(99, 10, "", "", 0, "R", false, 0, "")
	pdf.CellFormat(99, 10, "", "", 0, "C", false, 0, "")
	pdf.CellFormat(99, 10, "Instructor", "", 0, "L", false, 0, "")

	qr, err := generateQRCode(name)
	if err != nil {
		return nil, err
	}
	pdf.Image(qr, 250, 160, 30, 30, false, "", 0, "")

	filename := fmt.Sprintf("%s_%s_certificate.pdf", name, course)
	filePath := filepath.Join(cf.MinioPath, filename)
	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		return nil, err
	}
	
	return &filename, nil
}

func drawDecorativeBorder(pdf *gofpdf.Fpdf) {
	pdf.SetLineWidth(0.5)
	pdf.SetDrawColor(139, 69, 19)

	pdf.Rect(5, 5, 287, 200, "D")

	pdf.Rect(10, 10, 277, 190, "D")

	cornerSize := 20.0
	pdf.Line(5, 5, 5+cornerSize, 5)
	pdf.Line(5, 5, 5, 5+cornerSize)
	pdf.Line(292-cornerSize, 5, 292, 5)
	pdf.Line(292, 5, 292, 5+cornerSize)
	pdf.Line(5, 205-cornerSize, 5, 205)
	pdf.Line(5, 205, 5+cornerSize, 205)
	pdf.Line(292-cornerSize, 205, 292, 205)
	pdf.Line(292, 205-cornerSize, 292, 205)
}

func generateQRCode(data string) (string, error) {
	filename := "qr_code.png"
	err := qrcode.WriteFile(data, qrcode.Medium, 256, filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}

// func main() {
// 	err := generateCertificate("John Doe", "Golang Backend Development", "Najot Ta`lim", time.Now())
// 	if err != nil {
// 		panic(err)
// 	}
// }
