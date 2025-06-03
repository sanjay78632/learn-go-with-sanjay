package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
}

const fileName = "tasks.json"

func loadTasks() []Task {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{}
	}
	var tasks []Task
	json.Unmarshal(file, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatalf("Failed to save: %v", err)
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatalf("Cannot write file: %v", err)
	}
}

func main() {
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", 0, "Mark task as done by ID")
	flag.Parse()

	tasks := loadTasks()

	switch {
	case *add != "":
		newTask := Task{ID: len(tasks) + 1, Title: *add, Done: false}
		tasks = append(tasks, newTask)
		saveTasks(tasks)
		fmt.Println("✅ Task added.")
	case *list:
		for _, task := range tasks {
			status := "❌"
			if task.Done {
				status = "✅"
			}
			fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
		}
	case *done != 0:
		id := *done
		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				found = true
				break
			}
		}
		if found {
			saveTasks(tasks)
			fmt.Println("🎉 Task marked as done.")
		} else {
			fmt.Println("❌ Task not found.")
		}
	default:
		fmt.Println("⚠️ No valid flag provided. Use --add, --list or --done.")
	}
}