package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/minhnhatspk/chromedp-screenshot-pdf/models"
	"io/ioutil"
	"log"
	"path/filepath"
	"text/template"
	"time"
)

func main() {
	pdfTemplate := &models.PdfTemplate{
		Hospital: models.Hospital{
			Name:    "The CMS Center",
			Address: "89 Ngo quyen, hiep phu",
			Phone:   "0983746523",
		},
		Prescription: models.Prescription{
			Date:     time.Now().Format(time.RFC822),
			OrderNom: "dh24320947",
		},
		Patient: models.Patient{
			Name:    "Peter Parker",
			Address: "134 Le Hong Phong",
			Phone:   "0933746523",
		},
		Doctor: models.Doctor{
			Name: "John Thomson",
		},
		Drug: models.Drug{
			Name:    "Panadol Z 500MG CAP 1",
			Ammount: 10,
			Days:    9,
			Refill:  3,
			Price:   45.4,
		},
		Pay:                 23.5,
		PharmacistSignature: "byte array",
	}

	teml, err := template.New("pdf-template.html").ParseFiles("./pdf-template.html")
	if err != nil {
		log.Fatal(err)
	}

	templateBuf := new(bytes.Buffer)

	if err := teml.Execute(templateBuf, pdfTemplate); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("prescription-templ.html", []byte(templateBuf.String()), 0644); err != nil {
		log.Fatal(err)
	}

	//// create context
	//opts := []chromedp.ExecAllocatorOption{
	//	chromedp.Headless,
	//	chromedp.DisableGPU,
	//	chromedp.NoSandbox,
	//}
	//
	//allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	//defer cancel()
	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	var pdfBuffer []byte

	absTemplatePath, err := filepath.Abs("./prescription-templ.html")

	fmt.Println(absTemplatePath)

	if err != nil {
		log.Fatal(err)
	}

	if err := chromedp.Run(taskCtx, pdfScreenshot("file://"+absTemplatePath, "#main", &pdfBuffer)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("prescription.pdf", pdfBuffer, 0644); err != nil {
		log.Fatal(err)
	}
}

func pdfScreenshot(url string, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
