// greeting.go

package m00

func Greeting(name string) string {
	if name == "" {
		name = "Unknown"
	}
	return "Hello, " + name + "!"
}