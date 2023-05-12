package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// this is just a wrapper function to close the channel for WalkRecurse
func Walk(t *tree.Tree, ch chan int) {
	WalkRecurse(t, ch)
	close(ch)
}

func WalkRecurse(t *tree.Tree, ch chan int) {
	// walk the left tree
	if t.Left != nil {
		WalkRecurse(t.Left, ch)
	}
	
	// send value to chan
	ch <- t.Value
	
	// walk the right tree
	if t.Right != nil {
		WalkRecurse(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	
	ch2 := make(chan int)
	go Walk(t2, ch2)
	
	// compare the values on ch1 & ch2
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		
		if !ok1 && !ok2 {
			break
		} else if !(ok1 && ok2) {
			return false
		} else if v1 != v2 {
			return false
		}
	}
	return true
}
	

func main() {
	// ---- Testing Walk ----
	// create a channel
	ch := make(chan int)
	
	// create a tree 1, 2, 3, 4, ..., 10, and
	// walk the tree
	go Walk(tree.New(1), ch)
	
	// read and print 10 values from the channel
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	
	
	// ---- Testing Same ----
	fmt.Println(Same(tree.New(1), tree.New(1)))	// output: true
	fmt.Println(Same(tree.New(1), tree.New(2)))	// output: false
	
	
}
