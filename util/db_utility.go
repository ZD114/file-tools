package util

import (
	"database/sql"
	"reflect"
	"strconv"
	"unsafe"
)

func SnakeCasedName(name string) string {
	newstr := make([]byte, 0, len(name)+1)

	for i := 0; i < len(name); i++ {
		c := name[i]

		if isUpper := 'A' <= c && c <= 'Z'; isUpper {
			if i > 0 {
				newstr = append(newstr, '_')
			}

			c += 'a' - 'A'
		}

		newstr = append(newstr, c)
	}

	return BytesToString(newstr)
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// GetColumnsFromStruct 获取更新列
func GetColumnsFromStruct(source interface{}, ignoreFields ...string) []string {
	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source)

	// 要复制哪些字段
	m := SliceToMap(ignoreFields)
	filed := make([]string, 0)

	for i := 0; i < sourceValue.NumField(); i++ {
		name := SnakeCasedName(sourceType.Field(i).Name)

		if _, ok := m[name]; !ok {
			filed = append(filed, SnakeCasedName(name))
		}
	}

	return filed
}

func GetColumnsFromColumns(columns []string, ignoreFields ...string) []string {
	if len(ignoreFields) == 0 {
		return columns
	}

	ignoreFieldMap := make(map[string]string)

	for _, field := range ignoreFields {
		ignoreFieldMap[field] = field
	}

	cols := make([]string, 0, len(columns)-len(ignoreFieldMap))

	for _, col := range columns {
		if _, ok := ignoreFieldMap[col]; !ok {
			cols = append(cols, col)
		}
	}

	return cols
}

func GetQueryColumns(rows *sql.Rows) ([]string, map[string]string, error) {
	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		return nil, nil, err
	}

	length := len(columnTypes)

	columns := make([]string, length)
	columnTypeMap := make(map[string]string, length)

	for i, ct := range columnTypes {
		columns[i] = ct.Name()
		columnTypeMap[ct.Name()] = ct.DatabaseTypeName()
	}

	return columns, columnTypeMap, nil
}

func QueryForInterface(db *sql.DB, sqlInfo string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(sqlInfo, args...)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	columns, columnTypeMap, err := GetQueryColumns(rows)

	if err != nil {
		return nil, err
	}

	columnLength := len(columns)
	cache := make([]interface{}, columnLength) //临时存储每行数据

	for i := range cache { //为每一列初始化一个指针
		var a interface{}

		cache[i] = &a
	}

	var list []map[string]interface{} //返回的切片

	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})

		for i, data := range cache {
			if ct, ok := columnTypeMap[columns[i]]; ok {
				if (ct == "VARCHAR" || ct == "DATETIME" || ct == "TEXT") && *data.(*interface{}) != nil {
					v := string((*data.(*interface{})).([]byte))

					if ct == "DATETIME" && v == "0000-00-00 00:00:00" {
						item[columns[i]] = ""
					} else {
						item[columns[i]] = v
					}
				} else if ct == "DECIMAL" && *data.(*interface{}) != nil {
					if f, err := strconv.ParseFloat(string((*data.(*interface{})).([]byte)), 64); err == nil {
						item[columns[i]] = f
					} else {
						item[columns[i]] = *data.(*interface{})
					}
				} else {
					item[columns[i]] = *data.(*interface{})
				}
			} else {
				item[columns[i]] = *data.(*interface{}) //取实际类型
			}

		}

		list = append(list, item)
	}

	return list, nil
}
