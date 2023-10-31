package camera

// Let's create camera - it will hold all the information needed
// to have camera be at some place and everything that LookAtV takes

// LookAt has three 3D vectors - eye, center and up

type Camera struct {
  eye_x, eye_y, eye_z float64
  center_x, center_y, center_z float64
  up_x, up_y, up_z float64
}


