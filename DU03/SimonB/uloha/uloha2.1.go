package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	defer close(ch)
	walkRecursive(t, ch)
}

func walkRecursive(t *Tree, ch chan int) {
	if t != nil {
		walkRecursive(t.Left, ch)
		ch <- t.Value
		walkRecursive(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	b := make(chan bool)
	go func() {
		for {
			n1, ok1 := <- ch1
			n2, ok2 := <- ch2
			if ok1 != ok2 || n1 != n2 {
				b <- false
				break
			}
			if !ok1 {
				b <- true
				break
			}
		}
		b <- true
	} ()
	return <- b
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z, nil
}
/*
robil som https://tour.golang.org/methods/20 a https://tour.golang.org/concurrency/7

 */
func main() {
	/*
	ch := make(chan int)
	go Walk(New(1), ch)
	fmt.Println(Same(New(1), New(2)))
	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(2), New(1)))
	time.Sleep(9999999999999999)
	 */
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
