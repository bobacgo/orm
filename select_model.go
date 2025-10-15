package orm

type SelectModel struct {
	*Select[emptyModel, *emptyModel]
}

func SELECT(row Model) SelectModel {
	slect := SelectModel{
		&Select[emptyModel, *emptyModel]{
			dbCommon: dbCommon{}, // TODO
		},
	}
	for _, v := range row.Mapping() {
		slect.cols = append(slect.cols, v.Column)
		slect.res = append(slect.res, v.Result)
	}
	return slect
}