package generator

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/imantung/dirtmpl"
)

var (
	_Root   = "src/pages"
	_OutDir = "public"
	_Data   = SiteData{
		Meta: Meta{
			Title:       "demand-sense.ai | Strategic Demand Planning",
			Description: "AI-Driven Platform for Reliable Strategic Planning, and Operational Mitigation Actions with Zero Learning",
		},
		HomeURL: "index.html",
		GroupMenus: []GroupMenu{
			{
				Name: "PRINCIPAL",
				Menus: []Menu{
					{Name: "Master Distributor", URL: "#", Icon: SvgTable},
					{Name: "Store Reports", URL: "#", Icon: SvgPie},
					{Name: "Order Forecasting", URL: "#", Icon: SvgInsight},
					{Name: "Sales Forecasting", URL: "#", Icon: SvgInsight},
					{Name: "Principal Inventory", URL: "#", Icon: SvgBox},
				},
			},
			{
				Name: "DISTRIBUTOR",
				Menus: []Menu{
					{Name: "Master Salesman", URL: "#", Icon: SvgTable},
					{Name: "Master Store", URL: "#", Icon: SvgTable},
					{Name: "Salesman Reports", URL: "#", Icon: SvgPie},
					{Name: "Distributor Inventory", URL: "#", Icon: SvgBox},
					{Name: "Distributor Replenishment", URL: "#", Icon: SvgBox},
					{Name: "Store Replenishment", URL: "#", Icon: SvgBox},
				},
			},
			{
				Name: "SALESMAN",
				Menus: []Menu{
					{Name: "Find Nearby Store", URL: "#", Icon: SvgDevice},
					{Name: "Survey Store", URL: "#", Icon: SvgDevice},
					{Name: "Order Proposal", URL: "#", Icon: SvgDevice},
					{Name: "Salesman Activity Logs", URL: "#", Icon: SvgPie},
				},
			},
		},
		HomePage: HomePage{
			Summaries: []CardSummary{
				{
					Name:            "Total views",
					Value:           "$3.456K",
					ValuePercentage: "0.43%",
					PositiveValue:   true,
					Icon:            SvgEye,
				},
				{
					Name:            "Total profit",
					Value:           "$45,2K",
					ValuePercentage: "4.35%",
					PositiveValue:   true,
					Icon:            SvgShoppingCart,
				},
				{
					Name:            "Total Product",
					Value:           "2.450",
					ValuePercentage: "2.59%",
					PositiveValue:   true,
					Icon:            SvgShoppingBag,
				},
				{
					Name:            "Total Users",
					Value:           "3.456",
					ValuePercentage: "0.95%",
					PositiveValue:   false,
					Icon:            SvgUsers,
				},
			},
		},
	}
)

type (
	SiteData struct {
		PageName   string
		PageLevel  int
		Meta       Meta
		HomeURL    string
		GroupMenus []GroupMenu
		HomePage   HomePage
	}
	Meta struct {
		Title       string
		Description string
	}
	GroupMenu struct {
		Name  string
		Menus []Menu
	}
	Menu struct {
		Name string
		URL  string
		Icon template.HTML
	}
	HomePage struct {
		Summaries []CardSummary
	}
	CardSummary struct {
		Name            string
		Value           string
		ValuePercentage string
		PositiveValue   bool
		Icon            template.HTML
	}
)

func Generate() error {
	m, err := dirtmpl.HTMLTemplates(_Root)
	if err != nil {
		return err
	}

	for k, tmpl := range m {
		path := filepath.Join(_OutDir, k)
		os.Mkdir(filepath.Dir(path), os.ModePerm)

		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()

		data := _Data
		data.PageName = k
		data.PageLevel = GetPageLevel(k)

		if err := tmpl.Execute(file, data); err != nil {
			return err
		}
	}
	return nil
}

func GetPageLevel(pageName string) int {
	elems := strings.Split(pageName, "/")
	return len(elems)
}

func (s SiteData) RelativePath(path string) string {
	for i := 1; i < s.PageLevel; i++ {
		path = "../" + path
	}
	return path
}
