package main

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	program := initOpenGL()
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)

	visualizer := newVisualizer(20, program)

	for !window.ShouldClose() {
		t := time.Now()

		draw(visualizer, window, program)

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

func draw(vis *Visualizer, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	//TODO: implement drawing
	vis.draw()

	glfw.PollEvents()
	window.SwapBuffers()
}
