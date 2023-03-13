package main

import (
	"context"
	"fmt"

	gpt3 "github.com/sashabaranov/go-openai"
)

func gptCompleteContext(ori string) (ret string) {
	// Get the context.
	ctx := context.Background()

	// 主要 API Open AI Completion https://beta.openai.com/docs/guides/completion
	req := gpt3.CompletionRequest{
		// Model: The GPT-3.5 turbo model is the most powerful model available.
		Model: gpt3.GPT3Dot5Turbo,
		// 最大輸出內容，可以調整一下。
		MaxTokens: 300,
		// 輸入文字，也就是你平時在 ChatGPT 詢問他的問題。
		Prompt: ori,
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		ret = fmt.Sprintf("Err: %v", err)

	} else {
		// 回來的成果中，拿精準度最高的為答案。
		ret = resp.Choices[0].Text
	}
	return ret
}
