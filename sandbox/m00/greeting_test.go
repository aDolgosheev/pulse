// greeting_test.go — файл с суффиксом _test.go компилируется только для тестов

package m00

import "testing"

func TestGreeting(t *testing.T) {
	if got := Greeting("Aleksandr"); got != "Hello, Aleksandr!" {
		t.Errorf("Greeting() = %q", got)
	}
}