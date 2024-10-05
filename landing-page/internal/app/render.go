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
		SiteURL  string
		Problems []KV
		Features []string
	}{
		SiteURL: "#",
		Problems: []KV{
			{"images/icon_bad_data.png", "Scattered and Unclean Data"},
			{"images/icon_inaccurate.png", "Inaccuracy of Planning / Forecasting"},
			{"images/icon_disaster.png", "Lack of mitigation in real-life operational"},
		},
		Features: []string{
			"Data collection at store level with WhatsApp",
			"AI Model for Data Cleansing",
			"AI Model for Demand Forecasting",
			"Collaboration tool for uplift and business strategy",
			"Review and mitigation action tool (if needed)",
		},
	}
)

type (
	KV struct {
		Key   string
		Value string
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

		if err := tmpl.Execute(file, &Data); err != nil {
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
