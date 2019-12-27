package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/imath"
)

type vec3 struct {
	x, y, z int
}

func (v vec3) cmp(w vec3) vec3 {
	return vec3{
		x: imath.Cmp(v.x, w.x),
		y: imath.Cmp(v.y, w.y),
		z: imath.Cmp(v.z, w.z),
	}
}

func (v vec3) add(w vec3) vec3 {
	return vec3{
		x: v.x + w.x,
		y: v.y + w.y,
		z: v.z + w.z,
	}
}

func (v vec3) negate() vec3 {
	return vec3{
		x: -v.x,
		y: -v.y,
		z: -v.z,
	}
}

func (v vec3) String() string {
	return fmt.Sprintf("<x=%d, y=%d, z=%d>", v.x, v.y, v.z)
}
