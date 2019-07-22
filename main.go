package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

type files []string

func (f *files) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *files) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func main() {
	var command string
	var filesToWatch files
	flag.StringVar(&command, "command", "",
		"command to execute when a file changes")
	flag.Var(&filesToWatch, "file", "file to watch for changes;"+
		"flag can be used multiple times to specify multiple files")
	flag.Parse()

	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	for _, file := range filesToWatch {
		err = w.Add(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				return
			}
			log.Println("file changed", event)
			cmd := exec.Command("sh", "-c", command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Println(err)
			}
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		case <-signals:
			w.Close()
			return
		}
	}
}
