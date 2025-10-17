package orm

import "strings"

type SelectString struct {
	*selec[SelectString]
}

func SELECT2(cols ...string) *SelectString {
	s := &SelectString{
		&selec[SelectString]{
			dbCommon: dbCommon{},
		}}
	s.setT(s)

	s.sql = "SELECT " + strings.Join(cols, ", ")
	return s
}