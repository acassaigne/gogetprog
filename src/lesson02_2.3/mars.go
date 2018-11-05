package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	const spaceXSpeed = 100800/3600
	var distance = 56000000 // km

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	distance = 96300000
	fmt.Println(distance/spaceXSpeed, "seconds")
	fmt.Println(distance/spaceXSpeed/3600/24, "days")

}
