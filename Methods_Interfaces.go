package main

type Vector struct {
	x int16
	y int16
	z int16
}
type Student interface {
	getName() string
	getHeight() int8
	getStudentId() int8
	getLoc() Coordinates
	getParkingId() int8
}
type Employee interface {
	getName() string
	getHeight() int8
	getEmployeeId() int8
	getLoc() Coordinates
	getParkingId() int8
}
type person struct {
	name        string
	height      int8
	employee_id int8
	location    Coordinates
	student_id  int8
	parking_id  int8
}

func (p person) getName() string {
	return p.name
}
func (p person) getHeight() int8 {
	return p.height
}
func (p person) getEmployeeId() int8 {
	return p.employee_id
}
func (p person) getLoc() Coordinates {
	return p.location
}
func (p person) getStudentId() int8 {
	return p.student_id
}
func (p person) getParkingId() int8 {
	return p.parking_id
}
func (v Vector) getX() int16 {
	return v.getX()
}
func (v Vector) getY() int16 {
	return v.y
}
func (c Coordinates) getZ() float64 {
	return c.z
}
func (c Coordinates) getX() float64 {
	return c.x
}
func (c Coordinates) getY() float64 {
	return c.y
}

type Coordinates struct {
	x float64
	y float64
	z float64
}
