package main

import (
    "fmt"
    "time"
)

type Color int
type Model int
type Classification int

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

const (
    Heavy Classification = iota + 1
    Luxury
    OldSchool
)

type AddModelFunc func(m Model) AddColorFunc
type AddColorFunc func(c Color) AddPlateFunc
type AddPlateFunc func(plate string) Vehicle

type Vehicle struct {
    Plate          string
    Color          Color
    Model          Model
    Classification Classification
    Created        int64
}

func NewVehicleCurrying(cl Classification) AddModelFunc {
    return func(m Model) AddColorFunc {
        return func(c Color) AddPlateFunc {
            return func(plate string) Vehicle {
                return Vehicle{
                    Plate:          plate,
                    Color:          c,
                    Model:          m,
                    Classification: cl,
                    Created:        time.Now().Unix(),
                }
                
            }
            
        }
    }
}

func main() {
    corolla := NewVehicleCurrying(OldSchool)(Corolla)(Blue)("AJB-1908")
    fmt.Println(corolla)
    
    fox := NewVehicleCurrying(Luxury)(Fox)(Yellow)("AOI-0900")
    fmt.Println(fox)
    
    uno := NewVehicleCurrying(Heavy)(Uno)(Green)("AOI-0900")
    fmt.Println(uno)
}
