package pipeline

import "fmt"

type ImageJob struct {
	Input        string
	Output       string
	Format       string
	ResizeWidth  int
	Quality      int
	AttemptsLeft int
}

func ProcessImage(job ImageJob) error {
	fmt.Println("processing image btw.......")
	return nil
}
