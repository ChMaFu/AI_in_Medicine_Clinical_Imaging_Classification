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

var wg sync.WaitGroup

func main() {
	var srcFolderSpec = "/data3/home/chmafu/AI_in_Medicine_Clinical_Imaging_Classification/data//train/*.jpeg"
	var dstFolder = "/data3/home/chmafu/AI_in_Medicine_Clinical_Imaging_Classification/data/train-resized-256/"

	files, err := filepath.Glob(srcFolderSpec)
	if err != nil {
		log.Fatal(err)
	}
	var fileCount = len(files)
	fmt.Printf("Found %d files to resize\n", fileCount)

	var batchSize = fileCount / runtime.NumCPU()
	var firstIndex = 0
	var lastIndex = batchSize

	for i := 1; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go processImages(files[firstIndex:lastIndex], dstFolder)

		firstIndex = lastIndex + 1
		lastIndex = lastIndex + batchSize
	}
	wg.Wait()
}

func processImages(files []string, outFolder string) {
	defer wg.Done()

	for _, f := range files {
		resizeImage(f, outFolder)
	}

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
