package main

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"

	ico "github.com/Kodeworks/golang-image-ico"
)

func main() {
	pngFileName := "stick.png"
	absPath, err := filepath.Abs(pngFileName)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ConvertPngToIco(file)
}

func ConvertPngToIco(file *os.File) {
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	fmt.Print("Download Pfad angeben:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	outputPath := filepath.Join(scanner.Text(), "output.ico")

	icoFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer icoFile.Close()

	err = ico.Encode(icoFile, rgba)
	if err != nil {
		log.Fatal(err)
	}
	println("Konvertiert")
}
