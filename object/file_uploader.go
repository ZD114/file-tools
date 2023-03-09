package object

import (
	"sync"
)

type Uploader struct {
	ClientFileMetadata                // 文件元数据
	SliceSeq                          // 需要重传的序号
	waitGoroutine      sync.WaitGroup // 同步goroutine
	NewLoader          bool           // 是否是新创建的上传器
	FilePath           string         // 上传文件路径
	SliceBytes         int            // 切片大小
	RetryChannel       chan *FilePart // 重传channel通道
	MaxGtChannel       chan struct{}  // 限制上传的goroutine的数量通道
	StartTime          int64          // 上传开始时间
}
