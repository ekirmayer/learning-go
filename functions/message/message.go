package message

import "fmt"

// THe name of the function has to be in capital to make it public
func Message(name string, message string) string {
	return fmt.Sprintf("%s, %s", message, name)
}

func New_Message(name string, message string) (result string) {
	result = fmt.Sprintf("%s, %s", message, name)
	return result
}

func Very_New_Message(name string, message string) (result string, index string) {
	result = fmt.Sprintf("%s, %s", message, name)
	index = "1"
	return
}
