package inject

import (
	"fmt"
	"go-Dependency-Injection/app/api/controllers"
	"go-Dependency-Injection/app/dal/db"
	"go-Dependency-Injection/app/dal/user"
    userDomain "go-Dependency-Injection/app/domain/user"
)


var GContainer = &Container{
	singletons: make(map[string]interface{}),
	transients: make(map[string]factory),
}

func Init() {
	//db
	GContainer.SetSingleton("MockDBRead", &db.MockDB{Host: "192.168.1.12:3036", User: "root", Pwd: "123456", Alias: "Read"})
	GContainer.SetSingleton("MockDBWrite", &db.MockDB{Host: "192.168.1.25:3036", User: "root", Pwd: "123456", Alias: "Write"})

	//仓储
	GContainer.SetSingleton("UserRepository", &user.UserRepository{})

	//服务
	GContainer.SetSingleton("UserService", &userDomain.UserService{})

	//控制器
	GContainer.SetSingleton("UserController", &controllers.UserController{})

	//控制器工厂
	ctlFactory := &CtrlFactory{}
	GContainer.SetSingleton("CtrlFactory", ctlFactory)

	GContainer.Entry(ctlFactory) //注入
	fmt.Println("--------------------------")
	fmt.Println(GContainer.String())
}
