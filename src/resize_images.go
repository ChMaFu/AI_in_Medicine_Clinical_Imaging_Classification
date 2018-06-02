package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/disintegration/imaging"
)

type empty struct{}
type semaphore chan empty

// acquire n resources
func (s semaphore) P(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func main() {
	var srcFolderSpec = "/data3/home/chmafu/AI_in_Medicine_Clinical_Imaging_Classification/data//train/*.jpeg"
	var dstFolder = "/data3/home/chmafu/AI_in_Medicine_Clinical_Imaging_Classification/data/train-resized-256/"
	files, err := filepath.Glob(srcFolderSpec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d files to resize\n", len(files))

	var wg sync.WaitGroup
	sem := make(semaphore, runtime.NumCPU())
	for _, f := range files {
		wg.Add(1)
		sem.P(1)
		go func(file string, dstFolder string) {
			defer sem.V(1)
			go resizeImage(file, dstFolder)
		}(f, dstFolder)
	}
	wg.Wait()
}

func resizeImage(file string, outFolder string) {
	src, err := imaging.Open(file)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	var img = imaging.Resize(src, 250, 250, imaging.Lanczos)
	//var img = imaging.Blur(src, 2.0)

	var filename string
	_, filename = filepath.Split(file)
	var outfile = path.Join(outFolder, filename)
	fmt.Printf("Writing file: %q\n", outfile)
	err = imaging.Save(img, outfile)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
