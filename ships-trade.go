package main

import (
  "fmt"
  "go/build"
  "image"
  "image/draw"
  _ "image/png"
  "log"
  "os"
  "runtime"
  "strings"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/glfw/v3.3/glfw"
  "github.com/go-gl/mathgl/mgl32"
)

const WINDOW_WIDTH = 1024
const WINDOW_HEIGHT = 768

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
    log.Fataln("Failed to initialize glfw:", err)
  }
  defer glfw.Terminate();

  glfw.WindowHint(glfw.Resizable, glfw.False)
  glfw.WindowHint(glfw.ContextVersionMajor, 4)
  glfw.WindowHint(glfw.ContextVersionMinor, 1)
  glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
  glfw.WindowHint(glfw.OpenGLForwardCompativle, glfw.True)

  window, err := glfw.CreateWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Ships Trade", nil, nil)
  if err != nil {
    panic(err);
  }
  window.MakeContextCurrent();


}

func init() {
  dir, err := importPathToDir("github.com/go-gl/example/gl41core-cube")
  if err != nil {
    log.Fatalln("Unable to find Go package in your GOPATH, it's needed to load assets:", err)
  }
  err = os.Chdir(dir)
  if err != nil {
    log.Panicln("os.Chdir:", err)
  }
}

func importPathToDir(importPath string) (string, error) {
  p, err := build.Import(importPath, "", build.FindOnly)
  if err != nil {
    return "", err
  }
  return p.Dir, nil
}

