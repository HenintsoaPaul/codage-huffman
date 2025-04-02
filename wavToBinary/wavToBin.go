package main

import (
	"fmt"
	"log"
	"os"
)

func lireWav(nomFichier string) []byte {
	// Ouvrir le fichier
	f, err := os.Open(nomFichier)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Lire tout le fichier
	data, err := os.ReadFile(nomFichier)
	if err != nil {
		log.Fatal(err)
	}

	// Ignorer l'entête WAV (généralement 44 octets)
	if len(data) > 44 {
		data = data[44:]
	} else {
		log.Fatal("Fichier WAV trop petit pour contenir une en-tête valide")
	}

	return data
}

func bytesEnBinaire(data []byte) []string {
	// Convertir chaque byte en une chaîne binaire de 8 bits
	binaire := make([]string, len(data))
	for i, b := range data {
		binaire[i] = fmt.Sprintf("%08b", b) // Format binaire sur 8 bits
	}
	return binaire
}

func main() {
	nomFichier := "test.wav"
	dataAudio := lireWav(nomFichier)
	binaire := bytesEnBinaire(dataAudio)

	// Afficher les 10 premiers octets en binaire
	fmt.Println(dataAudio[:])
	fmt.Println(binaire[:])
}
