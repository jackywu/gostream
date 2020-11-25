package stream

import (
	"fmt"
	"testing"
)

func TestStream(t *testing.T) {
	s := NewStringStream([]string{"a", "b", "c"})
	result := s.Filter(func(a string) bool {
		return a != "b"
	}).Map(func(a string) string {
		return fmt.Sprintf("%s.2", a)
	}).Collect()
	fmt.Println(result)
}