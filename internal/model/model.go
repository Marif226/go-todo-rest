package model

type User struct {
	Id 			int		`json:"id"`
	Name		string	`json:"name"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type TodoList struct {
	Id			int		`json:"id"`
	Title		string	`json:"title"`
	Description	string	`json:"description"`
}

type UsersList struct {
	Id			int
	UserId		int		
	ListId		int		
}

type TodoItem struct {
	Id			int		`json:"id"`
	Title		string	`json:"title"`
	Description	string	`json:"description"`
	Done		bool	`json:"done"`
}

type ListsItem struct {
	Id 		int
	ListId	int
	ItemId	int
}