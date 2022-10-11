package main

import "fmt"

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		switch t {
		case Inch:
			value = value / 2.54

		default:
			value = value * 2.54
		}
	}
	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

func main() {
	Rectangle := NewRectangle(50, CM)
	Table := NewTable(120, 60, 70, 4, Inch)

	Volume(Rectangle, Inch)
	Volume(Table, CM)

	Car1 := NewCar("BMW", "X6", 250, 200, 400, 150, 190, CM)
	Car2 := NewCar("Mercedes", "Bens", 300, 220, 450, 160, 200, CM)
	Car3 := NewCar("Dodge", "Durango", 270, 270, 470, 170, 170, Inch)
	CarInfo(Car1, CM)
	CarInfo(Car2, CM)
	CarInfo(Car3, Inch)
}

type Rectangle struct {
	length Unit
}

type Table struct {
	length Unit
	width  Unit
	height Unit
	leg    int
}

func (a Rectangle) Length() Unit {
	return a.length
}
func (a Rectangle) Width() Unit {
	return a.length
}
func (a Rectangle) Height() Unit {
	return a.length
}
func NewRectangle(length float64, T UnitType) Rectangle {
	return Rectangle{NewUnit(length, T)}
}

func (a Table) Length() Unit {
	return a.length
}
func (a Table) Width() Unit {
	return a.width
}
func (a Table) Height() Unit {
	return a.height
}
func NewTable(length, width, height float64, leg int, T UnitType) Table {
	return Table{NewUnit(length, T), NewUnit(width, T), NewUnit(height, T), 4}
}

func NewUnit(value float64, T UnitType) Unit {
	return Unit{value, T}
}

func Volume(d Dimensions, T UnitType) {
	fmt.Printf("Объем: %f %s^3\n", d.Height().Get(T)*d.Length().Get(T)*d.Width().Get(T), T)
}

type Car struct {
	brand       string
	model       string
	maxSpeed    int
	enginePower int
	Table
}

func (a Car) Brand() string {
	return a.brand
}
func (a Car) Model() string {
	return a.model
}
func (a Car) MaxSpeed() int {
	return a.maxSpeed
}
func (a Car) EnginePower() int {
	return a.enginePower
}
func (a Car) Dimensions() Dimensions {
	var d Dimensions = a.Table
	return d
}

func NewCar(brand, model string, maxSpeed, enginePower int, length, width, height float64, T UnitType) Car {
	return Car{brand, model, maxSpeed, enginePower, NewTable(length, width, height, 0, T)}
}

func CarInfo(a Auto, T UnitType) {
	d := a.Dimensions()
	fmt.Printf("Брэнд: %s\n", a.Brand())
	fmt.Printf("Модель: %s\n", a.Model())
	fmt.Printf("Макс скорость: %d\n", a.MaxSpeed())
	fmt.Printf("Кол-во лошадиных сил: %d\n", a.EnginePower())
	Volume(d, T)
}
