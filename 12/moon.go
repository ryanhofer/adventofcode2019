package main

import "fmt"

type Moon struct {
	pos vec3
	vel vec3
}

func (m *Moon) energy() int {
	return m.potential() * m.kinetic()
}

func (m *Moon) potential() int {
	return abs(m.pos.x) + abs(m.pos.y) + abs(m.pos.z)
}

func (m *Moon) kinetic() int {
	return abs(m.vel.x) + abs(m.vel.y) + abs(m.vel.z)
}

func (m *Moon) String() string {
	return fmt.Sprintf("pos=%s vel=%s", m.pos, m.vel)
}
