package sayhi

import "fmt"

func Sayhi(name string) string {
	message := fmt.Sprintf("Hi there, %v!", name)
	return message
}
