package orm

type TestModel struct {
	ID int
	A  string
	B  int
	C  string
}

//func (m *TestModel) TableName() string {
//	return "test_model"
//}

func (m *TestModel) Mapping() []*Mapping {
	return []*Mapping{
		{"id", &m.ID, m.ID},
		{"a", &m.A, m.A},
		{"b", &m.B, m.B},
		{"c", &m.C, m.C},
	}
}

type Mapping struct {
	Column string
	Result any
	Value  any
}