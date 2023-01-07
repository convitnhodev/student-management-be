package _const

type Role int

const (
	Teacher Role = iota
	Manager
)

func (r Role) String() string {
	return []string{"Teacher", "Manager"}[r]
}
