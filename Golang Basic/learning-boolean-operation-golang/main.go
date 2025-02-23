package main

import "fmt"

func main() {
	finalScore := 90
	attendance := 80

	passFinalGrade := finalScore > 80
	passFinalAttendance := attendance > 80

	passed1 := passFinalGrade && passFinalAttendance
	passed2 := passFinalGrade || passFinalAttendance

	fmt.Println(passed1)
	fmt.Println(passed2)
}
