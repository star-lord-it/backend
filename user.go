package main

var exp = 0
var level = 1

func UpLevel(e int) {
	exp += e

	level = (exp / 10) + 1
}
