package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/orca-cpfr/orca-cpfr.github.io/landing-page/internal/generator"
)

var (
	_Root = "src/pages"
)

func main() {
	err := generator.Generate()
	if err != nil {
		log.Println("error:", err)
	}

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	build := false

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if !build && !event.Has(fsnotify.Chmod) {
					build = true
					fmt.Println("build pages..")
					err := generator.Generate()
					if err != nil {
						log.Println("error:", err)
					}
					build = false
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	filepath.WalkDir(_Root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			err := watcher.Add(path)
			if err != nil {
				return err
			}
			fmt.Println("Add path: ", path)
		}
		return nil
	})

	// Block main goroutine forever.
	<-make(chan struct{})
}

func PrintIfError() {

}
