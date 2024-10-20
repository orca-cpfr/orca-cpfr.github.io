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
		SiteURL: "https://demand-sense.ai",
		Meta: Meta{
			Title:       "demand-sense.ai | Strategic Demand Planning",
			Description: "AI-Driven Platform for Reliable Strategic Planning, and Operational Mitigation Actions with Zero Learning",
		},
		HomeURL: "index.html",
		UserProfile: UserProfile{
			Name:     "John Doe",
			JobTitle: "Sales Manager",
			PicURL:   "https://xsgames.co/randomusers/avatar.php?g=male",
		},
		CompanyProfile: CompanyProfile{
			Name: "XYZ Lte.",
		},
		Notifications: []Notification{
			{
				Message: "Edit your information in a swipe Sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim.",
				Date:    "12 May, 2025",
			},
			{
				Message: "It is a long established fact that a reader will be distracted by the readable.",
				Date:    "24 Feb, 2025",
			},
			{
				Message: "There are many variations of passages of Lorem Ipsum available, but the majority have suffered",
				Date:    "04 Jan, 2025",
			},
			{
				Message: "There are many variations of passages of Lorem Ipsum available, but the majority have suffered",
				Date:    "01 Dec, 2024",
			},
		},

		GroupMenus: []GroupMenu{
			{
				Name: "MAIN DASHBOARD",
				Menus: []Menu{
					{Name: "Home", URL: "#", Icon: SvgHome},
					{Name: "User Controls", URL: "#", Icon: SvgUsers},
					{Name: "Billings", URL: "#", Icon: SvgUsers},
				},
			},
			{
				Name: "MASTER DATA",
				Menus: []Menu{
					{Name: "Master Distributor", URL: "#", Icon: SvgTable},
					{Name: "Master Store", URL: "#", Icon: SvgTable},
					{Name: "Master Salesman", URL: "#", Icon: SvgTable},
					{Name: "Master Product", URL: "#", Icon: SvgTable},
					{Name: "Master Campaign", URL: "#", Icon: SvgTable},
				},
			},
			{
				Name: "PERFORMANCE TRACKER",
				Menus: []Menu{
					{Name: "Salesman Evaluation", URL: "#", Icon: SvgPie},
					{Name: "Store Evaluation", URL: "#", Icon: SvgPie},
					{Name: "Merchandising Evaluation", URL: "#", Icon: SvgPie},
					{Name: "Competitor Evaluation", URL: "#", Icon: SvgPie},
				},
			},
			{
				Name: "ANALYTICS PLATFORM",
				Menus: []Menu{
					{Name: "Data Cleansing", URL: "#", Icon: SvgLab},
					{Name: "Baseline Forecasting", URL: "#", Icon: SvgLab},
				},
			},
			{
				Name: "COLLABORATIVE PLANNING",
				Menus: []Menu{
					{Name: "Sales Forecasting", URL: "#", Icon: SvgInsight},
					{Name: "Order Forecasting", URL: "#", Icon: SvgInsight},
					{Name: "Join Business Plan", URL: "#", Icon: SvgPresentation},
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
		SiteURL        string
		PageName       string
		PageLevel      int
		UserProfile    UserProfile
		CompanyProfile CompanyProfile
		Notifications  []Notification
		Meta           Meta
		HomeURL        string
		GroupMenus     []GroupMenu
		HomePage       HomePage
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
	UserProfile struct {
		Name     string
		JobTitle string
		PicURL   string
	}
	CompanyProfile struct {
		Name string
	}
	Notification struct {
		Message string
		Date    string
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
