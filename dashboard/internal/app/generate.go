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
		PageName string
		Meta     Meta
	}{
		Meta: Meta{
			Title:       "orca-cpfr.io | AI-Driven CPFR Platform",
			Description: "AI-Driven Platform for Reliable Strategic Planning, and Operational Mitigation Actions with Zero Learning",
		},
	}
)

type (
	Meta struct {
		Title       string
		Description string
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