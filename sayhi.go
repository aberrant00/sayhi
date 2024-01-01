package sayhi

import "fmt"

func Sayhi(name string) string {
	message := fmt.Sprintf("Hi there bro, %v!", name)
	return message
}
