package inject

import (
	"go-Dependency-Injection/app/api/controllers"
)

type CtrlFactory struct {
	UserCtrl *controllers.UserController `inject:"UserController"`
}