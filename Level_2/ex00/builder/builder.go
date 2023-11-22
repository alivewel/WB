package main

type Builder interface {
	SetTypeHouse()
	SetColor()
	SetCountFloor()
	BuildHouse() House
}
