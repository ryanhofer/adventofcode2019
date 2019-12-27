package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/imath"
)

type Moon struct {
	pos vec3
	vel vec3
}

func (m *Moon) energy() int {
	return m.potential() * m.kinetic()
}

func (m *Moon) potential() int {
	return imath.Abs(m.pos.x) + imath.Abs(m.pos.y) + imath.Abs(m.pos.z)
}

func (m *Moon) kinetic() int {
	return imath.Abs(m.vel.x) + imath.Abs(m.vel.y) + imath.Abs(m.vel.z)
}

func (m *Moon) String() string {
	return fmt.Sprintf("pos=%s vel=%s", m.pos, m.vel)
}
