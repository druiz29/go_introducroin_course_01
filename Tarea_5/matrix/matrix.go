package matrix

import (
	"errors"
	"fmt"
)

type Matrix struct {
	Values    [][]float64
	Height    int  // filas
	Width     int  // columnas
	Quadratic bool // si es cuadrática
}

// New recibe N filas, cada fila es un []float64
func New(rows ...[]float64) (*Matrix, error) {
	if len(rows) == 0 {
		return nil, errors.New("matrix cannot be empty")
	}

	// Validar que la primera fila tenga columnas
	width := len(rows[0])
	if width == 0 {
		return nil, errors.New("matrix must have at least one column")
	}

	// Validar que todas las filas tengan el mismo número de columnas
	for i, r := range rows {
		if len(r) != width {
			return nil, fmt.Errorf("row %d has length %d, expected %d", i, len(r), width)
		}
	}

	height := len(rows)

	// Copia defensiva (evita que desde afuera modifiquen la matriz sin querer)
	values := make([][]float64, height)
	for i := range rows {
		values[i] = make([]float64, width)
		copy(values[i], rows[i])
	}

	m := &Matrix{
		Values:    values,
		Height:    height,
		Width:     width,
		Quadratic: height == width,
	}

	return m, nil
}

func (m *Matrix) Print() {
	// Imprime filas como: [ 2 7 8 ]
	for _, row := range m.Values {
		fmt.Print("[ ")
		for _, v := range row {
			// %g evita .0 si no hace falta (2 en vez de 2.0)
			fmt.Printf("%g ", v)
		}
		fmt.Println("]")
	}

	// Dimensión: "3 x 3"
	fmt.Printf("%d x %d\n", m.Height, m.Width)

	// Tal cual el output pedido
	fmt.Println("Cuadratic: ", m.Quadratic)
}
