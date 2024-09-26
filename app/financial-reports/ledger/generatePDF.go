package ledger

import (
	"fmt"
	"log"

	// accountingentry "mazza-backend/accountingEntry"
	"mazza/app/utils"
	ent "mazza/ent/generated"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type details struct {
	Title        string
	Subtitle     string
	Keywords     string
	Subject      string
	ColumnWidths []int
	Company      ent.Company
	TableHeader  []string
	TableBody    [][]string
	TableFooter  []string
}

func generatePDFDoc(data *[]ledgerStruct, company *ent.Company, startDate time.Time, endDate time.Time) (file *string, size float64, err error) {
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		tz = time.UTC
	}
	start := startDate.In(tz).Format("02-Jan-2006")
	end := endDate.In(tz).Format("02-Jan-2006")
	subtitle := "Período: " + start + " - " + end
	tableHeader := []string{"Data", "Conta", "Nome", "Descrição", "Débito", "Crédito"}
	columnWidths := []int{2, 1, 3, 2, 2, 2}
	keywords := "ledger,diario,lancamentos"
	subject := "invoice/fatura"
	tableBody, tableFooter := getTableRows(data)

	details := details{
		Title:        "Diário",
		Subtitle:     subtitle,
		Keywords:     keywords,
		Subject:      subject,
		ColumnWidths: columnWidths,
		Company:      *company,
		TableHeader:  tableHeader,
		TableBody:    tableBody,
		TableFooter:  tableFooter,
	}
	engine := getMaroto(&details)
	document, err := engine.Generate()
	if err != nil {
		log.Fatal(err.Error())
		return nil, 0, err
	}

	err = document.Save("billing.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
	size = document.GetReport().SizeMetric.Size.Value * 1024
	fmt.Println("size:", size, "byte")

	doc := document.GetBase64()
	return &doc, size, nil
}

func getMaroto(details *details) core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(25).
		WithTopMargin(15).
		WithRightMargin(25).
		WithCreationDate(time.Now()).
		WithAuthor("Mazza", true).
		WithCreator("Mazza", true).
		WithSubject(details.Subject, true).
		WithKeywords(details.Keywords, true).
		WithTitle(details.Title, true).
		WithOrientation(orientation.Horizontal).
		Build()

	// darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	doc := maroto.NewMetricsDecorator(mrt)

	err := doc.RegisterHeader(getPageHeader(details.Company.Name, *details.Company.TaxId))
	if err != nil {
		log.Fatal(err.Error())
	}

	err = doc.RegisterFooter(getPageFooter())
	if err != nil {
		log.Fatal(err.Error())
	}

	// Title
	doc.AddRows(getCompanyDetails(&details.Company))
	doc.AddRow(10)
	doc.AddRows(getTitleAndSubtitle(&details.Title, &details.Subtitle)...)

	// Body
	doc.AddRows(getTable(&details.ColumnWidths, &details.TableHeader, &details.TableBody, &details.TableFooter)...)

	// Notes
	doc.AddRow(15)

	// Don't remove the line below. The doc gets corrupted if this line is removed
	doc.AddRow(2, col.New(6).Add(code.NewBar("5123.151231", props.Barcode{})))

	return doc
}

func getPageHeader(companyName string, taxId string) core.Row {
	value := companyName + " | NUIT: " + taxId
	return row.New(15).Add(
		col.New(12).Add(text.New(value, props.Text{Top: 1, Size: 9, Align: align.Left})),
	)
}

func getPageFooter() core.Row {
	year := strconv.Itoa(time.Now().Year())
	link := "https://mazza.jp"
	return row.New(20).Add(
		col.New(12).Add(
			text.New("Gerado por Mazza - © "+year, props.Text{Top: 13, Size: 8, Align: align.Left}),
			text.New("www.mazza.jp", props.Text{Top: 16, Size: 8, Align: align.Left, Hyperlink: &link}),
		),
	)
}

