package main

import "fmt"

type vec3 struct {
	x, y, z int
}

func (v vec3) cmp(w vec3) vec3 {
	return vec3{
		x: cmp(v.x, w.x),
		y: cmp(v.y, w.y),
		z: cmp(v.z, w.z),
	}
}

func (v vec3) add(w vec3) vec3 {
	return vec3{
		x: v.x + w.x,
		y: v.y + w.y,
		z: v.z + w.z,
	}
}

func (v vec3) String() string {
	return fmt.Sprintf("<x=%d, y=%d, z=%d>", v.x, v.y, v.z)
}
