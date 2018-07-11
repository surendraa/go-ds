package list

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Single Linked List", func() {
		g.Describe("create a new list", func() {
			g.It("should return a list.Single", func() {
				s := Single{}
				arr := make([]int, 0)
				arr = append(arr, 1)
				s.New(arr)
				// a := 1
				// fmt.Printf("%T\n", s.Head.Value)
				// fmt.Printf("%T\n", a)
				// fmt.Println(s.Head.Value == 1)
				g.Assert(s.Head.Value).Equal(1)
			})
		})
	})
}
