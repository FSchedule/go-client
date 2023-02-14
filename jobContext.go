package fSchedule

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/stopwatch"
	"time"
)

// JobContext 任务在运行时，更新状态
type JobContext struct {
	id           int64                                  // 主键
	name         string                                 // 实现Job的特性名称（客户端识别哪个实现类）
	Data         collections.Dictionary[string, string] // 数据
	nextTimespan int64                                  // 下次执行时间
	progress     int                                    // 当前进度
	status       TaskStatus                             // 执行状态
	sw           *stopwatch.Watch                       // 运行时间
}

// SetNextAt 设置下次运行时间
func (receiver *JobContext) SetNextAt(t time.Time) {
	receiver.nextTimespan = t.UnixMicro()
}

// SetProgress 设置任务进度
func (receiver *JobContext) SetProgress(progress int) {
	receiver.progress = progress
}