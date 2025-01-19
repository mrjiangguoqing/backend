//模型交互是不是可以直接通过前端进行而不需要后端，这样可能只需要复制一份数据。

//sql记录下用户每次调用模型所产生的费用就行。
package main

import (
	"encoding/json"
	"fmt"
)

// ModelResponse 定义模型返回的数据结构
type ModelResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

// CalculateCostFromJSON 从 JSON 数据中解析 usage 并计算费用
func CalculateCostFromJSON(jsonData string, pricePerThousandTokens float64, promptPricePerThousandTokens float64, completionPricePerThousandTokens float64) {
	// 解析 JSON 数据
	var response ModelResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return
	}

	// 计算费用（按总 token 计算）
	totalCost := (float64(response.Usage.TotalTokens) / 1000) * pricePerThousandTokens
	fmt.Printf("总费用（按总 token 计算）: $%.6f\n", totalCost)

	// 计算费用（区分输入和输出 token 计算）
	promptCost := (float64(response.Usage.PromptTokens) / 1000) * promptPricePerThousandTokens
	completionCost := (float64(response.Usage.CompletionTokens) / 1000) * completionPricePerThousandTokens
	detailedCost := promptCost + completionCost
	fmt.Printf("总费用（区分输入和输出 token 计算）: $%.6f\n", detailedCost)
}

func main() {
	// 获取模型返回数据，考虑作为一个中间件处理，计算费用，同时写入数据库。
	jsonData := response.json[]

	// 定义单价
	pricePerThousandTokens := 0.002              // 每 1000 个 token 的价格
	promptPricePerThousandTokens := 0.0015       // 输入 token 单价
	completionPricePerThousandTokens := 0.002    // 输出 token 单价

	// 调用函数计算费用
	CalculateCostFromJSON(jsonData, pricePerThousandTokens, promptPricePerThousandTokens, completionPricePerThousandTokens)
}
