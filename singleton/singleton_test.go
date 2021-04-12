package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	s := GetInstance()
	fmt.Println(s)
}
