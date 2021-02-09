package main

// For print
import "fmt"

type Video struct {
	name        string
	user        string
	description string
}

// go fmt hello-world.go format file
func main() {
	fmt.Println("Hello World from GO")

	// is mandatory '}' in the last property
	var video = Video{
		name:        "Video 0",
		user:        "Woz",
		description: "Description 0"}

	fmt.Println(video)
}
