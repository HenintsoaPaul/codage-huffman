package main

import (
	"fmt"
	"huffman/utils/imgToBinary"
	"huffman/utils/textToBinary"
	"huffman/utils/wavToBinary"
	"log"
	"strings"
)

func getHiddenMessage(data []byte, targetIndexes []int) string {
	var hiddenMessage strings.Builder

	for _, index := range targetIndexes {
		hiddenMessage.WriteString(fmt.Sprintf("%08b", data[index]))
	}

	return hiddenMessage.String()
}

func text() {
	text := "This is a test"
	dico := textToBinary.GetDictionary(text)
	tree := textToBinary.BuildTree(dico)

	codes := make(map[rune]string)
	textToBinary.GetHuffmanCodes(tree, "", codes)

	textToBinary.PrintCodes(codes)

	// image -> bin[] -> encoded -> text normal

	//encoded := "1110101010110111000"
	encoded := "001101001011"
	decoded := textToBinary.DecodeHuffman(encoded, codes)
	fmt.Printf("Texte décodé: '%s'\n", decoded)
}

func img() {
	img, err := imgToBinary.LoadPngImage("resources/test_3x3.png")
	if err != nil {
		log.Fatal(err)
	}

	intensities := imgToBinary.GetGreyscaleIntensities(img)

	bins := imgToBinary.ConvertGreyscalesToBinaries(intensities)
	for _, bin := range bins {
		fmt.Println(bin)
	}
}

func wav() {
	nomFichier := "resources/test.wav"
	dataAudio := wavToBinary.LoadWav(nomFichier)
	binaire := wavToBinary.BytesEnBinaire(dataAudio)

	fmt.Println(dataAudio[:])
	fmt.Println(binaire[:])
}

func main() {
	//text()
	//img()
	wav()
}
