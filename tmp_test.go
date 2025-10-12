package orm

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	// xxx
	// xxx
	rows := make([]*TestModel, 0)
	SL(&rows)
}

func SL[T any, P PModel[T]](rows *[]*T) {
	newElem := new(T)
	mappings := P(newElem).Mapping()
	for _, v := range mappings {
		fmt.Println(v)
	}
}