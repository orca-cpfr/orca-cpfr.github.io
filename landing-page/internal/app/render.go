package app

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	Path   = "src/pages"
	OutDir = "public"
	Data   = struct {
		PageName     string
		Meta         Meta
		HomeURL      string
		ContactUsURL string
		ProductURL   string
		Menus        []TwoFields
		Problems     []TwoFields
		KeyFeatures  []string
		Founders     []Founder
		Products     []Product
	}{
		Meta: Meta{
			Title:       "orca-cpfr.io | AI-Driven CPFR Platform",
			Description: "AI-Driven Platform for Reliable Strategic Planning, and Operational Mitigation Actions with Zero Learning",
		},
		HomeURL:      "index.html",
		ContactUsURL: "contact-us.html",
		ProductURL:   "product-features.html",
		Menus: []TwoFields{
			{"index.html", "Home"},
			{"product-features.html", "Product Features"},
			{"blog.html", "Blogs & News"},
			{"about-us.html", "About Us"},
		},
		Problems: []TwoFields{
			{"images/icon_bad_data.svg", "Scattered and Unclean Data"},
			{"images/icon_inaccurate.svg", "Inaccuracy of Planning / Forecasting"},
			{"images/icon_disaster.svg", "Lack of mitigation in real-life operational"},
		},
		KeyFeatures: []string{
			"Data collection at store level with WhatsApp",
			"AI Model for Data Cleansing",
			"AI Model for Demand Forecasting",
			"Collaboration tool for uplift and business strategy",
			"Review and mitigation action tool (if needed)",
		},
		Founders: []Founder{
			{
				Name:        "Pelitawan Tjandrasa",
				Title:       "CEO",
				Description: "30+ years business leader in FMCG  & Consumer Electronics industry including Nestle, PZ Cussons, Perfetti Van Melle, Samsung & Philips.",
				ImageURL:    "https://loremflickr.com/320/320/dog",
				LinkedInURL: "https://www.linkedin.com/in/pelitawan-t-19178526/",
			},
			{
				Name:        "Roy Djunaidi",
				Title:       "COO / CFO",
				Description: "6 years of Venture experience with focus on B2B corporate innovation. Working with distribution partners for freight forwarding business.",
				ImageURL:    "https://loremflickr.com/320/320/dog",
				LinkedInURL: "https://www.linkedin.com/in/roydjunaidi/",
			},
			{
				Name:        "Iman Tunggono",
				Title:       "CTIO",
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
	TwoFields struct {
		Field1 string
		Field2 string
	}
	Meta struct {
		Title       string
		Description string
	}
	Founder struct {
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

func Render() error {
	var m map[string][]string = make(map[string][]string)
	WalkTemplates(m, Path, []string{})

	os.Mkdir(OutDir, os.ModePerm)

	for k, v := range m {
		file, err := os.Create(filepath.Join(OutDir, k))
		if err != nil {
			return err
		}
		defer file.Close()
		tmpl := template.Must(template.ParseFiles(v...))

		data := Data
		data.PageName = k

		if err := tmpl.Execute(file, data); err != nil {
			return err
		}
	}
	return nil
}

func WalkTemplates(m map[string][]string, parent string, list []string) {
	entries, _ := os.ReadDir(parent)
	for _, entry := range entries {
		filename := entry.Name()
		fullPath := filepath.Join(parent, filename)

		if strings.HasPrefix(filename, "_") {
			list = append(list, fullPath)
		} else {
			if entry.IsDir() {
				WalkTemplates(m, fullPath, list)
			} else {
				list2 := make([]string, len(list))
				copy(list2, list)

				key := fullPath[len(Path)+1:]
				m[key] = append(list2, fullPath)
			}
		}
	}
}
