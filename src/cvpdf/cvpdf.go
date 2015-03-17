package main

import (
	"flag"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
	"log/syslog"
	"strconv"
	"toml_parser"
	//	"domains"
	//	"path/filepath"
	//	"os"
)

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	bcv := toml_parser.Parse("/home/juno/git/go_cv/cv.toml")

	header := []string{"Item", "Duration", "Info"}

	pdf := gofpdf.New("P", "mm", "A4", "fonts")
	pdf.SetHeaderFunc(func() {
		pdf.Image("/home/juno/git/cv/src/assets/images/aleksander_mazurov.jpg", 10, 6, 30, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 15)
		pdf.Cell(80, 0, "")
		pdf.CellFormat(70, 10, "Aleksander Mazurov CV", "B", 0, "C", false, 0, "")
		pdf.Ln(20)
	})
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})
	pdf.AliasNbPages("")
	pdf.AddPage()

	w := []float64{35, 20, 115}
	wSum := 0.0
	for _, v := range w {
		wSum += v
	}

	// Color and font restoration

	for _, cv := range bcv.Cv {
		
		if cv.Name =="Cloud Computing" {
			
			pdf.AddPage()
		}
		
		pdf.SetFont("Arial", "B", 15)
		pdf.CellFormat(0, 10, cv.Name, "", 1, "", false, 0, "")

		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(128, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("Arial", "", 8)

		for j, str := range header {
			pdf.CellFormat(w[j], 7, str, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)

		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 0)

		fill := false

		for _, item := range cv.Item {

			pdf.CellFormat(w[0], 6, item.Title, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[1], 6, strconv.Itoa(item.Duration), "LR", 0, "R", fill, 0, "")
			
			pdf.CellFormat(w[2], 6, item.Extra, "LR", 0, "", fill, 0, "")
//			pdf.MultiCell(w[2], 6, item.Extra, "", "L", fill)
			pdf.Ln(-1)
			fill = !fill

		}

	}


	err = pdf.OutputFileAndClose("/home/juno/git/cv/src/mazurov_cv.pdf")
	if err == nil {
		fmt.Println("Successfully generated /home/juno/git/cv/src/mazurov_cv.pdf")
	} else {
		fmt.Println(err)
	}

}
