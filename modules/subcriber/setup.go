package subcriber

import (
	"context"
	"managerstudent/component"
)

func Setup(ctx component.AppContext) {
	DeleteStudentInClassAfterDeleteClass(ctx, context.Background())
	DeleteStudentInCourseAfterDeleteCourse(ctx, context.Background())
	SendNotifyAfterUserRegister(ctx, context.Background())
	SendNotifyAfterAddStudentToCourse(ctx, context.Background())
	SendNotifyAfterAddStudentToClass(ctx, context.Background())
	ChangeAcpUserAfterChangeNotify(ctx, context.Background())
}
