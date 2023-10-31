package camera

// Let's create camera - it will hold all the information needed
// to have camera be at some place and everything that LookAtV takes

// LookAt has three 3D vectors - eye, center and up

type Camera struct {
  Eye_x, Eye_y, Eye_z float64
  Center_x, Center_y, Center_z float64
  Up_x, Up_y, Up_z float64
}


