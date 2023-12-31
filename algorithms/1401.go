package algorithms

func CheckOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	ret := false
	square_radius := float64(radius) * float64(radius)
	//enumerate: circle center in rectangle
	if xCenter <= x2 && xCenter >= x1 && yCenter <= y2 && yCenter >= y1 {
		ret = true
	}
	//enumerate: four corners in circle
	if squareDistanceFloat64(float64(x1), float64(y1), float64(xCenter), float64(yCenter)) <= square_radius {
		ret = true
	}
	if squareDistanceFloat64(float64(x1), float64(y2), float64(xCenter), float64(yCenter)) <= square_radius {
		ret = true
	}
	if squareDistanceFloat64(float64(x2), float64(y1), float64(xCenter), float64(yCenter)) <= square_radius {
		ret = true
	}
	if squareDistanceFloat64(float64(x2), float64(y2), float64(xCenter), float64(yCenter)) <= square_radius {
		ret = true
	}

	//enumerate: four midpoints in circle
	if squareDistanceFloat64(float64(x1)/2+float64(x2)/2, float64(y1), float64(xCenter), float64(yCenter)) <= square_radius {
		ret = true
	}
	if squareDistanceFloat64(float64(x1)/2+float64(x2)/2, float64(y2), float64(xCenter), float64(yCenter)) <= square_radius {
		ret = true
	}
	if squareDistanceFloat64(float64(x1), float64(y1)/2+float64(y2)/2, float64(xCenter), float64(yCenter)) <= square_radius {
		ret = true
	}
	if squareDistanceFloat64(float64(x2), float64(y1)/2+float64(y2)/2, float64(xCenter), float64(yCenter)) <= square_radius {
		ret = true
	}
	return ret
}
