func Recharge(c *gin.Context) {
    // 从上下文中获取用户ID
    userID := c.MustGet("userID").(int)

    // 解析请求体中的充值金额和支付方式
    var rechargeRequest struct {
        Amount float64 `json:"amount"`
        Method string  `json:"method"` // 支付方式：支付宝、微信、银行卡等
    }
    if err := c.ShouldBindJSON(&rechargeRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
        return
    }

    // 调用充值逻辑
    err := RechargeUserBalance(userID, rechargeRequest.Amount, rechargeRequest.Method)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "充值失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "充值成功"})
}



func GetBalance(c *gin.Context) {
    // 从上下文中获取用户ID
    userID := c.MustGet("userID").(int)

    // 获取用户余额
    balance, err := GetUserBalance(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取余额失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"balance": balance})
}



func GetConsumptionRecords(c *gin.Context) {
    // 从上下文中获取用户ID
    userID := c.MustGet("userID").(int)

    // 获取消费记录
    records, err := GetUserConsumptionRecords(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取消费记录失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"records": records})
}

func PurchasePackage(c *gin.Context) {
    // 从上下文中获取用户ID
    userID := c.MustGet("userID").(int)

    // 解析请求体中的套餐ID
    var packageRequest struct {
        PackageID int `json:"package_id"`
    }
    if err := c.ShouldBindJSON(&packageRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
        return
    }

    // 调用套餐购买逻辑
    err := PurchaseUserPackage(userID, packageRequest.PackageID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "套餐购买失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "套餐购买成功"})
}


func CalculateCost(c *gin.Context) {
    // 从上下文中获取用户ID
    userID := c.MustGet("userID").(int)

    // 解析请求体中的使用情况
    var usageRequest struct {
        CallCount        int `json:"call_count"`
        ResourceUsage    int `json:"resource_usage"`
    }
    if err := c.ShouldBindJSON(&usageRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
        return
    }

    // 计算费用
    cost, err := CalculateUserCost(userID, usageRequest.CallCount, usageRequest.ResourceUsage)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "费用计算失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"cost": cost})
}

r := gin.Default()

// 充值功能
r.POST("/api/v1/recharge", Recharge)

// 余额查询
r.GET("/api/v1/balance", GetBalance)

// 消费记录
r.GET("/api/v1/consumption-records", GetConsumptionRecords)

// 套餐购买
r.POST("/api/v1/purchase-package", PurchasePackage)

// 费用计算
r.POST("/api/v1/calculate-cost", CalculateCost)