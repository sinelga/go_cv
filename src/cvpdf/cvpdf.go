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

	bconfig := toml_parser.Parse("/home/juno/git/go_cv/cv.toml")

	header := []string{"Item", "Duration", "Info"}
	jobplaceheader := []string{"Company", "Duration","Position","Details","Location","Country"}
	pdf := gofpdf.New("P", "mm", "A4", "fonts")


	pdf.SetHeaderFunc(func() {

		pdf.Image("/home/juno/git/cv/src/assets/images/aleksander_mazurov.jpg", 10, 6, 40, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "I", 10)
		pdf.SetX(110)
		pdf.CellFormat(70, 10, bconfig.Subtitle, "", 0, "C", false, 0, "")
		pdf.Ln(-1)
		pdf.SetFont("Arial", "B", 15)
		pdf.SetX(90)
		pdf.CellFormat(70, 10, bconfig.Maintitle, "", 0, "C", false, 0, "")
		pdf.Ln(-1)
		pdf.SetFont("Arial", "", 10)
		pdf.SetX(95)
		pdf.CellFormat(20, 6, "Phone:", "LRT", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "+358451202801", "LRT", 0, "", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(95)
		pdf.CellFormat(20, 6, "Email:", "LRT", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "support@mazurov.eu", "LRT", 0, "", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(95)
		pdf.CellFormat(20, 6, "Web:", "LRT", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "http://mazurov.eu", "LRT", 0, "", false, 0, "")
		pdf.Ln(-1)

		pdf.SetX(95)
		pdf.CellFormat(20, 6, "Skype:", "LRT", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "mazurovfi", "LRT", 0, "", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(95)
		pdf.CellFormat(20, 6, "Address:", "LRT", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "Hogberginkuja 1", "LRT", 0, "", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(95)
		pdf.CellFormat(20, 6, "", "LRB", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "10820 Lappohja Finland", "LRB", 0, "", false, 0, "")

		pdf.Ln(20)
	})
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})
	pdf.AliasNbPages("")
	pdf.AddPage()

	w := []float64{35, 15, 125}

	for _, cv := range bconfig.Cv {

		if cv.Name == "WebApplication Development" {

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

		for i, item := range cv.Item {

			if i == len(cv.Item)-1 {

				pdf.CellFormat(w[0], 6, item.Title, "LRB", 0, "", fill, 0, "")
				pdf.CellFormat(w[1], 6, strconv.Itoa(item.Duration)+" y.", "LRB", 0, "R", fill, 0, "")
				pdf.CellFormat(w[2], 6, item.Extra, "LRB", 0, "", fill, 0, "")

			} else {

				pdf.CellFormat(w[0], 6, item.Title, "LR", 0, "", fill, 0, "")
				pdf.CellFormat(w[1], 6, strconv.Itoa(item.Duration)+" y.", "LR", 0, "R", fill, 0, "")
				pdf.CellFormat(w[2], 6, item.Extra, "LR", 0, "", fill, 0, "")
			}
			pdf.Ln(-1)

			fill = !fill

		}

	}

	bjob := toml_parser.ParseWorkPlaces("/home/juno/git/go_cv/jobs.toml")

	pdf.SetHeaderFunc(func() {

		pdf.Image("/home/juno/git/cv/src/assets/images/aleksander_mazurov.jpg", 10, 6, 40, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "I", 10)
		pdf.SetX(110)
		pdf.CellFormat(70, 10, bjob.Subtitle, "", 0, "C", false, 0, "")
		pdf.Ln(-1)
		pdf.SetFont("Arial", "B", 15)
		pdf.SetX(90)
		pdf.CellFormat(70, 10, bjob.Maintitle, "", 0, "C", false, 0, "")
		pdf.Ln(-1)
		pdf.SetFont("Arial", "", 10)
		pdf.SetX(95)
		pdf.CellFormat(20, 6, "Phonelss:", "LRT", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "+358451202801", "LRT", 0, "", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(95)
		pdf.CellFormat(20, 6, "Email:", "LRT", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "support@mazurov.eu", "LRT", 0, "", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(95)
		pdf.CellFormat(20, 6, "Web:", "LRT", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "http://mazurov.eu", "LRT", 0, "", false, 0, "")
		pdf.Ln(-1)

		pdf.SetX(95)
		pdf.CellFormat(20, 6, "Skype:", "LRT", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "mazurovfi", "LRT", 0, "", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(95)
		pdf.CellFormat(20, 6, "Address:", "LRT", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "Hogberginkuja 1", "LRT", 0, "", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(95)
		pdf.CellFormat(20, 6, "", "LRB", 0, "", false, 0, "")
		pdf.CellFormat(65, 6, "10820 Lappohja Finland", "LRB", 0, "", false, 0, "")

		pdf.Ln(20)
	})

	pdf.AddPage()

	wjobplace := []float64{30,15,20,60,35,15}
	
	for _, job := range bjob.Jobs {

		pdf.SetFont("Arial", "B", 15)
		pdf.CellFormat(0, 10, job.Name, "", 1, "", false, 0, "")

		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(128, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("Arial", "", 8)

		for j, str := range jobplaceheader  {
			pdf.CellFormat(wjobplace[j], 7, str, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)

		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 0)

		fill := false

		for i, item := range job.Item {

			if i == len(job.Item)-1 {

				pdf.CellFormat(wjobplace[0], 6, item.Title, "LRB", 0, "", fill, 0, "")
				pdf.CellFormat(wjobplace[1], 6, item.Duration, "LRB", 0, "R", fill, 0, "")
				pdf.CellFormat(wjobplace[2], 6, item.Position, "LRB", 0, "", fill, 0, "")
				pdf.CellFormat(wjobplace[3], 6, item.Details, "LRB", 0, "", fill, 0, "")
				pdf.CellFormat(wjobplace[4], 6, item.Location, "LRB", 0, "", fill, 0, "")
				pdf.CellFormat(wjobplace[5], 6, item.Country, "LRB", 0, "", fill, 0, "")

			} else {

				pdf.CellFormat(wjobplace[0], 6, item.Title, "LR", 0, "", fill, 0, "")
				pdf.CellFormat(wjobplace[1], 6, item.Duration, "LR", 0, "R", fill, 0, "")
				pdf.CellFormat(wjobplace[2], 6, item.Position, "LR", 0, "", fill, 0, "")
				pdf.CellFormat(wjobplace[3], 6, item.Details, "LR", 0, "", fill, 0, "")
				pdf.CellFormat(wjobplace[4], 6, item.Location, "LR", 0, "", fill, 0, "")
				pdf.CellFormat(wjobplace[5], 6, item.Country, "LR", 0, "", fill, 0, "")
				
			}
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
