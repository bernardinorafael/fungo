package main

import (
	"fmt"
	"time"
)

type Category int
type Gender int

const (
	Clothing Category = iota + 1
	Sports
	Books
	Toys
)

const (
	Male Gender = iota + 1
	Female
	Unisex
)

type Product struct {
	Name     string
	Price    int
	Category Category
	Gender   Gender
	Created  int64
}

func ProductSpawner(c Category, g Gender) func(name string, price int) Product {
	return func(name string, price int) Product {
		return Product{
			Name:     name,
			Price:    price,
			Category: c,
			Gender:   g,
			Created:  time.Now().Unix(),
		}
	}
}

var (
	newMaleClothing = ProductSpawner(Clothing, Male)
	newFemaleSport  = ProductSpawner(Sports, Female)
	newBook         = ProductSpawner(Books, Unisex)
	newToy          = ProductSpawner(Toys, Unisex)
)

func main() {
	maleShirt := newMaleClothing("some male shirt", 10099)
	femaleSport := newFemaleSport("some female sport clothes", 29999)
	hp := newBook("some harry potter book", 5999)
	barbie := newToy("some barbie doll", 89990)

	fmt.Println(maleShirt)
	fmt.Println(femaleSport)
	fmt.Println(hp)
	fmt.Println(barbie)
}
