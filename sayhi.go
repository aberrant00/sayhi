package sayhi

import "fmt"

func Sayhi(name string) string {
	message := fmt.Sprintf("Hi, %v!", name)
	return message
}
