package inventory

import (
	"fmt"
	"log"
	"mazza/app/utils"
	"mazza/mazza/generated/model"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// var ExampleData = invoice{
// 	Date:     time.Now().String(),
// 	Filename: "invoice",
// 	Number:   "564",
// 	Title:    "Invoice no.342",
// 	IssuerDetails: invoiceIssuer{
// 		Name:    "Company Name, Ltd",
// 		TaxID:   "398298392",
// 		Address: "Av. 25 de Julho, nr. 434",
// 		City:    "Cidade",
// 		Country: "Pais",
// 		Phone:   "55 024 123-2323",
// 		Email:   "dejdieo@mail.com",
// 	},
// 	CustomerDetails: invoiceCustomer{
// 		Name:    "Customer Name",
// 		TaxID:   "89585489384",
// 		Address: "Av. 25 de Julho, nr. 434",
// 		City:    "Cidade",
// 		Country: "Pais",
// 		Phone:   "55 024 123-2323",
// 		Email:   "dejdieo@mail.com",
// 	},
// 	Body: [][]string{
// 		{"Swamp SwamppSwampSwp", "128787", "12,00", "R$ 4,00"},
// 		{"Sorin, A Planeswalker", "5", "4,00", "R$ 90,00"},
// 		{"Tassa", "1", "4,00", "R$ 30,00"},
// 		{"Skinrender", "3", "4,00", "R$ 9,00"},
// 	},
// 	Totals: invoiceTotals{
// 		Subtotal: "4598",
// 		VATRate:  "16%",
// 		VAT:      "594",
// 		Total:    "5199",
// 	},
// }

func GenerateInvoicePDF(details *model.Invoice) (file []byte, size float64, err error) {
	title := "Fatura n." + details.Number
	details.Title = &title
	details.Keywords = "invoice,sales,fatura,vendas"
	return __generatePDF(details, true)
}

func GenerateSalesQuotationPDF(details *model.Invoice) (file []byte, size float64, err error) {
	details.Keywords = "sales quotation,quotacao de vendas,quotação"
	return __generatePDF(details, false)
}

func __generatePDF(details *model.Invoice, isInvoice bool) (file []byte, size float64, err error) {
	fmt.Println("due:", details.PaymentDetails)
	m := getMaroto(details, isInvoice)
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
		return nil, 0, err
	}

	// err = document.Save("billing.pdf")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	size = document.GetReport().SizeMetric.Size.Value * 1024
	scale := document.GetReport().SizeMetric.Size.Scale
	fmt.Println("size:", size, scale)

	return document.GetBytes(), size, nil
}

func getMaroto(details *model.Invoice, isInvoice bool) core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(20).
		WithTopMargin(15).
		WithRightMargin(20).
		WithCreationDate(time.Now()).
		WithAuthor("Mazza", true).
		WithCreator("Mazza (www.mazza.jp)", true).
		WithSubject("invoice/fatura", true).
		WithKeywords(details.Keywords, true).
		WithTitle(*details.Title, true).
		Build()

	// darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	doc := maroto.NewMetricsDecorator(mrt)

	err := doc.RegisterHeader(getPageHeader(details.IssuerDetails))
	if err != nil {
		log.Fatal(err.Error())
	}

	notes := ""
	err = doc.RegisterFooter(append(getNotes(details.PaymentDetails, &notes), getPageFooter())...)
	if err != nil {
		log.Fatal(err.Error())
	}

	doc.AddRows(text.NewRow(15, *details.Title, props.Text{Size: 25, Style: fontstyle.Bold, Align: align.Left, VerticalPadding: 10}))
	doc.AddRows(getIssuerAndDocDetails(details.IssuerDetails, &details.Date, &details.PaymentDetails.DueDate, &isInvoice))
	doc.AddRows(getCustomerDetails(details.CustomerDetails))
	doc.AddRow(15)

	// Body
	doc.AddRows(getTransactions(&details.Body, details.Totals)...)

	// Don't remove the line below. The doc gets corrupted if this line is removed
	// doc.AddRow(0.01, code.NewBarCol(1, "5123.151231"))

	return doc
}

