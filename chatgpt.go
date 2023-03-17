package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/sashabaranov/go-openai"
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

// Create image by DALL-E 2
func gptImageCreate(prompt string) (string, error) {
	ctx := context.Background()

	// Sample image by link
	reqUrl := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize512x512,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	respUrl, err := client.CreateImage(ctx, reqUrl)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return "", errors.New("Image creation error")
	}
	fmt.Println(respUrl.Data[0].URL)

	return respUrl.Data[0].URL, nil
}
