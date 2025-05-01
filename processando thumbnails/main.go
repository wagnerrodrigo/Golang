package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	inputDir := "imagens"
	outputDir := "thumbnails"

	processImagens(inputDir, outputDir)
}

func processImagens(inputDir, outputDir string) {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			processThumbnail(filepath.Join(inputDir, file.Name()), outputDir)
		}

	}

	fmt.Println(("Thumbnails criadas"))
}

func processThumbnail(fileInput, outputDir string) {
	file, err := os.Open(fileInput)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	thumbnail := resize.Resize(100, 0, image, resize.Lanczos3)

	outputPath := filepath.Join(outputDir, filepath.Base(fileInput))
	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	switch strings.ToLower(filepath.Ext(fileInput)) {
	case ".jpg", ".jpeg":
		err := jpeg.Encode(outputFile, thumbnail, nil)
		if err != nil {
			panic(err)
		}
	case ".png":
		err := png.Encode(outputFile, thumbnail)
		if err != nil {
			panic(err)
		}
	}

}
