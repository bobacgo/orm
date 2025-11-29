package orm

import "database/sql"

// 基本类型结果为空时，转换为go的零值

// createNullableRes 为 res 中的每个字段创建对应的 sql.Null* 类型
func createNullableRes(res []any) []any {
	nullableRes := make([]any, len(res))

	for i, r := range res {
		switch r.(type) {
		case *string:
			nullableRes[i] = &sql.NullString{}
		case *int:
			nullableRes[i] = &sql.Null[int]{}
		case *int8:
			nullableRes[i] = &sql.Null[int8]{}
		case *int16:
			nullableRes[i] = &sql.Null[int16]{}
		case *int32:
			nullableRes[i] = &sql.Null[int32]{}
		case *int64:
			nullableRes[i] = &sql.NullInt64{}
		case *uint:
			nullableRes[i] = &sql.Null[uint]{}
		case *uint8:
			nullableRes[i] = &sql.NullByte{}
		case *uint16:
			nullableRes[i] = &sql.Null[uint16]{}
		case *uint32:
			nullableRes[i] = &sql.Null[uint32]{}
		case *uint64:
			nullableRes[i] = &sql.Null[uint64]{}
		case *float32:
			nullableRes[i] = &sql.Null[float32]{}
		case *float64:
			nullableRes[i] = &sql.NullFloat64{}
		case *bool:
			nullableRes[i] = &sql.NullBool{}
		default:
			// 其他类型直接使用原始指针
			nullableRes[i] = r
		}
	}

	return nullableRes
}

// unwrapNullableRes 将 sql.Null* 类型的值转换回原始类型
func unwrapNullableRes(nullableRes []any, originalRes []any) {
	for i, nullable := range nullableRes {
		original := originalRes[i]

		switch n := nullable.(type) {
		case *sql.NullString:
			if n.Valid {
				if ptr, ok := original.(*string); ok {
					*ptr = n.String
				}
			}
		case *sql.Null[int]:
			if n.Valid {
				if ptr, ok := original.(*int); ok {
					*ptr = n.V
				}
			}
		case *sql.Null[int8]:
			if n.Valid {
				if ptr, ok := original.(*int8); ok {
					*ptr = n.V
				}
			}
		case *sql.Null[int16]:
			if n.Valid {
				if ptr, ok := original.(*int16); ok {
					*ptr = n.V
				}
			}
		case *sql.Null[int32]:
			if n.Valid {
				if ptr, ok := original.(*int32); ok {
					*ptr = n.V
				}
			}
		case *sql.NullInt64:
			if n.Valid {
				if ptr, ok := original.(*int64); ok {
					*ptr = n.Int64
				}
			}
		case *sql.Null[uint]:
			if n.Valid {
				if ptr, ok := original.(*uint); ok {
					*ptr = n.V
				}
			}
		case *sql.NullByte:
			if n.Valid {
				if ptr, ok := original.(*uint8); ok {
					*ptr = n.Byte
				}
			}
		case *sql.Null[uint16]:
			if n.Valid {
				if ptr, ok := original.(*uint16); ok {
					*ptr = n.V
				}
			}
		case *sql.Null[uint32]:
			if n.Valid {
				if ptr, ok := original.(*uint32); ok {
					*ptr = n.V
				}
			}
		case *sql.Null[uint64]:
			if n.Valid {
				if ptr, ok := original.(*uint64); ok {
					*ptr = n.V
				}
			}
		case *sql.Null[float32]:
			if n.Valid {
				if ptr, ok := original.(*float32); ok {
					*ptr = n.V
				}
			}
		case *sql.NullFloat64:
			if n.Valid {
				if ptr, ok := original.(*float64); ok {
					*ptr = n.Float64
				}
			}
		case *sql.NullBool:
			if n.Valid {
				if ptr, ok := original.(*bool); ok {
					*ptr = n.Bool
				}
			}
		}
	}
}
