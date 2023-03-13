package main

import (
	"context"
	"fmt"
	gpt3 "github.com/sashabaranov/go-openai"
)

func gptCompleteContext(ori string) (ret string) {
	// Get the context.
	ctx := context.Background()

	// For more details about the API of Open AI Chat Completion: https://platform.openai.com/docs/guides/chat
	req := gpt3.ChatCompletionRequest{
		// Model: The GPT-3.5 turbo model is the most powerful model available.
		Model: gpt3.GPT3Dot5Turbo,
		// The message to complete.
		Messages: []gpt3.ChatCompletionMessage{{
			Role:    gpt3.ChatMessageRoleUser,
			Content: ori,
		}},
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		ret = fmt.Sprintf("Err: %v", err)
	} else {
		// The response contains a list of choices, each with a score.
		// The score is a float between 0 and 1, with 1 being the most likely.
		// The choices are sorted by score, with the first choice being the most likely.
		// So we just take the first choice.
		ret = resp.Choices[0].Message.Content
	}

	return ret
}
