package entity 

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date"`
	Completed   bool   `json:"completed"`
	Username    string `json:"username"`
}


