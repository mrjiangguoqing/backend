func GetModelList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	search := c.Query("search")
	category := c.Query("category")
	popularity := c.Query("popularity")
	price := c.Query("price")

	// 调用服务层获取模型列表
	models, total, err := service.GetModelList(page, limit, search, category, popularity, price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取模型列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  models,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func GetModelDetail(c *gin.Context) {
	modelID := c.Param("id")

	// 调用服务层获取模型详情
	modelDetail, err := service.GetModelDetail(modelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取模型详情失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": modelDetail})
}

func CallModel(c *gin.Context) {
	modelID := c.Param("id")
	var inputData map[string]interface{}
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "输入数据格式错误"})
		return
	}

	// 调用服务层执行模型调用
	output, cost, err := service.CallModel(modelID, inputData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "模型调用失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"output": output,
		"cost":   cost,
	})
}

func GetModelRecords(c *gin.Context) {
	modelID := c.Param("id")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// 调用服务层获取使用记录
	records, total, err := service.GetModelRecords(modelID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取使用记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  records,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func FavoriteModel(c *gin.Context) {
	modelID := c.Param("id")
	userID := c.MustGet("userID").(int)

	// 调用服务层收藏模型
	err := service.FavoriteModel(userID, modelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "收藏模型失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "收藏成功"})
}

func UnfavoriteModel(c *gin.Context) {
	modelID := c.Param("id")
	userID := c.MustGet("userID").(int)

	// 调用服务层取消收藏
	err := service.UnfavoriteModel(userID, modelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "取消收藏失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "取消收藏成功"})
}

func SubmitModelReview(c *gin.Context) {
	modelID := c.Param("id")
	userID := c.MustGet("userID").(int)

	var reviewData struct {
		Rating  int    `json:"rating"`
		Comment string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&reviewData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "输入数据格式错误"})
		return
	}

	// 调用服务层提交评价
	err := service.SubmitModelReview(userID, modelID, reviewData.Rating, reviewData.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "提交评价失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "评价提交成功"})
}



func GetModelReviews(c *gin.Context) {
	modelID := c.Param("id")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// 调用服务层获取评价列表
	reviews, total, err := service.GetModelReviews(modelID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评价列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  reviews,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func GetUserFavorites(c *gin.Context) {
	userID := c.MustGet("userID").(int)
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// 调用服务层获取收藏列表
	favorites, total, err := service.GetUserFavorites(userID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取收藏列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  favorites,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func GetUserModelRecords(c *gin.Context) {
	userID := c.MustGet("userID").(int)
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// 调用服务层获取调用记录
	records, total, err := service.GetUserModelRecords(userID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取调用记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  records,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 模型相关路由
	r.GET("/api/models", GetModelList)
	r.GET("/api/models/:id", GetModelDetail)
	r.POST("/api/models/:id/call", CallModel)
	r.GET("/api/models/:id/records", GetModelRecords)
	r.POST("/api/models/:id/favorite", FavoriteModel)
	r.DELETE("/api/models/:id/favorite", UnfavoriteModel)
	r.POST("/api/models/:id/review", SubmitModelReview)
	r.GET("/api/models/:id/reviews", GetModelReviews)

	// 用户相关路由
	r.GET("/api/user/favorites", GetUserFavorites)
	r.GET("/api/user/records", GetUserModelRecords)

	return r
}