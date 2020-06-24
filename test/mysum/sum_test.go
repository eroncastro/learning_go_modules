package mysum

import (
	"fmt"
	"testing"
	"proj1/mysum"
)

func TestAbc(t *testing.T) {
	if (4 == mysum.Sum(2, 2)) {
		fmt.Println("Success")
	} else {
		t.Error() // to indicate test failed
	}
}
