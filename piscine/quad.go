package piscine

import (
	"fmt"
	"strings"
)

func QuadA(x, y int) {
	characters := map[string]rune{
		"topLeft":     'o',
		"topRight":    'o',
		"bottomLeft":  'o',
		"bottomRight": 'o',
		"horizontal":  '-',
		"vertical":    '|',
	}

	fmt.Print(getQuad(x, y, characters))
}

func QuadB(x, y int) {
	characters := map[string]rune{
		"topLeft":     '/',
		"topRight":    '\\',
		"bottomLeft":  '/',
		"bottomRight": '\\',
		"horizontal":  '*',
		"vertical":    '*',
	}

	fmt.Print(getQuad(x, y, characters))
}

func QuadC(x, y int) {
	characters := map[string]rune{
		"topLeft":     'A',
		"topRight":    'A',
		"bottomLeft":  'C',
		"bottomRight": 'C',
		"horizontal":  'B',
		"vertical":    'B',
	}

	fmt.Print(getQuad(x, y, characters))
}

func QuadD(x, y int) {
	characters := map[string]rune{
		"topLeft":     'A',
		"topRight":    'C',
		"bottomLeft":  'A',
		"bottomRight": 'C',
		"horizontal":  'B',
		"vertical":    'B',
	}

	fmt.Print(getQuad(x, y, characters))
}

func QuadE(x, y int) {
	characters := map[string]rune{
		"topLeft":     'A',
		"topRight":    'C',
		"bottomLeft":  'C',
		"bottomRight": 'A',
		"horizontal":  'B',
		"vertical":    'B',
	}

	fmt.Print(getQuad(x, y, characters))
}

func getQuad(x, y int, characters map[string]rune) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	var quad strings.Builder

	for row := range y {
		for col := range x {
			switch {
			case row == 0 && col == 0:
				quad.WriteRune(characters["topLeft"])
			case row == 0 && col == x-1:
				quad.WriteRune(characters["topRight"])
			case row == y-1 && col == 0:
				quad.WriteRune(characters["bottomLeft"])
			case row == y-1 && col == x-1:
				quad.WriteRune(characters["bottomRight"])
			case row == 0 || row == y-1:
				quad.WriteRune(characters["horizontal"])
			case col == 0 || col == x-1:
				quad.WriteRune(characters["vertical"])
			default:
				quad.WriteRune(' ')
			}
		}
		quad.WriteRune('\n')
	}

	return quad.String()
}
