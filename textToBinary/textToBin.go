package main

import (
	"container/heap"
	"fmt"
)

// Définition de l'interface Node
type Node struct {
	Char      rune // only pour les feuilles
	Frequency int  // p_i (probability) ou f_i (frequency)
	Left      *Node
	Right     *Node
}

// Définition de l'interface MyQueue (tas min) afin de definir les regles d'un heap(tas)
type MyQueue []*Node

// 1. Nombre d'éléments dans le tas
func (pq MyQueue) Len() int {
	return len(pq)
}

// 2. Définir l'ordre du tas min : comparer les fréquences
func (pq MyQueue) Less(i, j int) bool {
	return pq[i].Frequency < pq[j].Frequency
}

// 3. Échanger deux éléments
func (pq MyQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// 4. Ajouter un élément à la fin (méthode imposée par heap.Interface)
func (pq *MyQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

// 5. Supprimer le plus petit élément (racine)
func (pq *MyQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1] // Récupérer le dernier élément
	*pq = old[:n-1]  // Réduire la taille du tas
	return item
}

func buildTree(dico map[rune]int) *Node {
	tree := make(MyQueue, 0)
	heap.Init(&tree)

	// create a node for each ktr
	for ktr, freq := range dico {
		heap.Push(&tree, &Node{Char: ktr, Frequency: freq})
	}

	// build tree by combinaison
	for tree.Len() > 1 {
		// a chaque fois que l'on fait heap.Pop() on obtient toujours l'elmt Less()
		left := heap.Pop(&tree).(*Node)
		right := heap.Pop(&tree).(*Node)

		// create a new node using the parents
		parent := &Node{
			Frequency: left.Frequency + right.Frequency,
			Left:      left,
			Right:     right,
		}

		heap.Push(&tree, parent)
	}

	// retourne la racine
	return heap.Pop(&tree).(*Node)
}

func printTree(tree *Node, code string) {
	if tree == nil {
		return
	}

	// if feuille -> afficher
	if tree.Left == nil && tree.Right == nil {
		fmt.Printf("Char: %c, Code: %s\n", tree.Char, code)
		return
	}

	printTree(tree.Left, code+"0")
	printTree(tree.Right, code+"1")
}

func getDictionary(text string) map[rune]int {
	// rune => caractere unicode
	dico := make(map[rune]int)

	for _, char := range text {
		if _, contains := dico[char]; !contains {
			dico[char] = 1
		} else {
			dico[char]++
		}
	}

	return dico
}

func printDictionary(dictionary map[rune]int) {
	for key, value := range dictionary {
		fmt.Printf("%c\t%d\n", key, value)
	}
}

func getHufmannCodes(tree *Node, code string, result map[rune]string) {
	if tree == nil {
		return
	}

	if tree.Left == nil && tree.Right == nil {
		result[tree.Char] = code
	}

	getHufmannCodes(tree.Left, code+"0", result)
	getHufmannCodes(tree.Right, code+"1", result)
}

func printCodes(myMap map[rune]string) {
	for key, value := range myMap {
		fmt.Printf("%c\t%s\n", key, value)
	}
}

func main() {
	text := "This is a test"

	dico := getDictionary(text)
	//printDictionary(dico)

	tree := buildTree(dico)
	//printTree(tree, "")
	codes := make(map[rune]string)
	getHufmannCodes(tree, "", codes)
	printCodes(codes)
}