func getPageHeader(issuer *model.InvoiceIssuer) core.Row {
	value := issuer.Name + " | NUIT: " + issuer.TaxID
	overhead := row.New(10).Add(col.New(12).Add(text.New(value, props.Text{Top: 1, Size: 9, Align: align.Left})))
	return overhead
}

func getIssuerAndDocDetails(issuer *model.InvoiceIssuer, issueDate *string, dueDate *string, isInvoice *bool) core.Row {
	var due string
	if *isInvoice {
		due = "Vencimento: "
	} else {
		due = "Válido até: "
	}
	return row.New(35).Add(
		col.New(8).Add(
			text.New(issuer.Name, props.Text{Top: 1, Size: 13, Align: align.Left, Style: fontstyle.Bold}),
			text.New("NUIT: "+issuer.TaxID, props.Text{Top: 7, Size: 12, Align: align.Left}),
			text.New(issuer.City+", "+issuer.Country, props.Text{Top: 12, Size: 12, Align: align.Left}),
			text.New("Tel: "+issuer.Phone, props.Text{Top: 17, Size: 12, Align: align.Left}),
			text.New("Email: "+issuer.Email, props.Text{Top: 23, Size: 12, Align: align.Left}),
		),
		col.New(2).Add(
			text.New("Data: ", props.Text{Top: 1, Size: 13, Align: align.Left}),
			text.New(due, props.Text{Top: 7, Size: 12, Align: align.Left}),
		),
		col.New(2).Add(
			text.New(*issueDate, props.Text{Top: 1, Size: 13, Align: align.Right, Style: fontstyle.Bold}),
			text.New(*dueDate, props.Text{Top: 7, Size: 12, Align: align.Right}),
		),
	)
}

func getCustomerDetails(customer *model.InvoiceCustomer) core.Row {
	red := utils.Colors.Red
	return row.New(30).Add(
		col.New(8).Add(
			text.New("CLIENTE:", props.Text{Top: 1, Size: 12, Align: align.Left, Style: fontstyle.Bold, Color: red}),
			text.New(customer.Name, props.Text{Top: 7, Size: 13, Align: align.Left, Style: fontstyle.Bold}),
			text.New("NUIT: "+customer.TaxID, props.Text{Top: 13, Size: 12, Align: align.Left}),
			text.New(customer.City+", "+customer.Country, props.Text{Top: 18, Size: 12, Align: align.Left}),
			text.New("Tel: "+customer.Phone, props.Text{Top: 24, Size: 12, Align: align.Left}),
			text.New("Email: "+customer.Email, props.Text{Top: 30, Size: 12, Align: align.Left}),
		),
		// col.New(4).Add(
		// 	text.New(issuer.Name, props.Text{Top: 1, Size: 13, Align: align.Right, Style: fontstyle.Bold, Color: red}),
		// 	text.New("NUIT: "+issuer.TaxID, props.Text{Top: 7, Size: 12, Align: align.Right, Color: red}),
		// 	text.New(issuer.City+", "+issuer.Country, props.Text{Top: 12, Size: 12, Align: align.Right, Color: red}),
		// 	text.New("Tel: "+issuer.Phone, props.Text{Top: 17, Size: 12, Align: align.Right, Color: red}),
		// 	text.New("Email: "+issuer.Email, props.Text{Top: 23, Size: 12, Align: align.Right, Color: red}),
		// ),
	)
}

func getNotes(info *model.PaymentDetails, notes *string) []core.Row {
	return []core.Row{
		row.New(15).WithStyle(&props.Cell{BorderType: border.Bottom, BorderThickness: 0.5}),
		row.New(2),
		row.New(7).Add(text.NewCol(6, "Termos e condiçōes", props.Text{Size: 13})),
		row.New(5).WithStyle(&props.Cell{BorderType: border.Top, BorderThickness: 0.5}),
		row.New(22).Add(
			col.New(22).Add(
				text.New(*notes, props.Text{Top: 1, Size: 12, Align: align.Left}),
				text.New("Banco: "+info.BankName, props.Text{Top: 10, Size: 12, Align: align.Left}),
				text.New("NIB: "+info.Iban, props.Text{Top: 16, Size: 12, Align: align.Left}),
			),
		).WithStyle(&props.Cell{BorderType: border.Bottom, BorderThickness: 0.5}),
	}
}

