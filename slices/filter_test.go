package slices_test

import (
	"reflect"
	"testing"

	"github.com/tobiashort/utils-go/slices"
)

func TestFilter(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	even := slices.Filter(numbers, func(number int) bool { return number%2 == 0 })
	if !reflect.DeepEqual(even, []int{0, 2, 4, 6, 8, 10, 12}) {
		t.Fatalf("%v\n", even)
	}
	t.Logf("%v\n", even)
}
