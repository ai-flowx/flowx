package gpt

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type DoubaoChat struct {
	Api      string
	Key      string
	Endpoint string
}

type DoubaoChatResponse struct {
	Id      string             `json:"id"`
	Model   string             `json:"model"`
	Created int64              `json:"created"`
	Object  string             `json:"object"`
	Choices []DoubaoChatChoice `json:"choices"`
	Usage   DoubaoChatUsage    `json:"usage"`
}

type DoubaoChatChoice struct {
	Index        int                `json:"index"`
	FinishReason string             `json:"finish_reason"`
	Message      DoubaoChatMessage  `json:"message"`
	Logprobs     DoubaoChatLogprobs `json:"logprobs"`
}

type DoubaoChatMessage struct {
	Role      string               `json:"role"`
	Content   string               `json:"content"`
	ToolCalls []DoubaoChatToolCall `json:"tool_calls"`
}

type DoubaoChatUsage struct {
	PromptTokens        int                           `json:"prompt_tokens"`
	CompletionTokens    int                           `json:"completion_tokens"`
	TotalTokens         int                           `json:"total_tokens"`
	PromptTokensDetails DoubaoChatPromptTokensDetails `json:"prompt_tokens_details"`
}

type DoubaoChatPromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

type DoubaoChatLogprobs struct {
	Content []DoubaoChatTokenLogprob `json:"content"`
}

type DoubaoChatTokenLogprob struct {
	Token       string                 `json:"token"`
	Bytes       []int                  `json:"bytes"`
	Logprob     float64                `json:"logprob"`
	TopLogprobs []DoubaoChatTopLogprob `json:"top_logprobs"`
}

type DoubaoChatTopLogprob struct {
	Token   string  `json:"token"`
	Bytes   []int   `json:"bytes"`
	Logprob float64 `json:"logprob"`
}

type DoubaoChatToolCall struct {
	Id       string             `json:"id"`
	Type     string             `json:"type"`
	Function DoubaoChatFunction `json:"function"`
}

type DoubaoChatFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

func (c *DoubaoChat) Init(_ context.Context) error {
	return nil
}

func (c *DoubaoChat) Deinit(_ context.Context) error {
	return nil
}

func (c *DoubaoChat) Chat(_ context.Context, request *ChatRequest) (ChatResponse, error) {
	var buf bytes.Buffer
	var res DoubaoChatResponse

	request.Model = c.Endpoint

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return ChatResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, c.Api, &buf)
	if err != nil {
		return ChatResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ChatResponse{}, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ChatResponse{}, err
	}

	if err := json.Unmarshal((body), &res); err != nil {
		return ChatResponse{}, err
	}

	return c.parseContent(&res), nil
}

func (c *DoubaoChat) parseContent(res *DoubaoChatResponse) ChatResponse {
	var buf ChatResponse

	buf.Id = res.Id

	for _, item := range res.Choices {
		var b ChatChoice
		b.Message.Role = item.Message.Role
		b.Message.Content = item.Message.Content
		buf.Choices = append(buf.Choices, b)
	}

	return buf
}
