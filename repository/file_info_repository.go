package repository

import (
	"time"
	"xorm.io/xorm"
)

type FileInfo struct {
	Id         int64     `xorm:"bigint 'id' notnull default('') comment('')" json:"id"`
	FilePath   string    `xorm:"varchar(100) 'file_path' notnull default('') comment('文件存储路径')" json:"filePath"` // 文件存储路径
	UserId     int64     `xorm:"bigint 'user_id' notnull default(0) comment('导入人编号')" json:"userId"`             // 导入人编号
	UserName   string    `xorm:"varchar(50) 'user_name' notnull default('') comment('导入人姓名')" json:"userName"`   // 导入人姓名
	Detail     string    `xorm:"varchar(200) 'detail' notnull default('') comment('导入说明')" json:"detail"`        // 导入说明
	CreateTime time.Time `xorm:"datetime 'create_time' notnull  comment('创建时间')" json:"createTime"`              // 创建时间
	UpdateTime time.Time `xorm:"datetime 'update_time' notnull  comment('更新时间')" json:"updateTime"`              // 更新时间
}

func CreateFileInfo(session *xorm.Session, entity *FileInfo) error {
	_, err := session.Insert(entity)

	return err
}
