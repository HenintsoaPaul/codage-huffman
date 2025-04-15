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

func (pq MyQueue) Len() int { return len(pq) }
func (pq MyQueue) Less(i, j int) bool {
	if pq[i].Frequency == pq[j].Frequency {
		return pq[i].Char < pq[j].Char // Ordre alphabétique pour la stabilité
	}
	return pq[i].Frequency < pq[j].Frequency
}
func (pq MyQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
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

func BuildTree(text string) *Node {
	tree := &MyQueue{}
	heap.Init(tree)

	dico := GetDictionary(text)

	for char, freq := range dico {
		heap.Push(tree, &Node{Char: char, Frequency: freq})
	}

	for tree.Len() > 1 {
		left := heap.Pop(tree).(*Node)
		right := heap.Pop(tree).(*Node)
		// parent := &Node{Frequency: left.Frequency + right.Frequency, Left: left, Right: right}
		parent := &Node{Frequency: left.Frequency + right.Frequency, Left: left, Right: right, Char: left.Char + right.Char}
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
