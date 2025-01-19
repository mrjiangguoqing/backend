package main

import (
    "net/http"
    "time"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "github.com/gin-contrib/cors"
)

var jwtKey = []byte("your_secret_key")

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func main() {
    r := gin.Default()

    // 登录路由

    // API路由组，需要认证

    r.Use(cors.New(cors.Config{
        // 明确指定允许的源，不能使用通配符 *
        AllowOrigins:     []string{"http://152.42.242.240:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
    r.POST("/login", login)
	r.POST("/api/v1/register", Register)
	r.POST("/api/v1/login", Login)
    r.GET("/api/v1/user", GetUserInfo)

    

    api := r.Group("/api")
    api.Use(authMiddleware())
    {
        api.GET("/profile", getProfile)
    }

    r.Run(":8080")
}

func login(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // 这里应该查询数据库验证用户名和密码
    // 这里使用简单的示例验证
    if user.Username != "test" || user.Password != "test" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // 创建JWT Token
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: user.Username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": tokenString,
    })
}

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        bearerToken := strings.Split(authHeader, " ")
        if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
            c.Abort()
            return
        }

        tokenStr := bearerToken[1]
        claims := &Claims{}

        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("username", claims.Username)
        c.Next()
    }
}

func getProfile(c *gin.Context) {
    username, _ := c.Get("username")
    c.JSON(http.StatusOK, gin.H{
        "message": "Protected profile route",
        "user":    username,
        "email":    "testmail",
    })
}

