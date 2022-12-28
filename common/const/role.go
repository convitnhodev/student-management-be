package _const

type Role int

const (
	Student Role = iota
	Teacher
	Headmaster
)

const CurrentUser = "User"

func (r Role) String() string {
	return []string{"Student", "Teacher", "AssistantPrincipal", "Headmaster"}[r]
}

type Requester interface {
}
