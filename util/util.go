package util

import (
	"fmt"
)

func Blue() {
	fmt.Print("\033[34m")
}

func Red() {
	fmt.Print("\033[31m")
}

func Green() {
	fmt.Print("\033[32m")
}

func Yellow() {
	fmt.Print("\033[33m")
}

func Magenta() {
	fmt.Print("\033[35m")
}

func Cyan() {
	fmt.Print("\033[36m")
}

func Reset() {
	fmt.Print("\033[0m")
}
