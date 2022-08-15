package log

import "fmt"

func LogInfo(f string, args ...interface{}) {
	msg := fmt.Sprintf(f, args...)
	fmt.Println(msg)
}
