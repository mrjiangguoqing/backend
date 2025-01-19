//1. 获取系统通知
func GetSystemNotifications(c *gin.Context) {
    userID := c.MustGet("userID").(int)

    notifications, err := GetNotificationsByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取系统通知失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}


//2. 标记通知为已读
func MarkNotificationAsRead(c *gin.Context) {
    notificationID := c.Param("id")
    userID := c.MustGet("userID").(int)

    err := MarkNotificationReadByID(userID, notificationID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "标记通知为已读失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "通知已标记为已读"})
}

//3. 删除通知
func DeleteNotification(c *gin.Context) {
    notificationID := c.Param("id")
    userID := c.MustGet("userID").(int)

    err := DeleteNotificationByID(userID, notificationID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "删除通知失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "通知已删除"})
}

//4. 获取消息中心历史消息
func GetMessageHistory(c *gin.Context) {
    userID := c.MustGet("userID").(int)

    messages, err := GetMessagesByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取历史消息失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"messages": messages})
}


//5. 标记消息为已读
func MarkMessageAsRead(c *gin.Context) {
    messageID := c.Param("id")
    userID := c.MustGet("userID").(int)

    err := MarkMessageReadByID(userID, messageID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "标记消息为已读失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "消息已标记为已读"})
}

//6. 删除消息
func DeleteMessage(c *gin.Context) {
    messageID := c.Param("id")
    userID := c.MustGet("userID").(int)

    err := DeleteMessageByID(userID, messageID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "删除消息失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "消息已删除"})
}

//7. 发送邮件/短信通知
func SendEmailOrSMSNotification(c *gin.Context) {
    var notification NotificationRequest
    if err := c.ShouldBindJSON(&notification); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
        return
    }

    userID := c.MustGet("userID").(int)
    err := SendNotification(userID, notification.Type, notification.Content)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "发送通知失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "通知已发送"})
}


r := gin.Default()

// 系统通知相关路由
r.GET("/api/v1/notifications", GetSystemNotifications)
r.PUT("/api/v1/notifications/:id/read", MarkNotificationAsRead)
r.DELETE("/api/v1/notifications/:id", DeleteNotification)

// 消息中心相关路由
r.GET("/api/v1/messages", GetMessageHistory)
r.PUT("/api/v1/messages/:id/read", MarkMessageAsRead)
r.DELETE("/api/v1/messages/:id", DeleteMessage)

// 邮件/短信通知相关路由
r.POST("/api/v1/notifications/send", SendEmailOrSMSNotification)