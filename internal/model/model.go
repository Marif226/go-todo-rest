package model

type User struct {
	Id 			int		`json:"-" db:"id"`
	Name		string	`json:"name" binding:"required"`
	Username	string	`json:"username" binding:"required"`
	Password	string	`json:"password" binding:"required"`
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