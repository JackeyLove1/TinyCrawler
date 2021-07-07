package model

import "fmt"

type BookDetail struct {
	Name      string
	Author    string
	Publicer  string
	Bookpages string
	Price     string
	Score     string
}

func (b *BookDetail) String() string {
	return fmt.Sprintf("name: %s, author: %s, publicer: %s, bookpages: %s, price: %s, score: %s\n",
		b.Name, b.Author, b.Publicer, b.Bookpages, b.Price, b.Score)
}
