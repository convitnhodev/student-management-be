package _const

type Role int

const (
	Student Role = iota
	Teacher
	AssistantPrincipal
	Headmaster
)

func (r Role) String() string {
	return []string{"Student", "Teacher", "AssistantPrincipal", "Headmaster"}[r]
}
