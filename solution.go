package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
var SidesTriangle uint8 = 3
var SidesSquare uint8 = 4
var SidesCircle uint8 = 0

func CalcSquare(sideLen float64, sidesNum uint8) float64 {
	if sidesNum == SidesCircle{
		return math.Pi*sideLen*sideLen
	} else if sidesNum == SidesTriangle{
		return math.Sqrt(3)/4*sideLen*sideLen
	} else if sidesNum == SidesSquare{
		return sideLen*sideLen
	} else{
		return 0
	}
}
