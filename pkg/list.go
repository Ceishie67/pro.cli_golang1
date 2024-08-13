package pkg

import (
	"github.com/fatih/color"
)

type ShoppingList struct {
	Items  []string
	Marked []bool
}

func (s *ShoppingList) AddItem(item string) {
	s.Items = append(s.Items, item)
	s.Marked = append(s.Marked, false)
}

func (s *ShoppingList) ShowItems() {
	C := color.New(color.FgCyan)
	M := color.New(color.FgHiMagenta)
	for i, item := range s.Items {
		if s.Marked[i] {
			C.Printf("   %d * %s ✅ \n", i+1, item)
		} else {
			M.Printf("   %d * %s\n", i+1, item)
		}

	}

}

func (s *ShoppingList) MarkItemAndUnmarkItem(index int) {

	if index < len(s.Items) && index >= 0 {

		if s.Marked[index] == false {
			s.Marked[index] = true

		} else {
			s.Marked[index] = false
		}

	}
}

func (s *ShoppingList) RemoveItem(index int) {

	if index < len(s.Items) && index >= 0 {

		s.Items = append(s.Items[:index], s.Items[index+1:]...)
		s.Marked = append(s.Marked[:index], s.Marked[index+1:]...)
	}
}

func (s *ShoppingList) ReorderList(fromIndex, toIndex int) {

	if fromIndex < len(s.Items) && toIndex < len(s.Items) && fromIndex != toIndex &&
		fromIndex < len(s.Marked) && toIndex < len(s.Marked) {

		item := s.Items[fromIndex]
		s.Items = append(s.Items[:fromIndex], s.Items[fromIndex+1:]...)
		s.Items = append(s.Items[:toIndex], append([]string{item}, s.Items[toIndex:]...)...)

		mark := s.Marked[fromIndex]
		s.Marked = append(s.Marked[:fromIndex], s.Marked[fromIndex+1:]...)
		s.Marked = append(s.Marked[:toIndex], append([]bool{mark}, s.Marked[toIndex:]...)...)
	}
}
