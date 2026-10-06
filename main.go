package main

import "fmt"

func main() {
	liste := []int{5, 2, 8, 1, 7, 3, 6, 4}

	fmt.Println("liste de base :", liste)

	bubble := append([]int(nil), liste...)
	BubbleSort(bubble)
	fmt.Println("bubble :", bubble)

	selection := append([]int(nil), liste...)
	SelectionSort(selection)
	fmt.Println("selection :", selection)

	insertion := append([]int(nil), liste...)
	InsertionSort(insertion)
	fmt.Println("insertion :", insertion)

	merge := MergeSort(liste)
	fmt.Println("merge :", merge)

	quick := append([]int(nil), liste...)
	QuickSort(quick)
	fmt.Println("quick :", quick)

	players := []Score{
		{Player: "Léa", Score: 12},
		{Player: "Tom", Score: 15},
		{Player: "Ana", Score: 12},
		{Player: "Max", Score: 9},
	}

	fmt.Println("\njoueurs avant :", players)

	insertionPlayers := append([]Score(nil), players...)
	InsertionSortScores(insertionPlayers)
	fmt.Println("insertion joueurs :", insertionPlayers)
	fmt.Println("stable :", IsStable(insertionPlayers))

	selectionPlayers := append([]Score(nil), players...)
	SelectionSortScores(selectionPlayers)
	fmt.Println("selection joueurs :", selectionPlayers)
	fmt.Println("stable :", IsStable(selectionPlayers))
}
