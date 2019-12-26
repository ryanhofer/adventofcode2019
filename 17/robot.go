package main

import . "github.com/ryanhofer/adventofcode2019/geom"

type Robot struct {
	pos Vec
	dir Cardinal
}

func (r *Robot) TurnLeft() {
	r.dir = r.dir.RotateCCW()
}

func (r *Robot) TurnRight() {
	r.dir = r.dir.RotateCW()
}

func (r *Robot) MoveForward() {
	r.pos = r.Forward()
}

func (r *Robot) Left() Vec {
	v := r.dir.RotateCCW().Vec()
	return r.pos.Add(v)
}

func (r *Robot) Right() Vec {
	v := r.dir.RotateCW().Vec()
	return r.pos.Add(v)
}

func (r *Robot) Forward() Vec {
	return r.pos.Add(r.dir.Vec())
}

func (r *Robot) Rune() rune {
	switch r.dir {
	case North:
		return RobotN
	case East:
		return RobotE
	case South:
		return RobotS
	case West:
		return RobotW
	default:
		return '?'
	}
}

func isRobot(r rune) bool {
	switch r {
	case RobotN:
		return true
	case RobotE:
		return true
	case RobotS:
		return true
	case RobotW:
		return true
	case RobotFalling:
		return true
	default:
		return false
	}
}
