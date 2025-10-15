package orm

type SelectString struct {
	*Select[emptyModel, *emptyModel]
}

func SELECT2(cols ...string) SelectString {
	return SelectString{
		&Select[emptyModel, *emptyModel]{
			dbCommon: dbCommon{},
			cols:     cols,
		},
	}
}