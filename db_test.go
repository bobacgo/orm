package orm_test

import (
	"testing"

	"github.com/bobacgo/orm"
)

// test insert

func TestInsert(t *testing.T) {
	orm.TestInsert(t)
}

func TestInsertModel(t *testing.T) {
	orm.TestInsertModel(t)
}

func TestInsertModels(t *testing.T) {
	orm.TestInsertModels(t)
}

func TestDelete(t *testing.T) {
	orm.TestDelete(t)
}

// test select

func TestSelectOne(t *testing.T) {
	orm.TestSelectOne(t)
}

func TestSelectMany(t *testing.T) {
	orm.TestSelectMany(t)
}

func TestSelectString(t *testing.T) {
	orm.TestSelectString(t)
}