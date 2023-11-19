package main

import (
	"fmt"
	"github.com/gSpera/morse"
	"image"
	"image/png"
	"os"
	"os/exec"
	"strings"
)

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	for i := 999; i >= 0; i-- {
		chars := getChar(fmt.Sprintf("zip/%d/flag/pwd.png", i))
		pwd := strings.ToLower(morse.ToText(chars))
		fmt.Printf("Password for %d is %s (%s)\n", i, pwd, chars)
		output, err := exec.Command("unzip", "-o", "-P", pwd, "-d", fmt.Sprintf("zip/%d", i-1), fmt.Sprintf("zip/%d/flag/flag_%d.zip", i, i)).CombinedOutput()
		fmt.Println(string(output))
		if err != nil {
			panic(err)
		}
	}
}

func getChar(path string) string {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println("Error: File could not be opened")
		panic(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error: File could not be decoded")
		panic(err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	baseColor := rgbaToStr(img.At(0, 0).RGBA())

	var chars []string
	for y := 1; y < height-1; y += 2 {
		var char []string
		var cnt int
		for x := 1; x < width; x++ {
			if rgbaToStr(img.At(x, y).RGBA()) != baseColor {
				cnt++
			} else {
				if cnt == 1 {
					char = append(char, ".")
					cnt = 0
				} else if cnt > 1 {
					char = append(char, "-")
					cnt = 0
				}
			}
		}
		chars = append(chars, strings.Join(char, ""))
	}

	return strings.Join(chars, " ")
}

func rgbaToStr(r uint32, g uint32, b uint32, _ uint32) string {
	return fmt.Sprintf("%d/%d/%d", int(r/257), int(g/257), int(b/257))
}
