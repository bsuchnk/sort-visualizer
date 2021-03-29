package main

var (
	vertexShaderSource = `
#version 330

in vec3 vp;
out vec3 col_pos;

uniform vec3 transform;
uniform float bar;

void main() {
	gl_Position = vec4(vp+transform, 1.0);

	if (bar >= 0)
		col_pos = vec3(bar,1-bar,1);
	else
		col_pos = vec3(1,1,0.5);
}
` + "\x00"

	fragmentShaderSource = `
#version 330
out vec4 frag_colour;
in vec3 col_pos;
void main() {
	frag_colour = vec4(col_pos, 1);
}
` + "\x00"
)
