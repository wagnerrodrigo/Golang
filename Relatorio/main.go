package main

import (
	"log"
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	// inicializa o Pdf no formato A4
	cfg := config.NewBuilder().
		WithOrientation(orientation.Vertical).
		WithPageSize(pagesize.A4).
		WithLeftMargin(15).
		WithTopMargin(15).
		WithRightMargin(15).
		WithBottomMargin(15).
		Build()

	m := maroto.New(cfg)

	// Criando as sessões do pdf
	// Cabeçalho
	addHeader(m)

	// Corpo
	addInvoiceDetails(m)

	// Lista de item
	addItemList(m)

	// Rodape
	addFooter(m)

	// Salvando o pdf no disco

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("rel/fatura.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("PDF salvo com sucesso!")
}

func addHeader(m core.Maroto) {
	m.AddRow(50,
		image.NewFromFileCol(12, "assets/dev.png",
			props.Rect{
				Center:  true,
				Percent: 75,
			},
		),
	)
	m.AddRow(20,
		text.NewCol(12,
			"Relatório financeiro", props.Text{
				Top:   5,
				Style: fontstyle.Bold,
				Align: align.Center,
				Size:  16,
			},
		),
	)
	m.AddRow(20,
		text.NewCol(12,
			"Fatura", props.Text{
				Top:   5,
				Style: fontstyle.Bold,
				Align: align.Center,
				Size:  12,
			},
		),
	)
}

func addInvoiceDetails(m core.Maroto) {
	m.AddRow(10,
		text.NewCol(6, "Data: "+time.Now().Format("02 Jan 2024"), props.Text{
			Align: align.Left,
			Size:  10,
		}),
		text.NewCol(6, "Fatura #1001", props.Text{
			Align: align.Right,
			Size:  10,
		}),
	)
	m.AddRow(40, line.NewCol(12))

}

type InvoceItem struct {
	Item           string
	Description    string
	Quantity       string
	Price          string
	DiscoutedPrice string
	Total          string
}

func (o InvoceItem) GetHeader() core.Row {
	return row.New(10).Add(
		text.NewCol(2, "item", props.Text{Style: fontstyle.Bold}),
		text.NewCol(3, "Description", props.Text{Style: fontstyle.Bold}),
		text.NewCol(1, "Quantity", props.Text{Style: fontstyle.Bold}),
		text.NewCol(2, "Price", props.Text{Style: fontstyle.Bold}),
		text.NewCol(2, "DiscoutedPrice", props.Text{Style: fontstyle.Bold}),
		text.NewCol(2, "Total", props.Text{Style: fontstyle.Bold}),
	)
}

func (o InvoceItem) GetContent(i int) core.Row {
	r := row.New(5).Add(
		text.NewCol(2, o.Item),
		text.NewCol(3, o.Description),
		text.NewCol(1, o.Quantity),
		text.NewCol(2, o.Price),
		text.NewCol(2, o.DiscoutedPrice),
		text.NewCol(2, o.Total),
	)

	if i%2 == 0 {
		r.WithStyle(&props.Cell{
			BackgroundColor: &props.Color{Red: 240, Green: 240, Blue: 240},
		})
	}

	return r
}

func getObjects() []InvoceItem {
	var items []InvoceItem
	contents := [][]string{
		{"LapTop", "14-inch, 16GB RAM", "1", "R$5000", "R$4800", "R$4800"},
		{"Mouse", "Sem fio Wireless", "2", "R$400", "R$360", "R$360"},
		{"Teclado", "Mecânico, RGB", "1", "R$1000", "R$960", "R$960"},
	}

	for i := 0; i < len(contents); i++ {
		items = append(items, InvoceItem{
			Item:           contents[i][0],
			Description:    contents[i][1],
			Quantity:       contents[i][2],
			Price:          contents[i][3],
			DiscoutedPrice: contents[i][4],
			Total:          contents[i][5],
		})
	}
	return items

}

func addItemList(m core.Maroto) {
	rows, err := list.Build[InvoceItem](getObjects())
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(rows...)
}

func addFooter(m core.Maroto) {
	m.AddRow(15,
		text.NewCol(8, "Montante total", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  10,
			Align: align.Right,
		}),
		text.NewCol(4, "R$6400", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  10,
			Align: align.Center,
		}),
	)
	m.AddRow(40,
		signature.NewCol(6, "Signatário autorizado", props.Signature{
			FontFamily: fontfamily.Courier}),
		code.NewQrCol(10, "https://wagnerrodrigo.dev/", props.Rect{
			Percent:            100,
			Center:             true,
			JustReferenceWidth: true,
		}),
	)

	m.AddAutoRow(
		code.NewBarCol(5, "123456789123", props.Barcode{
			Center: true,
			Type:   barcode.EAN,
		}),
		code.NewMatrixCol(6, "https://wagnerrodrigo.dev/", props.Rect{
			Center:  true,
			Percent: 50,
		}),
	)
}
