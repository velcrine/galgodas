package main

import "fmt"

// let's take these are the three towers.
// why they are declared as package level vars: see Print function
var a = []int{7, 5, 3, 2, 1}
var b = make([]int, 0)
var c = make([]int, 0)

// u loose the ability to play with people's mind and make them fall into a loop,
// may be not an evil intention.
// u gain the power of claiming truth.
// truth is ordinary and hidden

func main() {
	PrintTower()

	// as this implementation is based on standard approach:
	// We pass the no. of discs that needs to be moved.
	// As the rules are, before moving that "one",
	// program must move all the ones that are above it.
	MoveTower(&a, &b, &c, 5)
}

// args:
// source tower, target tower, tempSpace tower,
// nth disc that needs to be shifted

func MoveTower(A *[]int, B *[]int, C *[]int, n int) {
	// to do array operations, get out the index of nth disc in array
	// (len(*A) - 1) : the index of last disc
	// minus
	// n -1 : the discs above the last(nth) disc
	index := (len(*A) - 1) - n + 1

	// imagine a case,
	// where the element you want to move is the top most element.
	// Then? Then you can directly transfer the element to the target tower.
	// Just pop from source and push to the target
	if n == 1 {
		*B = append(*B, (*A)[index])
		*A = (*A)[:index]
		PrintTower()
		return
	}

	// make sure all the discs above the last disc are transferred to tempSpace,
	// so that the last one can be simply moved to target,
	// hence do that all except one disc are transferred.
	MoveTower(A, C, B, n-1)

	// transfer the bottom most disc, out of all
	*B = append(*B, (*A)[index])
	*A = (*A)[:index]
	PrintTower()

	// transfer n-1 discs from tempSpace to target,
	// which were transferred to tempSpace tower,
	// so that bottom most disc can be transferred freely

	MoveTower(C, B, A, n-1)
}

// This function should always print in order : Source, Target and TempSpace.
// But as the towers a,b,c continuously change their role as source, target and temp,
// the function has to rely on global references, for accurate mapping.
// Hence, the towers have been declared globally to be accessed, instead of being passed.

func PrintTower() {
	x, y, z := 0, 0, 0
	for i := range [5]int{} {
		x, y, z = 0, 0, 0
		if len(a) > 4-i {
			x = a[4-i]
		}
		if len(b) > 4-i {
			y = b[4-i]
		}
		if len(c) > 4-i {
			z = c[4-i]
		}
		fmt.Printf("%v   %v   %v\n",
			x, y, z)
	}
	fmt.Println()
	fmt.Println()

}
