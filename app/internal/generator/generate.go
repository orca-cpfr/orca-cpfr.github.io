package generator

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	Root   = "src/pages"
	OutDir = "public"
	Data   = struct {
		PageName string
		Meta     Meta
		HomeURL  string
	}{
		Meta: Meta{
			Title:       "orca-cpfr.io | AI-Driven CPFR Platform",
			Description: "AI-Driven Platform for Reliable Strategic Planning, and Operational Mitigation Actions with Zero Learning",
		},
		HomeURL: "/index.html",
	}
)

type (
	Meta struct {
		Title       string
		Description string
	}
)

func Generate() error {
	var m map[string][]string = make(map[string][]string)
	WalkTemplates(m, Root, []string{})

	for k, v := range m {
		path := filepath.Join(OutDir, k)
		os.Mkdir(filepath.Dir(path), os.ModePerm)

		file, err := os.Create(path)
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

				key := fullPath[len(Root)+1:]
				m[key] = append(list2, fullPath)
			}
		}
	}
}
