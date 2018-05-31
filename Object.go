package main

import (
	"fmt"

	"./src/bean"
)

func main() {
	var newRoom room
	newRoom.bed = "new"
	newRoom.door = "wide"
	newRoom.glass = "water"
	newRoom.wall = "tall"
	fmt.Println(newRoom.bed)
	//	var niceRoom room
	//	niceRoom.glass = "玻璃"
	//	niceRoom.bed = "软软的床"
	//	niceRoom.wall = "高高的墙"
	//	niceRoom.door = "结实的门"
	//	fmt.Println(niceRoom.door)
}
