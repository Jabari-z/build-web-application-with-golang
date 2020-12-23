package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "go-gin-web/handler"
	"go-gin-web/model"
	"go-gin-web/pkg/errno"
)


//endpoint
//     服务

func AddUser(c *gin.Context)  {
	var r model.User
	// 绑定数据
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.User{
		Username: r.Username,
		Password: r.Password,
	}
	// 核验数据
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// 插入数据库
	if _,err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	// 处理
	SendResponse(c, nil, u)
}

// SelectUser 查询用户
func SelectUser(c *gin.Context)  {
	//user_name 字段的内容
	name := c.Query("user_name")
	if name == ""{
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	var  user model.User
	if err := user.SelectUserByName(name);nil != err {
		fmt.Println(err)
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	// 核验数据
	if err := user.Validate(); err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
