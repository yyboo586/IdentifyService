package timeTask

import (
	"sync"

	"IdentifyService/internal/app/system/model"
	"IdentifyService/internal/app/system/service"
)

func init() {
	service.RegisterTaskList(New())
}

func New() service.ITaskList {
	return &sTaskList{
		mu: new(sync.RWMutex),
	}
}

type sTaskList struct {
	taskList []*model.TimeTask
	mu       *sync.RWMutex
}

// AddTask 添加任务
func (s *sTaskList) AddTask(task *model.TimeTask) {
	if task.FuncName == "" || task.Run == nil {
		return
	}
	s.taskList = append(s.taskList, task)
}

// GetByName 通过方法名获取对应task信息
func (s *sTaskList) GetByName(funcName string) *model.TimeTask {
	var result *model.TimeTask
	for _, item := range s.taskList {
		if item.FuncName == funcName {
			result = item
			break
		}
	}
	return result
}

// EditParams 修改参数
func (s *sTaskList) EditParams(funcName string, params []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, item := range s.taskList {
		if item.FuncName == funcName {
			item.Param = params
			break
		}
	}
}
