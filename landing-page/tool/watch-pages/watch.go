package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/orca-cpfr/orca-cpfr.github.io/landing-page/internal/app"
)

var (
	Path = "src/pages"
)

func main() {
	err := app.Render()
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
					err := app.Render()
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

	err = watcher.Add(Path)
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

func PrintIfError() {

}
