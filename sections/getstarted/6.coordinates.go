package getstarted

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/raedatoui/glutils"
	"github.com/raedatoui/learn-opengl-golang/sections"
)

type HelloCoordinates struct {
	sections.BaseSketch
	shader             glutils.Shader
	va glutils.VertexArray
	texture1, texture2 uint32
	transform          mgl32.Mat4
	cubePositions      []mgl32.Mat4
	rotationAxis       mgl32.Vec3
}

func (hc *HelloCoordinates) InitGL() error {
	hc.Name = "6. Coordinate Systems"

	var err error
	hc.shader, err = glutils.NewShader(
		"_assets/getting_started/6.coordinates/coordinate.vs",
		"_assets/getting_started/6.coordinates/coordinate.frag", "")
	if err != nil {
		return err
	}

	vertices := []float32{
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,

		-0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}
	attr := make(glutils.AttributesMap)
	attr[hc.shader.Attributes["position"]] = [2]int{0, 3}
	attr[hc.shader.Attributes["texCoord"]] = [2]int{3, 2}

	hc.va = glutils.VertexArray{
		Data: vertices,
		Stride: 5,
		DrawMode: gl.STATIC_DRAW,
		Normalized: false,
		Attributes: attr,
	}
	hc.rotationAxis = mgl32.Vec3{1.0, 0.3, 0.5}.Normalize()
	hc.cubePositions = []mgl32.Mat4{
		mgl32.Translate3D(0.0, 0.0, 0.0),
		mgl32.Translate3D(2.0, 5.0, -15.0),
		mgl32.Translate3D(-1.5, -2.2, -2.5),
		mgl32.Translate3D(-3.8, -2.0, -12.3),
		mgl32.Translate3D(2.4, -0.4, -3.5),
		mgl32.Translate3D(-1.7, 3.0, -7.5),
		mgl32.Translate3D(1.3, -2.0, -2.5),
		mgl32.Translate3D(1.5, 2.0, -2.5),
		mgl32.Translate3D(1.5, 0.2, -1.5),
		mgl32.Translate3D(-1.3, 1.0, -1.5),
	}


	// Texture 1
	if tex, err := glutils.NewTexture(gl.REPEAT, gl.REPEAT, gl.LINEAR, gl.LINEAR, "_assets/images/container.png"); err != nil {
		return err
	} else {
		hc.texture1 = tex
	}

	// Texture 2
	if tex, err := glutils.NewTexture(gl.REPEAT, gl.REPEAT, gl.LINEAR, gl.LINEAR, "_assets/images/awesomeface.png"); err != nil {
		return err
	} else {
		hc.texture2 = tex
	}

	return nil
}

func (hc *HelloCoordinates) Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(hc.Color32.R, hc.Color32.G, hc.Color32.B, hc.Color32.A)

	// Bind Textures using texture units
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, hc.texture1)
	loc1 := gl.GetUniformLocation(hc.shader, gl.Str("ourTexture1\x00"))
	gl.Uniform1i(loc1, 0)

	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, hc.texture2)
	loc2 := gl.GetUniformLocation(hc.shader, gl.Str("ourTexture2\x00"))
	gl.Uniform1i(loc2, 1)

	// Activate shader
	gl.UseProgram(hc.shader)

	// Create transformations
	view := mgl32.Translate3D(0.0, 0.0, -3.0)
	projection := mgl32.Perspective(45.0, sections.RATIO, 0.1, 100.0)
	// Get their uniform location
	modelLoc := gl.GetUniformLocation(hc.shader, gl.Str("model\x00"))
	viewLoc := gl.GetUniformLocation(hc.shader, gl.Str("view\x00"))
	projLoc := gl.GetUniformLocation(hc.shader, gl.Str("projection\x00"))
	// Pass the matrices to the shader
	gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])
	// Note: currently we set the projection matrix each frame,
	// but since the projection matrix rarely changes it's often best practice to set it outside the main loop only once.
	gl.UniformMatrix4fv(projLoc, 1, false, &projection[0])

	// Draw container
	gl.BindVertexArray(hc.vao)

	for i := 0; i < 10; i++ {
		// Calculate the model matrix for each object and pass it to shader before drawing
		model := hc.cubePositions[i]

		angle := float32(glfw.GetTime()) * float32(i+1)

		model = model.Mul4(mgl32.HomogRotate3D(angle, hc.rotationAxis))
		gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])
		gl.DrawArrays(gl.TRIANGLES, 0, 36)
	}
	gl.BindVertexArray(0)
}

func (hc *HelloCoordinates) Close() {
	hc.shader.Delete()
	hc.va.Delete()
}
