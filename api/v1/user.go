package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"xiaowei_list/model/database"
	"xiaowei_list/model/response"
	"xiaowei_list/service"
	"xiaowei_list/utils"
)

// Register
//
//	@Summary		注册
//	@Description	注册新用户
//	@Tags			登录模块
//	@Accept			json
//	@Produce		json
//	@Param			data	body		response.RegisterQ	true	"用户名，密码"
//	@Success		200		{object}	response.CommonA	"是否成功，返回信息，Token"
//	@Router			/api/v1/register [post]
func Register(c *gin.Context) {
	// 获取请求数据
	data := utils.BindJsonAndValid(c, &response.RegisterQ{}).(*response.RegisterQ)
	// 用户的用户名已经注册过的情况
	if _, notFound := service.GetUserByUsername(data.Username); !notFound {
		c.JSON(http.StatusOK, response.CommonA{Success: false, Message: "用户已存在"})
		return
	}
	// 将密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 12)
	if err != nil {
		panic("CreateUser: hash password error")
	}
	fmt.Println(data.Username)
	fmt.Println(string(hashedPassword))
	fmt.Println(data.Email)
	fmt.Println(data.Nickname)
	// 成功创建用户
	if err := service.CreateUser(&database.User{Username: data.Username, Password: string(hashedPassword), Nickname: data.Nickname, Email: data.Email, UserType: "user"}); err != nil {
		panic("CreateUser: create user error")
	}
	// 返回响应
	c.JSON(http.StatusOK, response.CommonA{Success: true, Message: "注册成功"})
}

// Login
//
//	@Summary		用户登录
//	@Description	根据用户邮箱和密码等生成token，并将token返回给用户
//	@Tags			登录模块
//	@Accept			json
//	@Produce		json
//	@Param			data	body		response.LoginQ	true	"用户名，密码"
//	@Success		200		{object}	response.LoginA	"是否成功，返回信息，Token"
//	@Router			/api/v1/login [post]
func Login(c *gin.Context) {
	// 获取请求中的数据
	data := utils.BindJsonAndValid(c, &response.LoginQ{}).(*response.LoginQ)
	// 用于登录的邮箱未注册的情况
	user, notFound := service.GetUserByUsername(data.Username)
	if notFound {
		c.JSON(http.StatusOK, response.LoginA{Success: false, Message: "登录失败，用户名不存在", Token: "", UserId: 0})
		return
	}
	// 密码错误的情况
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.JSON(http.StatusOK, response.LoginA{Success: false, Message: "登录失败，密码错误", Token: "", UserId: 0})
		return
	}
	// 成功返回响应
	token := utils.GenerateToken(uint64(user.UserId))
	c.JSON(http.StatusOK, response.LoginA{Success: true, Message: "登录成功", Token: token, UserId: user.UserId})
}

// GetUser
//
//	@Summary		获取用户资料
//	@Description	获取一个用户公开的详细资料
//	@Tags			用户模块
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string				false	"token"
//	@Param			id		path		int					true	"用户ID"
//	@Success		200		{object}	response.GetUserA	"是否成功，返回信息，用户名"
//	@Router			/api/v1/users/{id} [get]
func GetUser(c *gin.Context) {
	// 获取请求数据
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.GetUserA{Success: false, Message: "请求参数非法", Code: 400})
		return
	}
	// 查询对应ID的用户
	if user, notFound := service.GetUserByID(int(id)); notFound {
		c.JSON(http.StatusOK, response.GetUserA{Success: false, Message: "找不到对应的用户", Code: 404})
	} else {
		// 返回响应
		c.JSON(http.StatusOK, response.GetUserA{Success: true, Message: "查找成功", Name: user.Username, Code: 200})
	}
}
