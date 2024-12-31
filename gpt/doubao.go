package gpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	chatMessage = `
		[
		  {
		    "role": "system",
		    "content": "You are a helpful assistant."
		  },
		  {
		    "role": "user",
		    "content": "%s"
		  }
		]`

	visionMessage = `
		[
		  {
		    "role": "user",
		    "content": [
		      {
		        "type": "text",
		        "text": "提取图片文字，严格保留段落和内容格式"
		      },
		      {
		        "type": "image_url",
		        "image_url": {
		          "detail": "high",
		          "url": "data:image/png;base64,%s"
		        }
		      }
		    ]
		  }
		]`
)

type DoubaoChat struct {
	Api   string
	Model string
	Key   string
}

type ChatPrompt struct {
	Messages []ChatMessages `json:"messages"`
}

type ChatMessages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Id      string       `json:"id"`
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Model   string       `json:"model"`
	Choices []ChatChoice `json:"choices"`
	Usage   ChatUsage    `json:"usage"`
}

type ChatChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	LogProbs     bool        `json:"log_probs"`
	FinishReason string      `json:"finish_reason"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatUsage struct {
	PromptTokens        int                     `json:"prompt_tokens"`
	CompletionTokens    int                     `json:"completion_tokens"`
	TotalTokens         int                     `json:"total_tokens"`
	PromptTokensDetails ChatPromptTokensDetails `json:"prompt_tokens_details"`
}

type ChatPromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

type DoubaoVision struct {
	Api   string
	Model string
	Key   string
}

type VisionPrompt struct {
	Messages []VisionMessages `json:"messages"`
}

type VisionMessages struct {
	Role    string          `json:"role"`
	Content []VisionContent `json:"content"`
}

type VisionContent struct {
	Type     string         `json:"type"`
	Text     string         `json:"text"`
	ImageUrl VisionImageUrl `json:"image_url"`
}

type VisionImageUrl struct {
	Url string `json:"url"`
}

type VisionResponse struct {
	Id      string         `json:"id"`
	Object  string         `json:"object"`
	Created int64          `json:"created"`
	Model   string         `json:"model"`
	Choices []VisionChoice `json:"choices"`
	Usage   VisionUsage    `json:"usage"`
}

type VisionChoice struct {
	Index        int           `json:"index"`
	Message      VisionMessage `json:"message"`
	LogProbs     bool          `json:"log_probs"`
	FinishReason string        `json:"finish_reason"`
}

type VisionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type VisionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func (c *DoubaoChat) Init(_ context.Context) error {
	return nil
}

func (c *DoubaoChat) Deinit(_ context.Context) error {
	return nil
}

func (c *DoubaoChat) Run(_ context.Context, content string) ([]string, error) {
	var res ChatResponse

	buf := fmt.Sprintf(`{
		"model": "%s",
		"messages": %s
	}`, c.Model, fmt.Sprintf(chatMessage, content))

	req, err := http.NewRequest("POST", c.Api, bytes.NewBuffer([]byte(buf)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal((body), &res); err != nil {
		return nil, err
	}

	return c.parseContent(&res), nil
}

func (v *DoubaoChat) parseContent(res *ChatResponse) []string {
	var buf []string

	for _, item := range res.Choices {
		buf = append(buf, item.Message.Content)
	}

	return buf
}

func (v *DoubaoVision) Init(_ context.Context) error {
	return nil
}

func (v *DoubaoVision) Deinit(_ context.Context) error {
	return nil
}

func (v *DoubaoVision) Run(_ context.Context, content string) ([]string, error) {
	var res VisionResponse

	buf := fmt.Sprintf(`{
		"model": "%s",
		"messages": %s
	}`, v.Model, fmt.Sprintf(visionMessage, content))

	req, err := http.NewRequest("POST", v.Api, bytes.NewBuffer([]byte(buf)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+v.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal((body), &res); err != nil {
		return nil, err
	}

	return v.parseContent(&res), nil
}

func (v *DoubaoVision) parseContent(res *VisionResponse) []string {
	var buf []string

	for _, item := range res.Choices {
		buf = append(buf, item.Message.Content)
	}

	return buf
}
