package utils

import "github.com/johnfercher/maroto/v2/pkg/props"

// Colors used in the PDF library
var Colors = struct {
	Blue     *props.Color
	DarkGray *props.Color
	Gray     *props.Color
	Red      *props.Color
	White    *props.Color
}{
	Blue:     &props.Color{Red: 10, Green: 10, Blue: 150},
	DarkGray: &props.Color{Red: 55, Green: 55, Blue: 55},
	Gray:     &props.Color{Red: 240, Green: 240, Blue: 240},
	Red:      &props.Color{Red: 150, Green: 10, Blue: 10},
	White:    &props.Color{Red: 255, Green: 255, Blue: 255},
}
