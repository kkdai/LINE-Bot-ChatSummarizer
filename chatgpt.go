package main

import (
	"context"
	"fmt"

	gpt3 "github.com/sashabaranov/go-gpt3"
)

func CompleteContext(ori string) (ret string) {
	ctx := context.Background()
	req := gpt3.CompletionRequest{
		Model:     gpt3.GPT3TextDavinci003,
		MaxTokens: 300,
		Prompt:    ori,
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		ret = fmt.Sprintf("Err: %v", err)

	} else {
		ret = resp.Choices[0].Text
	}
	return ret
}
