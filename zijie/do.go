package main

import "fmt"

// BehaviorStatus 表示节点执行的返回状态
type BehaviorStatus int

const (
	// Success 表示节点成功完成了它的任务
	Success BehaviorStatus = iota
	// Failure 表示节点未能完成它的任务
	Failure
	// Running 表示节点正在执行它的任务
	Running
)

// Behavior 表示行为树中的节点接口
type Behavior interface {
	Tick() BehaviorStatus
}

// Action 是一个行为节点，执行具体的任务
type Action struct {
	name string
}

// Tick 执行具体的任务，并返回执行状态
func (a *Action) Tick() BehaviorStatus {
	fmt.Printf("动作: %s 正在执行\n", a.name)
	// 假设这个动作总是成功
	return Success
}

// Sequence 是一个控制节点，按顺序执行子节点，直到一个子节点返回 Failure 或者所有子节点返回 Success
type Sequence struct {
	children []Behavior
}

// Tick 按顺序执行子节点
func (s *Sequence) Tick() BehaviorStatus {
	for _, child := range s.children {
		status := child.Tick()
		if status != Success {
			return status
		}
	}
	return Success
}

// Selector 是一个控制节点，按顺序执行子节点，直到一个子节点返回 Success
type Selector struct {
	children []Behavior
}

// Tick 按顺序执行子节点
func (s *Selector) Tick() BehaviorStatus {
	for _, child := range s.children {
		status := child.Tick()
		if status != Failure {
			return status
		}
	}
	return Failure
}

func main() {
	// 创建一个行为树
	root := &Selector{
		children: []Behavior{
			&Sequence{
				children: []Behavior{
					&Action{name: "检查敌人"},
					&Action{name: "射击"},
				},
			},
			&Action{name: "巡逻"},
		},
	}

	// 执行行为树
	status := root.Tick()
	fmt.Printf("行为树状态: %v\n", status)
}
