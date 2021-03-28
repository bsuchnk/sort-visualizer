package main

import (
	"fmt"
	"math/rand"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type barData struct {
	value int32
	ID    int32
}

func newData(n int32) []barData {
	vd := make([]barData, n)
	for i := int32(0); i < n; i++ {
		vd[i].value = rand.Int31()%(n*2-20) + 10
		vd[i].ID = i
	}
	return vd
}

type Bar struct {
	vao  uint32
	data barData
}

func newBar(bardata barData, barsN int32) *Bar {
	bar := new(Bar)
	bar.data = bardata

	l := float32(2. / float32(barsN))
	h := float32(bar.data.value) / float32(barsN)
	var vertices = []float32{
		0, 0, 0,
		0, h, 0,
		l, 0, 0,

		l, 0, 0,
		0, h, 0,
		l, h, 0,
	}

	bar.vao = makeVao(vertices)
	return bar
}

func (bar *Bar) draw() {
	gl.BindVertexArray(bar.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(18/3)) //18
}

type Visualizer struct {
	nElem  int32
	shader uint32
	data   []barData
	bars   []*Bar
}

func newVisualizer(n int32, shaderProgram uint32) *Visualizer {
	v := new(Visualizer)
	v.nElem = n
	v.shader = shaderProgram
	v.data = newData(n)
	v.bars = make([]*Bar, n)

	for i, d := range v.data {
		//fmt.Println(i)
		v.bars[i] = newBar(d, n)
	}

	//v.randomize()
	v.bubbleSort()

	return v
}

func (v *Visualizer) draw() {
	transform := gl.GetUniformLocation(v.shader, gl.Str("transform\x00"))
	for i, bar := range v.bars {
		gl.Uniform3f(transform, -1+float32(2)*float32(v.data[i].ID)/float32(v.nElem), -1, 0)
		bar.draw()
	}
}

func (v *Visualizer) randomize() {
	values := make([]int32, v.nElem)
	for i := int32(0); i < v.nElem; i++ {
		values[i] = i
	}
	for i := int32(0); i < v.nElem; i++ {
		r := rand.Int31() % (int32(len(values)) - i)
		v.data[i].value = values[r]
		values[r] = values[len(values)-1]
	}
}

//
func makeVao(vertices []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func (v *Visualizer) bubbleSort() {
	for i := 0; i < int(v.nElem); i++ {
		fmt.Printf("%d ", v.data[i].value)
	}
	fmt.Println()
	for i := 0; i < int(v.nElem); i++ {
		for j := 1; j < int(v.nElem); j++ {
			if v.data[j].value < v.data[j-1].value {
				v.data[j], v.data[j-1] = v.data[j-1], v.data[j]
			}
		}
	}

	for i := 0; i < int(v.nElem); i++ {
		fmt.Printf("%d ", v.data[i].value)
	}
}
