package main

var text = []string{
	"Chunky and noisy,",
	"but with stars in their black feathers,",
	"they spring from the telephone wire",
	"and instantly",
}

type flock struct {
	birds []bird
}

func newFlock() flock {
	birds := make([]bird, 200)
	count := 0
	for i, line := range text {
		for j, char := range line {
			birds[count] = bird{char: char, x: j, y: i}
			count++
		}
	}
	return flock{
		birds: birds,
	}
}
