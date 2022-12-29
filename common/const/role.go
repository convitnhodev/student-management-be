package _const

type Role int

const (
	Teacher Role = iota
	Manager
)

const CurrentUser = "User"

func (r Role) String() string {
	return []string{"Teacher", "Manager"}[r]
}

type Requester interface {
}
