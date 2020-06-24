// Algorithms file file
package cmd

import (
	"fmt"
	"math"
)

type stepper interface {
	step(x,y float64) stepper
	values(format string) []string
	ready() bool
}

type inPlaceStepper struct {
	xsaved, ysaved float64 // saved values for printing
	xstart, ystart, xprev, yprev float64
	anglemin, anglemax           float64
	stepNo                       int
	rdy                          bool
}

func (a inPlaceStepper) values(format string) []string {
	return []string{fmt.Sprintf(format,a.xstart),fmt.Sprintf(format,a.ystart)}
}

func (a inPlaceStepper) ready() bool {
	return a.rdy
}

func (a inPlaceStepper) step(x, y float64) stepper {
	a.rdy = false
	if a.stepNo == 0 {
		a.xstart, a.xprev, a.ystart, a.yprev = x, x, y, y
		a.stepNo++
		a.rdy = true
		return a
	}
	Dx, Dy := x-a.xstart, y-a.ystart
	loangle, hiangle, angle := math.Atan2(Dy-tolerance, Dx), math.Atan2(Dy+tolerance, Dx), math.Atan2(Dy, Dx)
	if a.stepNo == 1 {
		a.anglemin, a.anglemax = loangle, hiangle
	}
	// condition set to trigger on NaN too
	if !(angle >= a.anglemin) || !(angle <= a.anglemax) { // Finding a value steps our algoritm twice
		a.xstart, a.ystart, a.xprev, a.yprev = a.xprev, a.yprev, x, y
		Dx, Dy = x-a.xstart, y-a.ystart
		a.anglemin, a.anglemax = math.Atan2(Dy-tolerance, Dx), math.Atan2(Dy+tolerance, Dx)
		a.stepNo++
		a.rdy = true
		return a
	}
	if loangle >= a.anglemin {
		a.anglemin = loangle
	}
	if hiangle <= a.anglemax {
		a.anglemax = hiangle
	}
	a.stepNo++
	a.xprev, a.yprev = x, y
	return a
}

// Interpolator
type interpStepper struct {

	xstart, ystart, xprev, yprev float64
	anglemin, anglemax           float64
	stepNo                       int
	rdy                          bool
}

func (a interpStepper) step(x, y float64) stepper {
	a.rdy = false
	if a.stepNo == 0 {
		a.xstart, a.xprev, a.ystart, a.yprev = x, x, y, y
		a.stepNo++
		a.rdy = true
		return a
	}
	Dx, Dy := x-a.xstart, y-a.ystart
	loangle, hiangle := math.Atan2(Dy-tolerance, Dx), math.Atan2(Dy+tolerance, Dx)
	if a.stepNo == 1 {
		a.anglemin, a.anglemax = loangle, hiangle
	}
	// should admit NaN values
	if !(a.anglemax >= loangle ) || !(a.anglemin <= hiangle) { // Finding a value steps our algoritm twice
		a.ystart = a.ystart + (a.xprev - a.xstart) * (math.Tan(a.anglemax) + math.Tan(a.anglemin)) / 2 // interpolator
		a.xstart, a.xprev, a.yprev = a.xprev, x, y
		Dx, Dy = x-a.xstart, y-a.ystart
		a.anglemin, a.anglemax = math.Atan2(Dy-tolerance, Dx), math.Atan2(Dy+tolerance, Dx)
		a.stepNo++
		a.rdy = true
		return a
	}
	if loangle >= a.anglemin {
		a.anglemin = loangle
	}
	if hiangle <= a.anglemax {
		a.anglemax = hiangle
	}
	a.stepNo++
	a.xprev, a.yprev = x, y
	return a
}

func (a interpStepper) values(format string) []string {
	return []string{fmt.Sprintf(format,a.xstart),fmt.Sprintf(format,a.ystart)}
}

func (a interpStepper) ready() bool {
	return a.rdy
}