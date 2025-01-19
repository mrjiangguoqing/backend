package user_manage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查用户是否已存在
	// 这里假设有一个函数 CheckUserExists 来检查用户是否存在
	if CheckUserExists(req.Email, req.Phone) {
		c.JSON(http.StatusConflict, gin.H{"error": "用户已存在"})
		return
	}

	// 创建用户
	// 这里假设有一个函数 CreateUser 来创建用户
	if err := CreateUser(req.Email, req.Phone, req.Password, req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}



package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证用户登录信息
	// 这里假设有一个函数 AuthenticateUser 来验证用户登录信息
	user, err := AuthenticateUser(req.Email, req.Phone, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 生成token
	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}


func GetUserInfo(c *gin.Context) {
	// 从上下文中获取用户ID
	userID := c.MustGet("userID").(int)

	// 获取用户信息
	// 这里假设有一个函数 GetUserByID 来获取用户信息
	user, err := GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
		return
	}

	c.JSON(http.StatusOK, user)
}


func GetUserRole(c *gin.Context) {
	// 从上下文中获取用户ID
	userID := c.MustGet("userID").(int)

	// 获取用户角色
	// 这里假设有一个函数 GetUserRoleByID 来获取用户角色
	role, err := GetUserRoleByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户角色失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

func DeleteUser(c *gin.Context) {
	// 从上下文中获取用户ID
	userID := c.MustGet("userID").(int)

	// 删除用户
	// 这里假设有一个函数 DeleteUserByID 来删除用户
	if err := DeleteUserByID(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注销用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注销成功"})
}




type ForgotPasswordRequest struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func ForgotPassword(c *gin.Context) {
	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查用户是否存在
	// 这里假设有一个函数 CheckUserExists 来检查用户是否存在
	if !CheckUserExists(req.Email, req.Phone) {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	// 发送重置密码链接或验证码
	// 这里假设有一个函数 SendResetPasswordLink 来发送重置密码链接或验证码
	if err := SendResetPasswordLink(req.Email, req.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "发送重置密码链接失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "重置密码链接已发送"})
}

