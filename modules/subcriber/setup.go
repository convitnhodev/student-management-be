package subcriber

import (
	"context"
	"managerstudent/component"
)

func Setup(ctx component.AppContext) {
	DeleteStudentInClassAfterDeleteClass(ctx, context.Background())
	DeleteStudentInCourseAfterDeleteCourse(ctx, context.Background())
}
