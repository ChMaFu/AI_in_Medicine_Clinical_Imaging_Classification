package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"sync"

	"github.com/disintegration/imaging"
)

func main() {
	var srcFolderSpec = "/data3/home/chmafu/AI_in_Medicine_Clinical_Imaging_Classification/data//train/*.jpeg"
	var dstFolder = "/data3/home/chmafu/AI_in_Medicine_Clinical_Imaging_Classification/data/train-resized-256/"
	files, err := filepath.Glob(srcFolderSpec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d files to resize\n", len(files))

	sem := make(chan struct{}, 8)
	var wg sync.WaitGroup
	for _, f := range files {
		wg.Add(1)
		go func(file string, dstFolder string) {
			sem <- struct{}{}
			defer func() { <-sem }()
			defer wg.Done()

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
