package _const

type Role int

const (
	Student Role = iota
	Teacher
	AssistantPrincipals
	Headmaster
)

func (r Role) String() string {
	return []string{"Student", "Teacher", "AssistantPrincipals", "Headmaster"}[r]
}
