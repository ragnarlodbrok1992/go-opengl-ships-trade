package camera

// Let's create camera - it will hold all the information needed
// to have camera be at some place and everything that LookAtV takes

// LookAt has three 3D vectors - eye, center and up

type Camera struct {
  Eye_x, Eye_y, Eye_z float32
  Center_x, Center_y, Center_z float32
  Up_x, Up_y, Up_z float32
}

func (cam* Camera) Move(x, y, z float32) {
  cam.Center_x += x
  cam.Center_y += y 
  cam.Center_z += z 
}

type CameraMovement int64

// @TODO add vector to center vector of camera to move it around using WASDQE keys

