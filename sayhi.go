package sayhi

import "fmt"

func Sayhi(name string) string {
	message := fmt.Sprintf("Welcome there faggy, %v!", name)

	return message
}
