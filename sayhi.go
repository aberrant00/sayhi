package sayhi

import "fmt"

func Sayhi(name string) string {
	message := fmt.Sprintf("Well well well if it isnt heh, %v!", name)

	return message
}
