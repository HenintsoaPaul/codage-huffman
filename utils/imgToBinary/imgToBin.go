package imgToBinary

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
)

func LoadPngImage(imgPath string) (image.Image, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du décodage de l'image PNG: %v", err)
	}

	return img, nil
}

func GetGreyscaleIntensities(img image.Image) []uint8 {
	if img == nil {
		log.Fatal("l'image est nil")
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	intensities := make([]uint8, width*height)

	index := 0
	for y := bounds.Min.Y; y < height; y++ {
		for x := bounds.Min.X; x < width; x++ {
			intensities[index] = getPixelGreyscale(x, y, img)
			index++
		}
	}

	return intensities
}

func getPixelGreyscale(x int, y int, img image.Image) uint8 {
	// Method 1: Extract RGB (16 bits) then convert to 8 bits
	r, g, b, _ := img.At(x, y).RGBA()
	r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)
	rr := float64(r8) * 0.299
	gg := float64(g8) * 0.587
	bb := float64(b8) * 0.114
	grayColor := uint8(rr + gg + bb + 0.5) // 0.5 pour arrondir

	// Method 2: use librairy
	// grayColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y

	return grayColor
}

func ConvertGreyscalesToBinaries(intensities []uint8) []string {
	bins := make([]string, len(intensities))

	for i, intensity := range intensities {
		bins[i] = fmt.Sprintf("%08b", intensity)
	}

	return bins
}

func main() {
	img, err := LoadPngImage("test_3x3.png")
	if err != nil {
		log.Fatal(err)
	}

	intensities := GetGreyscaleIntensities(img)

	bins := ConvertGreyscalesToBinaries(intensities)
	for _, bin := range bins {
		fmt.Println(bin)
	}
}
