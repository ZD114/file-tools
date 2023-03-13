package object

import "time"

type FileInfoProperty struct {
	Id         int64     `json:"id"`
	FilePath   string    `json:"filePath"`
	UserId     int64     `json:"userId"`
	UserName   string    `json:"userName"`
	Detail     string    `json:"detail"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
