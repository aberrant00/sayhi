package sayhi

import "fmt"

func Sayhi(name string) string {
	message := fmt.Sprintf("Welcome faggy, %v!", name)

	return message
}
