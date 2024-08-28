package pdfmaker

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/jung-kurt/gofpdf/v2"
	"gitlab.com/muallimah/course_service/internal/pkg/config"
)

func GenerateCertificate(name, course, issuedBy string, cf *config.Config) (*string,error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 24)
	pdf.Cell(40, 10, "Certificate of Completion")
	pdf.Ln(20)

	pdf.SetFont("Arial", "", 18)
	pdf.Cell(40, 10, "This is to certify that")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(40, 10, name)
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 18)
	pdf.Cell(40, 10, "has successfully completed the course")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(40, 10, course)
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 18)
	pdf.Cell(40, 10, "Issued by")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(40, 10, issuedBy)
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 18)
	pdf.Cell(40, 10, "Date")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 20)
    date := time.Now()
	pdf.Cell(40, 10, date.Format("January 2, 2006"))

	filename := fmt.Sprintf("%s_%s_certificate.pdf", name, course)
	filepath := filepath.Join(cf.MinioPath, filename)

	// Save the file in the specified directory
	err := pdf.OutputFileAndClose(filepath)
	if err != nil {
		return nil, err
	}

	return &filename, nil
}

// func main() {
// 	err := generateCertificate("John Doe", "Golang Backend Development", "Najot Ta`lim", time.Now())
// 	if err != nil {
// 		panic(err)
// 	}
// }
