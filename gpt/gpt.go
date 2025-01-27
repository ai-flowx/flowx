package gpt

import (
	"context"

	"github.com/pkg/errors"
)

const (
	DefaultChatStream           = false
	DefaultChatMaxTokens        = 4096
	DefaultChatFrequencyPenalty = 0
	DefaultChatPresencePenalty  = 0
	DefaultChatTemperature      = 1
	DefaultChatTopP             = 0.7
)

const (
	providerDoubaoChat = "doubao-chat"
)

type Gpt interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Chat(context.Context, *ChatRequest) (ChatResponse, error)
}

type Config struct {
	Provider string
	Api      string
	Key      string
	Endpoint string
}

type ChatRequest struct {
	Model            string        `json:"model"`
	Messages         []ChatMessage `json:"messages"`
	Stream           bool          `json:"stream"`
	MaxTokens        int64         `json:"max_tokens"`
	FrequencyPenalty float64       `json:"frequency_penalty"`
	PresencePenalty  float64       `json:"presence_penalty"`
	Temperature      float64       `json:"temperature"`
	TopP             float64       `json:"top_p"`
}

type ChatResponse struct {
	Id      string       `json:"id"`
	Choices []ChatChoice `json:"choices"`
}

type ChatChoice struct {
	Message ChatMessage `json:"message"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type gpt struct {
	cfg *Config
	gt  Gpt
}

func New(_ context.Context, cfg *Config) Gpt {
	var gt Gpt

	if cfg.Provider == providerDoubaoChat {
		gt = &DoubaoChat{
			Api:      cfg.Api,
			Key:      cfg.Key,
			Endpoint: cfg.Endpoint,
		}
	} else {
		// BYPASS
	}

	return &gpt{
		cfg: cfg,
		gt:  gt,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (g *gpt) Init(ctx context.Context) error {
	if err := g.gt.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init\n")
	}

	return nil
}

func (g *gpt) Deinit(ctx context.Context) error {
	if err := g.gt.Deinit(ctx); err != nil {
		return errors.Wrap(err, "failed to deinit\n")
	}

	return nil
}

func (g *gpt) Chat(ctx context.Context, request *ChatRequest) (ChatResponse, error) {
	buf, err := g.gt.Chat(ctx, request)
	if err != nil {
		return ChatResponse{}, errors.Wrap(err, "failed to run\n")
	}

	return buf, nil
}