func getCompanyDetails(company *ent.Company) core.Row {
	taxID := fmt.Sprint("NUIT: ", company.TaxId)
	city := fmt.Sprint(company.City, ", ", company.Country)
	phone := fmt.Sprint("Tel: ", utils.GetValue(company.Phone, "--"))
	email := fmt.Sprint("Email: ", utils.GetValue(company.Email, "--"))
	return row.New(20).Add(
		col.New(8),
		col.New(4).Add(
			text.New(company.Name, props.Text{Top: 1, Size: 12, Align: align.Right, Color: utils.Colors.Red, Style: fontstyle.Bold}),
			text.New(taxID, props.Text{Top: 6, Size: 11, Align: align.Right, Color: utils.Colors.Red}),
			text.New(city, props.Text{Top: 10, Size: 11, Align: align.Right, Color: utils.Colors.Red}),
			text.New(phone, props.Text{Top: 14, Size: 11, Align: align.Right, Color: utils.Colors.Red}),
			text.New(email, props.Text{Top: 18, Size: 11, Align: align.Right, Color: utils.Colors.Red}),
		),
	)
}

func getTitleAndSubtitle(title *string, subtitle *string) []core.Row {
	content := []core.Row{
		text.NewRow(12, *title, props.Text{Size: 20, Style: fontstyle.Bold, Align: align.Center, VerticalPadding: 10}),
		text.NewRow(12, *subtitle, props.Text{Size: 14, Style: fontstyle.Bold, Align: align.Center, VerticalPadding: 10}),
	}

	return content
}

// Table header, body, footer
func getTable(columnWidths *[]int, tableHeader *[]string, body *[][]string, totals *[]string) []core.Row {
	// Add table header
	cols := []core.Col{}
	for i, value := range *tableHeader {
		textAlign := align.Left
		if i == 0 || i > 3 {
			textAlign = align.Center
		}
		cols = append(cols, text.NewCol((*columnWidths)[i], value, props.Text{Size: 11, Align: textAlign, Color: utils.Colors.White, Style: fontstyle.Bold}))
	}
	rows := []core.Row{row.New(7).Add(cols...).WithStyle(&props.Cell{BackgroundColor: utils.Colors.Red})}

	// Add table body
	for i, entry := range *body {
		cols = []core.Col{}
		for i, value := range entry {
			textAlign := align.Left
			if i > 3 {
				textAlign = align.Right
			}
			cols = append(cols, text.NewCol((*columnWidths)[i], value, props.Text{Size: 11, Align: textAlign}))
		}
		newRow := row.New(6).Add(cols...)
		if i%2 == 0 {
			gray := utils.Colors.Gray
			newRow.WithStyle(&props.Cell{BackgroundColor: gray})
		}
		rows = append(rows, newRow)
	}

	// Add table footer
	cols = []core.Col{}
	for i, value := range *totals {
		textAlign := align.Left
		if i == 0 {
			textAlign = align.Center
		} else if i == 0 || i > 3 {
			textAlign = align.Right
		}
		cols = append(cols, text.NewCol((*columnWidths)[i], value, props.Text{Size: 11, Align: textAlign, Style: fontstyle.Bold}))
	}
	newRow := row.New(8).Add(cols...).WithStyle(&props.Cell{
		BorderType: border.Top,
	})
	rows = append(rows, newRow)

	return rows
}

func getTableRows(data *[]ledgerStruct) (body [][]string, footer []string) {
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("location error:", err)
		tz = time.UTC
	}
	totalDebit := float64(0)
	totalCredit := float64(0)
	for i, entry := range *data {
		date := ""
		currentDate := entry.Date.In(tz).Format("02-Jan-2006")
		if i == 0 {
			date = currentDate
		} else {
			previousDate := (*data)[i-1].Date.In(tz).Format("02-Jan-2006")

			if previousDate != currentDate {
				date = currentDate
			}
		}

		description := fmt.Sprint("", entry.Description)
		if len(description) > 20 {
			description = description[:20]
		}

		debit, credit := utils.ToCurrencyStr(entry.Debit), utils.ToCurrencyStr(entry.Credit)
		if entry.Debit == 0 {
			debit = ""
		}
		if entry.Credit == 0 {
			credit = ""
		}

		if isDebitable(entry.AccountType) && entry.Credit < 0 {
			credit = utils.ToCurrencyStr(-entry.Credit)
		} else if !isDebitable(entry.AccountType) && entry.Debit < 0 {
			debit = utils.ToCurrencyStr(-entry.Debit)
		}

		row := []string{date, entry.Account, entry.Label, description, debit, credit}
		totalDebit += entry.Debit
		totalCredit += entry.Credit
		body = append(body, row)
	}
	footer = []string{"Total", "", "", "", utils.ToCurrencyStr(totalDebit), utils.ToCurrencyStr(totalCredit)}
	return body, footer
}
