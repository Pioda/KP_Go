package main

import "fmt"

func main()  {

	r := Rectangle{
		Width:10,
		Height:15,
		GeoLocation: GeoLocation{
			Color:"Red",
		},
	}
	c := Circle {
		Radius:42,
		GeoLocation: GeoLocation{
			Color: "Blue",
		},
	}

	t := Triangle {
		B:Position{
			X:10,
			Y:12,
		},
		C:Position{
			X:7,
			Y:5,
		},
		GeoLocation:GeoLocation{
			Color: "Green",
		},
	}
	var forms []interface{Paint}
	forms = append(forms, r)
	forms = append(forms, c)
	forms = append(forms, t)

	printElementList(forms)
}

func printElementList(elements []interface{Paint}) {
    for _, e := range elements {
        e.Paint()
    }
}

type Paint interface{
	Paint()
}

type Position struct {
	X float64
	Y float64
}

type GeoLocation struct {
	Color string
}

func (g GeoLocation) Paint(){
	fmt.Println(g.Color)
}

type Rectangle struct {
	Width float64
	Height float64
	GeoLocation
}

func (r Rectangle) Paint(){
	r.GeoLocation.Paint()
	fmt.Println("Width:", r.Width)
	fmt.Println("Height:", r.Height)
}

type Circle struct {
	Radius float64
	GeoLocation
}

func (c Circle) Paint(){
	c.GeoLocation.Paint()
	fmt.Println("Radius:", c.Radius)
}

type Triangle struct {
	B Position
	C Position
	GeoLocation
}

func (t Triangle) Paint(){
	t.GeoLocation.Paint()
	fmt.Println("B Y:", t.B.Y)	
	fmt.Println("B X:", t.B.X)
	fmt.Println("C X:", t.C.X)
	fmt.Println("C Y:", t.C.Y)
}