func getPageFooter() core.Row {
	year := strconv.Itoa(time.Now().Year())
	link := "https://mazza.jp"

	platformSignature := row.New(15).Add(
		col.New(8).Add(
			text.New("Gerado por Mazza - © "+year, props.Text{Top: 6, Size: 11, Align: align.Left}),
			text.New("www.mazza.jp", props.Text{Top: 10, Size: 11, Align: align.Left, Hyperlink: &link}),
		),
	)
	return platformSignature
}

// Body
func getTransactions(body *[][]string, totals *model.InvoiceTotals) []core.Row {
	white := utils.Colors.White
	red := utils.Colors.Red
	rows := []core.Row{
		row.New(7).Add(
			text.NewCol(6, " Produto", props.Text{Size: 12, Align: align.Left, Color: white, Style: fontstyle.Bold}),
			text.NewCol(1, "Qtd", props.Text{Size: 12, Align: align.Left, Color: white, Style: fontstyle.Bold}),
			text.NewCol(2, "Preço", props.Text{Size: 12, Align: align.Right, Color: white, Style: fontstyle.Bold}),
			text.NewCol(3, "Subtotal ", props.Text{Size: 12, Align: align.Right, Color: white, Style: fontstyle.Bold}),
		).WithStyle(&props.Cell{BorderType: border.Bottom, BackgroundColor: red}),
	}

	contentsRow := []core.Row{row.New(1)}

	for i, content := range *body {
		row := row.New(7).Add(
			text.NewCol(6, content[0], props.Text{Size: 12, Align: align.Left, BreakLineStrategy: breakline.DashStrategy}),
			text.NewCol(1, content[1], props.Text{Size: 12, Align: align.Left}),
			text.NewCol(2, content[2], props.Text{Size: 12, Align: align.Right}),
			text.NewCol(3, content[3], props.Text{Size: 12, Align: align.Right}),
		)
		if i%2 == 0 {
			gray := utils.Colors.Gray
			row.WithStyle(&props.Cell{BackgroundColor: gray})
		}
		if i == len(*body)-1 {
			row.WithStyle(&props.Cell{BorderType: border.Bottom, BorderThickness: 0.5})
		}
		contentsRow = append(contentsRow, row)
	}

	rows = append(rows, contentsRow...)
	appendTotalRows(&rows, totals.Subtotal, totals.VatRate, totals.Vat, totals.Total)

	return rows
}

func appendTotalRows(rows *[]core.Row, subtotal string, vatRate string, vat string, total string) {
	*rows = append(*rows, row.New(6).Add(
		col.New(7),
		text.NewCol(2, "Subtotal:", props.Text{Top: 5, Size: 12, Align: align.Right}),
		text.NewCol(3, subtotal, props.Text{Top: 5, Size: 12, Style: fontstyle.Bold, Align: align.Right}),
	))
	*rows = append(*rows, row.New(6).Add(
		col.New(7),
		text.NewCol(2, "IVA ("+vatRate+")", props.Text{Top: 5, Size: 12, Align: align.Right}),
		text.NewCol(3, vat, props.Text{Top: 5, Size: 12, Style: fontstyle.Bold, Align: align.Right}),
	))
	*rows = append(*rows, row.New(6).Add(
		col.New(7),
		text.NewCol(2, "Total:", props.Text{Top: 5, Size: 15, Align: align.Right}),
		text.NewCol(3, total, props.Text{Top: 5, Size: 15, Style: fontstyle.Bold, Align: align.Right}),
	))
	*rows = append(*rows, row.New(15))
}
