//将返回有用的数据都写入数据库，同时添加有用的数据。便于后续处理。
package models

import (
	"time"
)

type FullResponse struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"` // 唯一标识符
	Object            string    `gorm:"not null"`                 // response object
	Created           int64     `gorm:"not null"`                 // created timestamp
	Model             string    `gorm:"not null"`                 // model used
	Choices           string    `gorm:"not null"`                 // choices (JSON string)
	PromptTokens      int       `gorm:"not null"`                 // input tokens
	CompletionTokens  int       `gorm:"not null"`                 // output tokens
	TotalTokens       int       `gorm:"not null"`                 // total tokens used
	PromptCacheHitTokens int    `gorm:"not null"`                 // prompt cache hit tokens
	PromptCacheMissTokens int   `gorm:"not null"`                 // prompt cache miss tokens
	SystemFingerprint string    `gorm:"not null"`                 // system fingerprint
	CreatedAt         time.Time `gorm:"autoCreateTime"`           // created timestamp
}


package services

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"models"
)

type HTTPResponse struct {
	ID                string `json:"id"`
	Object            string `json:"object"`
	Created           int64  `json:"created"`
	Model             string `json:"model"`
	Choices           []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs       interface{} `json:"logprobs"`
		FinishReason   string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens      int `json:"prompt_tokens"`
		CompletionTokens  int `json:"completion_tokens"`
		TotalTokens       int `json:"total_tokens"`
		PromptCacheHitTokens int `json:"prompt_cache_hit_tokens"`
		PromptCacheMissTokens int `json:"prompt_cache_miss_tokens"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}

func WriteFullResponseToDB(db *gorm.DB, response string) error {
	// 解析 HTTP 响应
	var httpResponse HTTPResponse
	if err := json.Unmarshal([]byte(response), &httpResponse); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	// 创建 FullResponse 实例
	fullResponse := models.FullResponse{
		Object:              httpResponse.Object,
		Created:             httpResponse.Created,
		Model:               httpResponse.Model,
		Choices:             fmt.Sprintf("%+v", httpResponse.Choices), // 将 choices 转换为字符串，或者存储为 JSON 字符串
		PromptTokens:        httpResponse.Usage.PromptTokens,
		CompletionTokens:    httpResponse.Usage.CompletionTokens,
		TotalTokens:         httpResponse.Usage.TotalTokens,
		PromptCacheHitTokens: httpResponse.Usage.PromptCacheHitTokens,
		PromptCacheMissTokens: httpResponse.Usage.PromptCacheMissTokens,
		SystemFingerprint:   httpResponse.SystemFingerprint,
	}

	// 写入数据库
	if err := db.Create(&fullResponse).Error; err != nil {
		return fmt.Errorf("failed to insert full response into database: %w", err)
	}

	return nil
}




package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"services"
)

func main() {
	// 配置数据库连接
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移数据库
	db.AutoMigrate(&models.FullResponse{})

	// 示例 HTTP 响应
	response := `{
		"id": "8f91ea21-8570-492a-bb67-b7c0d4030ddc",
		"object": "chat.completion",
		"created": 1737283315,
		"model": "deepseek-chat",
		"choices": [{
			"index": 0,
			"message": {
				"role": "assistant",
				"content": "Hello! How can I assist you today? 😊"
			},
			"logprobs": null,
			"finish_reason": "stop"
		}],
		"usage": {
			"prompt_tokens": 11,
			"completion_tokens": 11,
			"total_tokens": 22,
			"prompt_cache_hit_tokens": 0,
			"prompt_cache_miss_tokens": 11
		},
		"system_fingerprint": "fingerprint_here"
	}`

	// 写入数据库
	err = services.WriteFullResponseToDB(db, response)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Full response data successfully written to the database.")
	}
}
