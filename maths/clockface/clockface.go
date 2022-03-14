package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

// Since every seconds hand rev --> minutes hands move one tick (1/60)
// Add it to radians minute hand at
// (minutes resolution) + (seconds resolution)
func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}
func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) +
		(math.Pi / (6 / float64(t.Hour()%12)))
}
func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}
func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
