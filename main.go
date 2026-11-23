
	//seccion: 307C2
	//Eduardo Miranda 32282061
	//Gabriel Aberg Forney 31229711
	//Valeria Fong 33054233
	//Juan Diaz 32487540

package main

import (
	"fmt"
	"math"
)

type Figura interface {
	Area() float64
	Perimetro() float64
	Nombre() string
}

type BaseFigura struct {
	nombre string
}

func (b BaseFigura) Nombre() string {
	return b.nombre
}
type Rectangulo struct {
	BaseFigura // Composición embebida
	largo      float64
	ancho      float64
}

func NuevoRectangulo(largo, ancho float64) Rectangulo {
	return Rectangulo{
		BaseFigura: BaseFigura{nombre: "Rectángulo"},
		largo:      largo,
		ancho:      ancho,
	}
}

func (r Rectangulo) Area() float64 {
	return r.largo * r.ancho
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.largo + r.ancho)
}


type Circulo struct {
	BaseFigura // Composición embebida
	radio      float64
}

func NuevoCirculo(radio float64) Circulo {
	return Circulo{
		BaseFigura: BaseFigura{nombre: "Círculo"},
		radio:      radio,
	}
}

func (c Circulo) Area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.radio
}


type Triangulo struct {
	BaseFigura // Composición embebida
	a, b, c    float64
}

func NuevoTriangulo(a, b, c float64) Triangulo {
	return Triangulo{
		BaseFigura: BaseFigura{nombre: "Triángulo"},
		a:          a,
		b:          b,
		c:          c,
	}
}

func (t Triangulo) Area() float64 {
	// Fórmula de Area: s = (a + b + c) / 2
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t Triangulo) Perimetro() float64 {
	return t.a + t.b + t.c
}
func ImprimirFigura(f Figura) {
	fmt.Printf("Figura: %-11s | Área: %6.2f | Perímetro: %6.2f\n",
		f.Nombre(), f.Area(), f.Perimetro())
}
func main() {

	rectangulo := NuevoRectangulo(3, 4)
	circulo := NuevoCirculo(5)
	triangulo := NuevoTriangulo(3, 4, 5)


	figuras := []Figura{rectangulo, circulo, triangulo}

	for _, f := range figuras {
		ImprimirFigura(f)
	}
}
