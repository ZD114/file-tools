package object

import (
	"fmt"
	"strings"
	"time"
	"zhangda/file-tools/repository"
	"zhangda/file-tools/util"
)

func CreateFileInfo(m FileInfoProperty) (*repository.FileInfo, error) {
	entity := new(repository.FileInfo)

	util.Copy(m, entity, "Id", "CreateTime", "UpdateTime")

	now := time.Now()

	entity.CreateTime = now
	entity.UpdateTime = now

	session := GetFileInfoDB().NewSession()

	defer session.Close()

	if err := repository.CreateFileInfo(session, entity); err != nil {
		session.Rollback()
		return nil, err
	}

	return entity, nil
}

func GetFileItems(searchParam FileSearchParam) (PageResult, error) {
	countSb := new(strings.Builder)
	where, params := buildSearchAreasWhere(searchParam)

	countSb.WriteString("SELECT COUNT(id) FROM file_info ")
	countSb.WriteString(where)

	total := int64(0)

	if _, err := GetFileInfoDB().SQL(countSb.String(), params...).Get(&total); err != nil {
		return *NewPageResult(), err
	}

	pageResult := &PageResult{Total: total}

	pageResult.TotalPages = pageResult.GetTotalPages(searchParam.PageSize)

	if total > 0 {
		querySb := new(strings.Builder)
		queryParams := append(params, searchParam.Start(), searchParam.PageSize)

		querySb.WriteString("SELECT * FROM file_info")
		querySb.WriteString(where)
		querySb.WriteString(" order by id desc LIMIT ?,?")

		entities := make([]FileInfoProperty, 0)

		if err := GetFileInfoDB().SQL(querySb.String(), queryParams...).Find(&entities); err != nil {
			return *pageResult, err
		}

		pageResult.Data = make([]interface{}, len(entities))

		for i, entity := range entities {
			pageResult.Data[i] = entity
		}

		return *pageResult, nil
	}

	return *pageResult, nil
}

func buildSearchAreasWhere(searchParam FileSearchParam) (string, []interface{}) {
	sb := new(strings.Builder)
	params := make([]interface{}, 0, searchParam.ParamLength())

	sb.WriteString(" WHERE 1=1")

	if len(searchParam.FileName) > 0 {
		sb.WriteString(" AND file_path like ? ")
		fileName := fmt.Sprintf("%%%s%%", searchParam.FileName)

		params = append(params, fileName)
	}

	if len(searchParam.UserName) > 0 {
		sb.WriteString(" AND user_name like ? ")
		userName := fmt.Sprintf("%%%s%%", searchParam.UserName)

		params = append(params, userName)
	}

	if len(searchParam.Detail) > 0 {
		sb.WriteString(" AND detail like ? ")
		detail := fmt.Sprintf("%%%s%%", searchParam.Detail)

		params = append(params, detail)
	}

	return sb.String(), params
}
