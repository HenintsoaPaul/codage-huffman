package textToBinary

import (
	"container/heap"
	"sort"
)

type Symbol struct {
	Char rune
	Idx  int
	Freq int
}

func GetDictionary2(text string) map[rune]Symbol {
	dico := make(map[rune]Symbol)
	for idx, char := range text {
		if symbol, found := dico[char]; found {
			freq := symbol.Freq + 1
			dico[char] = Symbol{Char: symbol.Char, Idx: symbol.Idx, Freq: freq}
		} else {
			dico[char] = Symbol{Char: char, Idx: idx, Freq: 1}
		}
	}
	return dico
}

type YourNode struct {
	Char       rune
	Frequency  int
	FirstIndex int
	Left       *YourNode
	Right      *YourNode
}

type YourQueue []*YourNode

func (pq YourQueue) Len() int { return len(pq) }
func (pq YourQueue) Less(i, j int) bool {
	if pq[i].Frequency == pq[j].Frequency {
		return pq[i].FirstIndex < pq[j].FirstIndex // Ordre d'occurence pour la stabilité
	}
	return pq[i].Frequency < pq[j].Frequency
}
func (pq YourQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *YourQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*YourNode))
}
func (pq *YourQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func BuildTree2(text string) *YourNode {
	tree := &YourQueue{}
	heap.Init(tree)

	dico := GetDictionary2(text)

	keys := make([]rune, 0, len(dico))
	for key := range dico {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return dico[keys[i]].Idx < dico[keys[j]].Idx
	})

	for _, key := range keys {
		symbol := dico[key]
		heap.Push(tree, &YourNode{Char: key, Frequency: symbol.Freq, FirstIndex: symbol.Idx})
	}

	for tree.Len() > 1 {
		left := heap.Pop(tree).(*YourNode)
		right := heap.Pop(tree).(*YourNode)
		parent := &YourNode{Frequency: left.Frequency + right.Frequency, Left: left, Right: right, FirstIndex: right.FirstIndex}
		heap.Push(tree, parent)
	}

	return heap.Pop(tree).(*YourNode)
}

func GetHuffmanCodes2(tree *YourNode, code string, result map[rune]string) {
	if tree == nil {
		return
	}

	if tree.Left == nil && tree.Right == nil {
		result[tree.Char] = code
	}

	GetHuffmanCodes2(tree.Left, code+"0", result)
	GetHuffmanCodes2(tree.Right, code+"1", result)
}
