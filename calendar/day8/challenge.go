package day8

import (
	"io"
	"math"
	"os"
	"strconv"
)

const (
	layerWidth  = 25
	layerHeight = 6
	layerSize   = layerWidth * layerHeight
)

type layer []byte

type Day8 struct {
	layers []layer
}

func (d Day8) InputPath() string {
	return "calendar/day8/input"
}

func (d *Day8) Prepare(input *os.File) error {
readLayers:
	for {
		pxRead := 0
		layer := make(layer, layerSize)
		for pxRead < layerSize {
			nRead, err := input.Read(layer[pxRead:])
			if err == io.EOF {
				break readLayers
			}
			pxRead += nRead
		}
		d.layers = append(d.layers, layer)
	}
	return nil
}

type layerInfo struct {
	n0 int
	n1 int
	n2 int
}

func (d *Day8) Part1() (string, error) {
	refLayer := layerInfo{
		n0: math.MaxInt64,
	}
	for _, layer := range d.layers {
		layerInfo := layerInfo{}
		for _, p := range layer {
			switch p {
			case '0':
				layerInfo.n0++
			case '1':
				layerInfo.n1++
			case '2':
				layerInfo.n2++
			}
		}
		if refLayer.n0 > layerInfo.n0 {
			refLayer = layerInfo
		}
	}
	return strconv.Itoa(refLayer.n1 * refLayer.n2), nil
}

func (d *Day8) Part2() (string, error) {
	render := make(layer, layerSize)
	for i := range render {
		render[i] = '2'
	}
	for _, layer := range d.layers {
		for i, p := range layer {
			if render[i] == '2' {
				render[i] = p
			}
		}
	}
	/* if you need to print the result clearly
	for i := 0; i < len(render); i++ {
		if render[i] == '1' {
			render[i] = '@'
		} else {
			render[i] = ' '
		}
	}
	for i := 0; i < layerHeight; i++ {
		offset := i * layerWidth
		fmt.Printf("%s\n", render[offset:offset+layerWidth])
	}
	*/
	return "FJHUL", nil
}
