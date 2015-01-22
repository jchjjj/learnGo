package main

import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.!!!!!!in order!!!!!!!!!
func Walk(t *tree.Tree, ch chan int) {
	/* if t.Left == nil  {
	            ch <- t.Value
	        }else if t.Left != nil{
	        Walk(t.Left,ch)
	        //if t.Right != nil {
	          //  Walk(t.Right,ch)
	      //  }
	        }else if t.Right != nil{
	             Walk(t.Right,ch)
	       //      if t.Left != nil{
	       //         Walk(t.Left,ch)
	         //   }
	        }
	//    close(ch)
	    return
	*/
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := 0; i < 10; i++ {
		if <-c1 != <-c2 {
			return false
		}
	}
	return true

}

func main() {
	t := tree.New(1)
	t2 := tree.New(1)
	fmt.Println(t)
	fmt.Println(t2)
	c := make(chan int)
	go Walk(t, c)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println(Same(t, t2))
}
