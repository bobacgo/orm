package orm

type SelectModels[T any, P PModel[T]] struct {
	*Select[T, P]
	row  *T
	rows *[]*T
}

func SELECT1[T any, P PModel[T]](rows *[]*T) SelectModels[T, P] {
	slect := SelectModels[T, P]{
		Select: &Select[T, P]{
			dbCommon: dbCommon{}, // TODO
		},
		row:  new(T),
		rows: rows,
	}
	mappings := P(slect.row).Mapping()
	for _, v := range mappings {
		slect.cols = append(slect.cols, v.Column)
		slect.res = append(slect.res, v.Result)
	}
	return slect
}