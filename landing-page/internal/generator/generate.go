package generator

import (
	"os"
	"path/filepath"

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
		HomeURL:      "index.html",
		ContactUsURL: "contact-us.html",
		ProductURL:   "product-features.html",
		LoginURL:     "/demo/app/login.html",
		Menus: []Menu{
			{
				Name: "Home",
				URL:  "index.html",
			},
			{
				Name: "Products",
				URL:  "product.html",
			},
			{
				Name: "Blogs & News",
				URL:  "blog.html",
			},
			{
				Name: "Contact Us",
				URL:  "contact-us.html",
			},
			{
				Name: "About",
				URL:  "about.html",
			},
		},
		Problems: []TwoFields{
			{"images/icon_bad_data.svg", "Scattered and Unclean Data"},
			{"images/icon_inaccurate.svg", "Inaccuracy of Planning / Forecasting"},
			{"images/icon_disaster.svg", "Lack of mitigation in real-life operational"},
		},
		KeyFeatures: []string{
			"Zero-learning data collection (Whatsapp)",
			"Tracking & Managing Performance",
			"AI-Driven Data Cleansing",
			"AI-Driven Baseline Forecasting",
			"CPFR for Strategic Business Planning",
		},
		TeamMembers: []TeamMember{
			{
				Name:        "Pelitawan Tjandrasa",
				Title:       "CEO / Co-Founder",
				Description: "30+ years business leader in FMCG  & Consumer Electronics industry including Nestle, PZ Cussons, Perfetti Van Melle, Samsung & Philips.",
				ImageURL:    "https://loremflickr.com/320/320/dog",
				LinkedInURL: "https://www.linkedin.com/in/pelitawan-t-19178526/",
			},
			{
				Name:        "Iman Tunggono",
				Title:       "CTIO / Co-Founder",
				Description: "14+ years of experience in software development for various industries like Ride-hailing, OTA, and Fintech.",
				ImageURL:    "https://loremflickr.com/320/320/dog",
				LinkedInURL: "https://www.linkedin.com/in/iman-tunggono/",
			},
		},
		Products: []Product{
			{
				Name: "WhatsApp ChatBot",
				Features: []Feature{
					{Name: "Find nearby store with Geotagging"},
					{Name: "Data Collection"},
					{Name: "Order plan"},
					{Name: "AI Assistance"},
				},
			},
			{
				Name: "Dashboard",
				Features: []Feature{
					{Name: "Master Data"},
					{Name: "Store Insight Reports"},
					{Name: "Sales-Forecast Report"},
					{Name: "Order-Forecast Reports"},
					{Name: "Front-end Agreement Forms"},
					{Name: "Join-Biz Plan Forms"},
				},
			},
			{
				Name: "AI Model",
				Features: []Feature{
					{Name: "Data Cleansing"},
					{Name: "Forecasting"},
				},
			},
		},
	}
)

type (
	SiteData struct {
		PageName     string
		Meta         Meta
		HomeURL      string
		ContactUsURL string
		ProductURL   string
		LoginURL     string
		Menus        []Menu
		Problems     []TwoFields
		KeyFeatures  []string
		TeamMembers  []TeamMember
		Products     []Product
	}
	Menu struct {
		Name string
		URL  string
	}
	TwoFields struct {
		Field1 string
		Field2 string
	}
	Meta struct {
		Title       string
		Description string
	}
	TeamMember struct {
		Name        string
		Title       string
		Description string
		ImageURL    string
		LinkedInURL string
	}
	Product struct {
		Name     string
		Features []Feature
	}
	Feature struct {
		Name string
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

		if err := tmpl.Execute(file, data); err != nil {
			return err
		}
	}
	return nil
}
