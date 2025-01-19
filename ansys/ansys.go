//用户活跃度统计

func GetUserActivity(c *gin.Context) {
    // 从上下文中获取用户ID
    userID := c.MustGet("userID").(int)

    // 获取用户活跃度数据
    activity, err := GetUserActivityByID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户活跃度失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"activity": activity})
}


//模型使用统计
func GetModelUsage(c *gin.Context) {
    // 从上下文中获取模型ID
    modelID := c.Param("modelID")

    // 获取模型使用数据
    usage, err := GetModelUsageByID(modelID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取模型使用数据失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"usage": usage})
}



//收入统计
func GetRevenue(c *gin.Context) {
    // 获取平台收入数据
    revenue, err := GetPlatformRevenue()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取收入数据失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"revenue": revenue})
}


//用户增长统计
func GetUserGrowth(c *gin.Context) {
    // 获取用户增长数据
    growth, err := GetUserGrowthData()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户增长数据失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"growth": growth})
}


func SetupRouter() *gin.Engine {
    r := gin.Default()

    // 用户活跃度统计
    r.GET("/api/v1/user/activity", GetUserActivity)

    // 模型使用统计
    r.GET("/api/v1/model/usage/:modelID", GetModelUsage)

    // 收入统计
    r.GET("/api/v1/revenue", GetRevenue)

    // 用户增长统计
    r.GET("/api/v1/user/growth", GetUserGrowth)

    return r
}
