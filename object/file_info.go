package object

import (
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
