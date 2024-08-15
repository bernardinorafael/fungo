package main

import (
	"fmt"
	"time"
)

type (
	Plate string
	Color int
	Model int
)

const (
	Blue Color = iota + 1
	Green
	Yellow
)

const (
	Corolla Model = iota + 1
	Fox
	Uno
)

type AddModelFunc func(m Model) AddPlateFunc
type AddPlateFunc func(p Plate) Vehicle

type Vehicle struct {
	Plate   Plate
	Color   Color
	Model   Model
	Created int64
}

func NewVehicleCurrying(c Color) AddModelFunc {
	return func(m Model) AddPlateFunc {
		return func(p Plate) Vehicle {
			return Vehicle{
				Plate: p,
				Model: m, Color: c,
				Created: time.Now().Unix(),
			}
		}
	}
}

func main() {
	corolla := NewVehicleCurrying(Blue)(Corolla)("A8I-90OK")
	fox := NewVehicleCurrying(Yellow)(Fox)("AOI-0900")
	uno := NewVehicleCurrying(Green)(Uno)("AOI-0900")

	fmt.Println(uno)
	fmt.Println(corolla)
	fmt.Println(fox)
}
