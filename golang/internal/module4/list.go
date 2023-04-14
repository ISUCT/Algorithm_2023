package module4

import (
	"errors"
	"fmt"
)

type List struct {
	Head *Item
	Tail *Item
	Size int
}

type Item struct {
	Value int
	Next  *Item
}

func (l *List) Add(item Item) {
	if l.Size == 0 {
		l.Head = &item
	} else {
		l.Tail.Next = &item
	}
	l.Tail = &item
	l.Size += 1
}

func (l *List) GetByIndex(index int) (*Item, error) {
	if index >= l.Size {
		return nil, errors.New("Index out of bound")
	}
	if index == l.Size-1 {
		return l.Tail, nil
	}
	if index == 0 {
		return l.Head, nil
	}

	item := l.Head
	for i := 0; i <= index-1; i++ {
		item = item.Next
	}
	return item, nil
}

func Sample() {
	list := List{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
	item := Item{
		Value: 5,
		Next:  nil,
	}
	list.Add(item)
	list.Add(Item{
		Value: 6,
		Next:  nil,
	})
	list.Add(Item{
		Value: 7,
		Next:  nil,
	})

	fmt.Println(list.GetByIndex(0))
	fmt.Println(list.GetByIndex(1))
	fmt.Println(list.GetByIndex(2))

}
