package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"regexp"
	"strings"
)

type TaskManager struct {
	TaskMap  map[int]*Task
	TaskTree *treemap.Map
}

func Constructor(tasks [][]int) TaskManager {
	taskManager := TaskManager{
		TaskMap:  make(map[int]*Task),
		TaskTree: treemap.NewWithStringComparator(),
	}

	for _, task := range tasks {
		tObject := &Task{Id: task[1], UserId: task[0], Priority: task[2]}
		taskManager.TaskMap[tObject.Id] = tObject
		key := generateKey(tObject.Priority, tObject.Id)
		taskManager.TaskTree.Put(key, tObject)
	}

	return taskManager
}

func generateKey(priority, taskId int) string {
	return fmt.Sprintf("%10d#%10d", priority, taskId)
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	tObject := &Task{Id: taskId, UserId: userId, Priority: priority}
	this.TaskMap[taskId] = tObject
	key := generateKey(priority, taskId)
	this.TaskTree.Put(key, tObject)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	tObject := this.TaskMap[taskId]
	this.Rmv(taskId)
	this.Add(tObject.UserId, taskId, newPriority)
}

func (this *TaskManager) Rmv(taskId int) {
	tObject := this.TaskMap[taskId]
	key := generateKey(tObject.Priority, tObject.Id)
	this.TaskTree.Remove(key)
	delete(this.TaskMap, tObject.Id)
}

func (this *TaskManager) ExecTop() int {
	if len(this.TaskMap) == 0 {
		return -1
	}
	_, highestObject := this.TaskTree.Max()
	if highestObject == nil {
		return -1
	}
	this.Rmv(highestObject.(*Task).Id)
	return highestObject.(*Task).UserId
}

type Task struct {
	Id       int
	UserId   int
	Priority int
}

func main() {
	taskList := [][]int{
		{1, 101, 10},
		{2, 102, 20},
		{3, 103, 15},
	}
	manager := Constructor(taskList)
	manager.Add(4, 104, 5)
	manager.Edit(102, 8)
	fmt.Println(manager.ExecTop()) // Output: 3
	manager.Rmv(101)
	manager.Add(5, 105, 15)
	fmt.Println(manager.ExecTop()) // Output: 5

}

func hasMatch(s string, p string) bool {
	parts := strings.Split(p, "*")
	p = parts[0] + ".*" + parts[1]

	re := regexp.MustCompile(p)
	return re.MatchString(s)
}

//20 102
//15 103
