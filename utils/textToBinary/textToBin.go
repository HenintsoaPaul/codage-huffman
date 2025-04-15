package textToBinary

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	Char      rune
	Frequency int
	Left      *Node
	Right     *Node
}

type MyQueue []*Node

func (pq MyQueue) Len() int           { return len(pq) }
func (pq MyQueue) Less(i, j int) bool { return pq[i].Frequency < pq[j].Frequency }
func (pq MyQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *MyQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}
func (pq *MyQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func BuildTree(dico map[rune]int) *Node {
	tree := &MyQueue{}
	heap.Init(tree)

	// Convertir le dictionnaire en une liste et trier par fréquence
	type CharFreq struct {
		Char      rune
		Frequency int
	}

	var charFreqList []CharFreq
	for char, freq := range dico {
		charFreqList = append(charFreqList, CharFreq{Char: char, Frequency: freq})
	}

	// Trier par fréquence, puis par caractère pour garantir un ordre stable
	sort.SliceStable(charFreqList, func(i, j int) bool {
		if charFreqList[i].Frequency == charFreqList[j].Frequency {
			return charFreqList[i].Char < charFreqList[j].Char
		}
		return charFreqList[i].Frequency < charFreqList[j].Frequency
	})

	// Ajouter les nœuds au tas dans l'ordre trié
	for _, cf := range charFreqList {
		heap.Push(tree, &Node{Char: cf.Char, Frequency: cf.Frequency})
	}

	for tree.Len() > 1 {
		left := heap.Pop(tree).(*Node)
		right := heap.Pop(tree).(*Node)
		parent := &Node{Frequency: left.Frequency + right.Frequency, Left: left, Right: right}
		heap.Push(tree, parent)
	}

	return heap.Pop(tree).(*Node)
}

func GetHuffmanCodes(tree *Node, code string, result map[rune]string) {
	if tree == nil {
		return
	}

	if tree.Left == nil && tree.Right == nil {
		result[tree.Char] = code
	}

	GetHuffmanCodes(tree.Left, code+"0", result)
	GetHuffmanCodes(tree.Right, code+"1", result)
}

func GetDictionary(text string) map[rune]int {
	dico := make(map[rune]int)
	for _, char := range text {
		dico[char]++
	}
	return dico
}

func PrintCodes(codes map[rune]string) {
	keys := make([]rune, 0, len(codes))
	for key := range codes {
		keys = append(keys, key)
	}

	// trier par binaire
	sort.Slice(keys, func(i, j int) bool {
		return codes[keys[i]] < codes[keys[j]]
	})

	for _, key := range keys {
		fmt.Printf("%c - %s\n", key, codes[key])
	}

	// for char, code := range codes {
	// 	fmt.Printf("%c: %s\n", char, code)
	// }
}

func DecodeHuffman(encodedText string, codes map[rune]string) string {
	invertedMap := make(map[string]rune)
	for key, value := range codes {
		invertedMap[value] = key
	}

	var decodedText strings.Builder
	currentCode := ""

	for _, bit := range encodedText {
		currentCode += string(bit)

		if char, found := invertedMap[currentCode]; found {
			decodedText.WriteRune(char)
			currentCode = ""
		}
	}

	return decodedText.String()
}
