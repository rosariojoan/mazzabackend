package balancesheet

import (
	"fmt"
	"log"
	"mazza/app/utils"
	ent "mazza/ent/generated"
	"mazza/mazza/generated/model"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type details struct {
	Title        string
	Subtitle     string
	Keywords     string
	Subject      string
	ColumnWidths []int
	Company      *ent.Company
	TableHeader  []string
	Data         *model.BalanceSheetOuput
}

func generatePDFDoc(data *model.BalanceSheetOuput, company *ent.Company, startDate time.Time, endDate time.Time) (file *string, size float64, err error) {
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		tz = time.UTC
	}
	start := startDate.In(tz).Format("02-Jan-2006")
	end := endDate.In(tz).Format("02-Jan-2006")
	subtitle := "Período: " + start + " - " + end
	tableHeader := []string{"Conta", "Rubrica", "Valor"}
	columnWidths := []int{2, 7, 3}
	keywords := "balance sheet,balanco patrimonial"
	subject := "balance sheet/balanco patrimonial"

	details := details{
		Title:        "Balanço Patrimonial",
		Subtitle:     subtitle,
		Keywords:     keywords,
		Subject:      subject,
		ColumnWidths: columnWidths,
		Company:      company,
		TableHeader:  tableHeader,
		Data:         data,
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
	// size = document.GetReport().SizeMetric.Size.Value * 1024
	// fmt.Println("size:", size, "byte")
	doc := document.GetBase64()
	return &doc, size, nil
}

func getMaroto(details *details) core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(15).
		WithTopMargin(15).
		WithRightMargin(15).
		WithCreationDate(time.Now()).
		WithAuthor("Mazza", true).
		WithCreator("Mazza", true).
		WithSubject(details.Subject, true).
		WithKeywords(details.Keywords, true).
		WithTitle(details.Title, true).
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
	doc.AddRows(getCompanyDetails(details.Company))
	doc.AddRow(10)
	doc.AddRows(getTitleAndSubtitle(&details.Title, &details.Subtitle)...)

	// Body
	doc.AddRows(getTable(&details.ColumnWidths, &details.TableHeader, details.Data)...)

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

func getRow(item *model.ReportRowItem, columnWidths *[]int) (newRow core.Row) {
	return row.New(6).Add(
		text.NewCol((*columnWidths)[0], item.Account, props.Text{Size: 12, Align: align.Left}),
		text.NewCol((*columnWidths)[1], item.Label, props.Text{Size: 12, Align: align.Left}),
		text.NewCol((*columnWidths)[2], utils.ToCurrencyStr(item.Value), props.Text{Size: 11, Align: align.Right}),
	)
}

func getTotalRow(item *model.ReportRowItem, columnWidths *[]int) (newRow core.Row) {
	return row.New(6).Add(
		text.NewCol((*columnWidths)[0], item.Account, props.Text{Size: 12, Align: align.Left, Style: fontstyle.Bold}),
		text.NewCol((*columnWidths)[1], item.Label, props.Text{Size: 12, Align: align.Left, Style: fontstyle.Bold}),
		text.NewCol((*columnWidths)[2], utils.ToCurrencyStr(item.Value), props.Text{Size: 11, Align: align.Right, Style: fontstyle.Bold}),
	).WithStyle(&props.Cell{BorderType: border.Top})
}

// Table header, body, footer
func getTable(columnWidths *[]int, tableHeader *[]string, data *model.BalanceSheetOuput) []core.Row {
	// Add table header
	cols := []core.Col{}
	for i, value := range *tableHeader {
		textAlign := align.Left
		if i > 1 {
			textAlign = align.Center
		}
		cols = append(cols, text.NewCol((*columnWidths)[i], value, props.Text{Size: 11, Align: textAlign, Color: utils.Colors.White, Style: fontstyle.Bold}))
	}
	header := row.New(7).Add(cols...).WithStyle(&props.Cell{BackgroundColor: utils.Colors.Red})

	// Add table body
	assetsHeader := row.New(6).Add(text.NewCol(12, "ACTIVO", props.Text{Size: 12, Align: align.Left, Style: fontstyle.Bold}))

	// Current assets
	currentAssets := []core.Row{row.New(6).Add(text.NewCol(12, "Activo corrente", props.Text{Size: 12, Align: align.Left, Style: fontstyle.Bold}))}
	for _, entry := range data.Assets.CurrentAssets {
		currentAssets = append(currentAssets, getRow(entry, columnWidths))
	}
	currentAssets = append(currentAssets,
		getRow(&model.ReportRowItem{Account: "", Label: "Activo corrente total", Value: data.Assets.TotalCurrentAssets}, columnWidths).WithStyle(&props.Cell{BorderType: border.Top}),
	)

	// Fixed assets
	fixedAssets := []core.Row{row.New(6).Add(text.NewCol(12, "Activo fixo", props.Text{Size: 12, Align: align.Left, Style: fontstyle.Bold}))}
	for _, entry := range data.Assets.FixedAssets {
		fixedAssets = append(fixedAssets, getRow(entry, columnWidths))
	}
	fixedAssets = append(fixedAssets,
		getRow(&model.ReportRowItem{Account: "", Label: "Activo fixo total", Value: data.Assets.TotalFixedAssets}, columnWidths).WithStyle(&props.Cell{BorderType: border.Top}),
	)

	// Total assets
	totalAssets := getTotalRow(&model.ReportRowItem{Account: "", Label: "Activo total", Value: data.Assets.TotalAssets}, columnWidths)

	liabilitiesHeader := row.New(6).Add(text.NewCol(12, "PASSIVO", props.Text{Size: 12, Align: align.Left, Style: fontstyle.Bold}))

	// Current liabilities
	currentLiabilities := []core.Row{row.New(6).Add(text.NewCol(12, "Passivo corrente", props.Text{Size: 12, Align: align.Left, Style: fontstyle.Bold}))}
	for _, entry := range data.Liabilities.CurrentLiabilities {
		currentLiabilities = append(currentLiabilities, getRow(entry, columnWidths))
	}
	currentLiabilities = append(currentLiabilities,
		getRow(&model.ReportRowItem{Account: "", Label: "Passivo corrente total", Value: data.Liabilities.TotalCurrentLiabilities}, columnWidths).WithStyle(&props.Cell{BorderType: border.Top}),
	)

	// Non-current liabilities
	nonCurrentLiabilities := []core.Row{row.New(6).Add(text.NewCol(12, "Passivo não-corrente", props.Text{Size: 12, Align: align.Left, Style: fontstyle.Bold}))}
	for _, entry := range data.Liabilities.NonCurrentLiabilities {
		nonCurrentLiabilities = append(nonCurrentLiabilities, getRow(entry, columnWidths))
	}
	nonCurrentLiabilities = append(nonCurrentLiabilities,
		getRow(&model.ReportRowItem{Account: "", Label: "Passivo não-corrente total", Value: data.Liabilities.TotalNonCurrentLiabilities}, columnWidths).WithStyle(&props.Cell{BorderType: border.Top}),
	)

	// Total liability
	totalLiabilities := getTotalRow(&model.ReportRowItem{Account: "", Label: "Passivo total", Value: data.Liabilities.TotalLiabilities}, columnWidths)

	// Equity
	equityHeader := row.New(6).Add(text.NewCol(12, "CAPITAL PRÓPRIO", props.Text{Size: 12, Align: align.Left, Style: fontstyle.Bold}))
	equity := []core.Row{}
	for _, entry := range data.Equity.Equity {
		equity = append(equity, getRow(entry, columnWidths))
	}

	totalEquity := getRow(&model.ReportRowItem{Account: "", Label: "Capital próprio total", Value: data.Equity.TotalEquity}, columnWidths).WithStyle(&props.Cell{BorderType: border.Top})
	totalLiabilityAndEquity := getTotalRow(&model.ReportRowItem{Account: "", Label: "Total passivo + capital próprio", Value: data.TotalLiabilityAndEquity}, columnWidths)

	rows := []core.Row{header, row.New(2), assetsHeader}
	rows = append(rows, currentAssets...)
	rows = append(rows, row.New(1))
	rows = append(rows, fixedAssets...)
	rows = append(rows, totalAssets)
	rows = append(rows, row.New(2))
	rows = append(rows, liabilitiesHeader)
	rows = append(rows, currentLiabilities...)
	rows = append(rows, nonCurrentLiabilities...)
	rows = append(rows, totalLiabilities)
	rows = append(rows, row.New(2))
	rows = append(rows, equityHeader)
	rows = append(rows, equity...)
	rows = append(rows, totalEquity)
	rows = append(rows, totalLiabilityAndEquity)

	return rows
}
