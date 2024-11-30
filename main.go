package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	name        string
	isCompleted bool
}

// 常量定义任务命令
const (
	CmdAdd    = 1
	CmdRemove = 2
	CmdList   = 3
	CmdChange = 4
)

var tasks []Task
const MAX_TASKS = 10 // 最大任务数量

// 移除第一个命令行参数（文件名），获取实际传入的参数
func removeFirstArg(args []string) ([]string, error) {
	if len(args) > 1 {
		return args[1:], nil // 排除第一个参数
	}
	return nil, errors.New("至少需要一个命令参数")
}

// 验证命令是否有效
func validateCommand(args []string) (int, error) {
	if len(args) < 1 {
		return -1, errors.New("命令为空，请输入有效命令")
	}

	switch strings.ToLower(args[0]) {
	case "add":
		return CmdAdd, nil
	case "remove":
		return CmdRemove, nil
	case "list":
		return CmdList, nil
	case "change":
		return CmdChange, nil
	default:
		return -1, fmt.Errorf("无效命令: %s", args[0])
	}
}

// 向任务列表中添加任务
func addTask(args []string) {
	if len(args) < 2 {
		fmt.Println("请提供任务名称。")
		return
	}

	if len(tasks) >= MAX_TASKS {
		fmt.Println("任务列表已满，无法添加更多任务。")
		return
	}

	// 任务名称去除空格并检查是否为空
	taskName := strings.TrimSpace(args[1])
	if taskName == "" {
		fmt.Println("任务名称不能为空。")
		return
	}

	// 创建任务并添加
	task := Task{name: taskName, isCompleted: false}
	tasks = append(tasks, task)
	fmt.Printf("任务 \"%v\" 已添加到任务列表。\n", task.name)
}

// 列出所有任务
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("任务列表为空。")
		return
	}

	fmt.Println("列出所有任务：")
	for i, task := range tasks {
		status := "未完成"
		if task.isCompleted {
			status = "已完成"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.name, status)
	}
}

// 处理用户命令
func handleCommand(commandId int, args []string) {
	switch commandId {
	case CmdAdd:
		addTask(args)
	case CmdList:
		listTasks()
	default:
		fmt.Println("不支持的命令或未实现。")
	}
}

func main() {
	args, err := removeFirstArg(os.Args)
	if err != nil {
		fmt.Println("错误:", err)
		return
	}

	// 验证命令并处理
	commandId, err := validateCommand(args)
	if err != nil {
		fmt.Println("错误:", err)
		return
	}

	handleCommand(commandId, args)
}
