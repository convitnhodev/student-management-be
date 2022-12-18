package studentBiz

type AddStudentToClassStore interface {
}

type addStudentToClassBiz struct {
	store AddStudentToClassStore
}

func NewAddStudentToClassBiz(store AddStudentToClassStore) *addStudentToClassBiz {
	return &addStudentToClassBiz{store: store}
}

