package workflow

type Workflow struct {
	ID     string
	Name   string
	Status string
	Tasks  []Task
}
