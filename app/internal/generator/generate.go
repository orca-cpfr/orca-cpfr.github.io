package generator

import (
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
					{Name: "Master Distributor", URL: "#", Icon: "table"},
					{Name: "Store Reports", URL: "#", Icon: "pie"},
					{Name: "Order Forecasting", URL: "#", Icon: "insight"},
					{Name: "Sales Forecasting", URL: "#", Icon: "insight"},
					{Name: "Principal Inventory", URL: "#", Icon: "box"},
				},
			},
			{
				Name: "DISTRIBUTOR",
				Menus: []Menu{
					{Name: "Master Salesman", URL: "#", Icon: "table"},
					{Name: "Master Store", URL: "#", Icon: "table"},
					{Name: "Salesman Reports", URL: "#", Icon: "pie"},
					{Name: "Distributor Inventory", URL: "#", Icon: "box"},
					{Name: "Distributor Replenishment", URL: "#", Icon: "box"},
					{Name: "Store Replenishment", URL: "#", Icon: "box"},
				},
			},
			{
				Name: "SALESMAN",
				Menus: []Menu{
					{Name: "Find Nearby Store", URL: "#", Icon: "device"},
					{Name: "Survey Store", URL: "#", Icon: "device"},
					{Name: "Order Proposal", URL: "#", Icon: "device"},
					{Name: "Salesman Activity Logs", URL: "#", Icon: "pie"},
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
		Icon string
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
