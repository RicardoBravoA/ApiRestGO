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
	var video = getVideo("User 0", "Woz", "Description 0")

	fmt.Println(video)

	fmt.Println(delivery(5))
}

func getVideo(name string, user string, description string) Video {
	// is mandatory '}' in the last property
	return Video{
		name:        name,
		user:        user,
		description: description}
}


// closures
func delivery(items int) (string, float32) {

	const price float32 = 10

	//closure
	total := func() float32 {
		return float32(items) * price
	}

	return "The total is:", total()

}
