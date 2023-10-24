package main

import (
  "fmt"
  "image"
  "image/draw"
  _ "image/png"
  "log"
  "os"
  "runtime"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/glfw/v3.3/glfw"
  "github.com/go-gl/mathgl/mgl32"

  "go-opengl-ships-trade/src/graphics"
  // "go-opengl-ships-trade/src/helpers"
)

const WINDOW_WIDTH = 1024
const WINDOW_HEIGHT = 768

var cubeVertices = []float32{
	//  X, Y, Z, U, V
	// Bottom
	-1.0, -1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,

	// Top
	-1.0, 1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, 1.0, 1.0, 1.0,

	// Front
	-1.0, -1.0, 1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,

	// Back
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 1.0,

	// Left
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,

	// Right
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
}


var vertexShader = `
  #version 330

  uniform mat4 projection;
  uniform mat4 camera;
  uniform mat4 model;

  in vec3 vert;
  in vec2 vertTexCoord;

  out vec2 fragTexCoord;

  void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera * model * vec4(vert, 1);
  }
` + "\x00"

var fragmentShader = `
  #version 330

  uniform sampler2D tex;

  in vec2 fragTexCoord;

  out vec4 outputColor;

  void main() {
    outputColor = texture(tex, fragTexCoord);
  }
` + "\x00"

func key_callback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
  // DEBUG
  fmt.Println("Pressing key!")

  // Escaping engine
  if key == glfw.KeyEscape && action == glfw.Press {
    window.SetShouldClose(true)
  }
} 

func init() {
  // GLFW event handling run on the main OS thread
  runtime.LockOSThread() // TODO check that that does
}

func main() {
  // Init - deinit messages
  fmt.Println("Ships trade!");
  defer fmt.Println("Goodbye.");

  // Engine start - glfw
  if err := glfw.Init(); err != nil {
    log.Fatalln("Failed to initialize glfw:", err)
  }
  defer glfw.Terminate();

  glfw.WindowHint(glfw.Resizable, glfw.False)
  glfw.WindowHint(glfw.ContextVersionMajor, 4)
  glfw.WindowHint(glfw.ContextVersionMinor, 1)
  glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
  glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

  window, err := glfw.CreateWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Ships Trade", nil, nil)
  if err != nil {
    panic(err);
  }
  window.MakeContextCurrent();

  if err := gl.Init(); err != nil {
    panic(err)
  }

  // Set GLFW callback functions
  window.SetKeyCallback(key_callback);

  version := gl.GoStr(gl.GetString(gl.VERSION))
  fmt.Println("OpenGL version", version)

  // Configure the vertex and fragment shaders
  program, err := graphics.NewProgram(vertexShader, fragmentShader)
  if err != nil {
    panic(err)
  }

  gl.UseProgram(program)

  // Projection camera model matrices my favourite!
  projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(WINDOW_WIDTH) / WINDOW_HEIGHT, 0.1, 1000.0)
  projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
  gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

  camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
  cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
  gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

  model := mgl32.Ident4()
  modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
  gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

  textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
  gl.Uniform1i(textureUniform, 0)

  gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

  // Load the texture
  texture, err := newTexture("assets/placeholder_texture.png")
  if err != nil {
    log.Fatalln(err)
  }

  // Configure the vertex data
  var vao uint32
  gl.GenVertexArrays(1, &vao)
  gl.BindVertexArray(vao)

  var vbo uint32
  gl.GenBuffers(1, &vbo)
  gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices) * 4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)

  vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
  gl.EnableVertexAttribArray(vertAttrib)
  gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 5 * 4, 0)

  texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
  gl.EnableVertexAttribArray(texCoordAttrib)
  gl.VertexAttribPointerWithOffset(texCoordAttrib, 2, gl.FLOAT, false, 5 * 4, 3 * 4)

  // Configure global settings
  gl.Enable(gl.DEPTH_TEST)
  gl.DepthFunc(gl.LESS)
  gl.ClearColor(1.0, 1.0, 1.0, 1.0)

  angle := 0.0
  previousTime := glfw.GetTime()

  for !window.ShouldClose() {
    // Update
    time := glfw.GetTime()
    elapsed := time - previousTime
    previousTime = time

    // Simulate
    angle += elapsed
    model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})

    // Render
    // Clear screen
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

    gl.UseProgram(program)
    gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

    gl.BindVertexArray(vao)

    gl.ActiveTexture(gl.TEXTURE0)
    gl.BindTexture(gl.TEXTURE_2D, texture)

    gl.DrawArrays(gl.TRIANGLES, 0, 6 * 2 * 3)

    // Maintenance
    window.SwapBuffers()
    glfw.PollEvents()
  }

}

func newTexture(file string) (uint32, error) {
  // DEBUG
  path, err := os.Getwd()
  if err != nil {
    log.Println(err)
  }

  fmt.Println(path)
  
  imgFile, err := os.Open(file)
  if err != nil {
    return 0, fmt.Errorf("texture %q not found on disk: %v", file, err)
  }

  img, _, err := image.Decode(imgFile)
  if err != nil {
    return 0, err
  }

  rgba := image.NewRGBA(img.Bounds())
  if rgba.Stride != rgba.Rect.Size().X*4 {
    return 0, fmt.Errorf("unsopported stride")
  }

  draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

  var texture uint32
  gl.GenTextures(1, &texture)
  gl.ActiveTexture(gl.TEXTURE0)
  gl.BindTexture(gl.TEXTURE_2D, texture)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
  gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
  gl.TexImage2D(
    gl.TEXTURE_2D,
    0,
    gl.RGBA,
    int32(rgba.Rect.Size().X),
    int32(rgba.Rect.Size().Y),
    0,
    gl.RGBA,
    gl.UNSIGNED_BYTE,
    gl.Ptr(rgba.Pix))

  return texture, nil
}


