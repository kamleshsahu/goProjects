package main

import (
	"fmt"
	"time"
)

func main() {
	task := []Task{
		&EmailProcessor{
			email:   "kamlesh.sahu@dotpe.in",
			message: "Greetings!",
		},
		&MeetingScheduler{
			email: "tikesh@preplaced.in",
		},
		&EmailProcessor{
			email:   "vaibhav.singh@dotpe.in",
			message: "Hi!",
		},
		&MeetingScheduler{
			email: "learning@preplaced.in",
		},
		&MeetingScheduler{
			email: "staging@preplaced.in",
		},
		&EmailProcessor{
			email:   "tikesh.sahu@dotpe.in",
			message: "Hi Gamer!",
		},
		&EmailProcessor{
			email:   "test@dotpe.in",
			message: "Test message",
		},
	}

	worker := WorkerPool{
		Tasks:      task,
		workerSize: 2,
	}

	fmt.Println("Job Started!")
	worker.Start()
	fmt.Println("Job Finished!")

	newTask := &EmailProcessor{
		email:   "kamlesh.sahu@dotpe.in",
		message: "Interrupted task 123",
	}

	worker.wg.Add(1)
	time.Sleep(0 * time.Second)
	worker.TaskChannel <- newTask

	close(worker.TaskChannel)

	worker.wg.Wait()
}
