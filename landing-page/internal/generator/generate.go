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
			Title:       "demand-sense.ai | Strategic Demand Planning Platform",
			Description: "AI-Driven Platform for Reliable Strategic Planning, and Operational Mitigation Actions with Zero Learning",
		},
		SiteURL:      "https://demand-sense.ai",
		ContactUsURL: "contact-us.html",
		Menus: []Menu{
			{Name: "Home", URL: "index.html"},
			{
				Name: "Product",
				SubMenus: []Menu{
					{Name: "Data Collection", URL: "product/data-collection.html"},
					{Name: "Performance Tracking", URL: "product/performance-tracking.html"},
					{Name: "Analytics Platform", URL: "product/analytics-platform.html"},
					{Name: "Collaborative Planning", URL: "product/collaborative-planning.html"},
				},
			},
			{Name: "Blog & News", URL: "blog.html"},
			{
				Name: "Demo",
				SubMenus: []Menu{
					{Name: "WhatsApp Demo", URL: "#"},
					{Name: "Dashboard Demo", URL: "demo/dashboard/index.html"},
				},
			},
			{Name: "About", URL: "about.html"},
		},
		Problems: []Problem{
			{Icon: "images/icon_bad_data.svg", Name: "Scattered and Unclean Data"},
			{Icon: "images/icon_inaccurate.svg", Name: "Inaccuracy of Planning / Forecasting"},
			{Icon: "images/icon_disaster.svg", Name: "Lack of mitigation in real-life operational"},
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
				Name:    "Data Collection",
				URL:     "product/data-collection.html",
				Tagline: "Zero-learning data collection with Whatsapp",
				Features: []Feature{
					{Name: "Salesman Check-in/Check-out"},
					{Name: "Stock Survey"},
					{Name: "Merchandising Survey"},
					{Name: "Competitor Survey"},
				},
			},
			{
				Name:    "Performance Tracking",
				URL:     "product/performance-tracking.html",
				Tagline: "Tracking & Managing Performance",
				Features: []Feature{
					{Name: "Salesman Evaluation"},
					{Name: "Store Evaluation"},
					{Name: "Channel Evaluation"},
					{Name: "Region Evaluation"},
				},
			},
			{
				Name:    "Analytics Platform",
				URL:     "product/analytics-platform.html",
				Tagline: "Advanced analytics to forecast future demand",
				Features: []Feature{
					{Name: "Data Cleansing"},
					{Name: "Baseline Forecasting"},
					{Name: "Machine Learning Algorithm"},
					{Name: "AI Predictive Model"},
				},
			},
			{
				Name:    "Collaborative Planning",
				URL:     "product/collaborative-planning.html",
				Tagline: "CPFR for Strategic Business Planning",
				Features: []Feature{
					{Name: "Join-Business Plan"},
				},
			},
		},
	}
)

type (
	SiteData struct {
		SiteURL      string
		PageName     string
		PageLevel    int
		Meta         Meta
		ContactUsURL string
		Menus        []Menu
		Problems     []Problem
		TeamMembers  []TeamMember
		Products     []Product
	}
	Menu struct {
		Name     string
		URL      string
		SubMenus []Menu
	}
	Problem struct {
		Name string
		Icon string
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
		URL      string
		Tagline  string
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

	os.Mkdir(_OutDir, os.ModePerm)

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
