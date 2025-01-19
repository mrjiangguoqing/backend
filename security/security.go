func GetUserRole(c *gin.Context) {
    // 从上下文中获取用户ID
    userID := c.MustGet("userID").(int)

    // 获取用户角色
    role, err := GetUserRoleByID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户角色失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"role": role})
}


func CheckAPIKey(c *gin.Context) {
    apiKey := c.GetHeader("API-Key")
    if apiKey == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "API Key 缺失"})
        c.Abort()
        return
    }

    // 验证API Key
    isValid, err := ValidateAPIKey(apiKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "验证API Key失败"})
        c.Abort()
        return
    }

    if !isValid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的API Key"})
        c.Abort()
        return
    }

    c.Next()
}

func RateLimit(c *gin.Context) {
    userID := c.MustGet("userID").(int)
    if IsRateLimited(userID) {
        c.JSON(http.StatusTooManyRequests, gin.H{"error": "请求过于频繁，请稍后再试"})
        c.Abort()
        return
    }

    c.Next()
}


func EncryptData(c *gin.Context) {
    var data map[string]interface{}
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
        return
    }

    encryptedData, err := Encrypt(data)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "数据加密失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"encrypted_data": encryptedData})
}


func LogUserAction(c *gin.Context) {
    userID := c.MustGet("userID").(int)
    action := c.Request.Method + " " + c.Request.URL.Path

    err := LogAction(userID, action)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "日志记录失败"})
        return
    }

    c.Next()
}


func DetectAnomalies(c *gin.Context) {
    userID := c.MustGet("userID").(int)
    action := c.Request.Method + " " + c.Request.URL.Path

    if IsAnomalous(userID, action) {
        c.JSON(http.StatusForbidden, gin.H{"error": "检测到异常操作，已阻止"})
        c.Abort()
        return
    }

    c.Next()
}

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // 添加中间件
    r.Use(CheckAPIKey)
    r.Use(RateLimit)
    r.Use(LogUserAction)
    r.Use(DetectAnomalies)

    // 添加路由
    r.POST("/api/v1/login", Login)
    r.GET("/api/v1/user/role", GetUserRole)
    r.POST("/api/v1/encrypt", EncryptData)

    return r
}

