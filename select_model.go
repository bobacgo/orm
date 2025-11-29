package orm

import (
	"context"
)

type SelectModel struct {
	*selec[SelectModel]
}

func SELECT1(row Model) *SelectModel {
	s := &SelectModel{
		selec: &selec[SelectModel]{
			dbCommon: dbCommon{},
		},
	}
	for _, v := range row.Mapping() {
		s.cols = append(s.cols, v.Column)
		s.res = append(s.res, v.Result)
	}
	s.setT(s)
	return s
}

func (d *SelectModel) Query(ctx context.Context, db Execer) error {
	if d.err != nil {
		return d.err
	}
	sqlText := d.SQL()
	d.debugPrint(ctx, sqlText)

	// 为每个字段创建对应的 sql.Null* 类型用于扫描
	nullableRes := createNullableRes(d.res)

	if err := db.QueryRowContext(ctx, sqlText, d.args...).Scan(nullableRes...); err != nil {
		return err
	}

	// 将扫描结果从 sql.Null* 类型转换回原始类型
	unwrapNullableRes(nullableRes, d.res)
	return nil
}
