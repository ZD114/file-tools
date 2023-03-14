package object

type FileSearchParam struct {
	PageParam
	FileName string `json:"fileName"`
	UserName string `json:"userName"`
	Detail   string `json:"detail"`
}
