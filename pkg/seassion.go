package pkg

import (
	"github.com/sashabaranov/go-openai"
	"sync"
)

type SessionConfig struct {
	ID         int64
	MaxHistory int
	MaxTokens  int
}

type Session struct {
	config  SessionConfig
	history []*openai.ChatCompletionMessage
	lock    sync.Mutex
}

func NewSession(id int64) *Session {
	return &Session{
		config: SessionConfig{
			ID:         id,
			MaxHistory: 8,
			MaxTokens:  2048,
		},
		history: make([]*openai.ChatCompletionMessage, 0),
	}
}

func (s *Session) trimHistory() {
	if len(s.history) > s.config.MaxHistory {
		s.history = s.history[len(s.history)-s.config.MaxHistory:]
	}
	tokens := 0
	for i := len(s.history) - 1; i >= 0; i-- {
		tokens += len(s.history[i].Content)
		if tokens > s.config.MaxTokens {
			s.history = s.history[i+1:]
			break
		}
	}
}

func (s *Session) Chat(with func([]*openai.ChatCompletionMessage) (*openai.ChatCompletionMessage, error)) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.trimHistory()
	answer, err := with(s.history)
	if err != nil {
		s.history = append(s.history, answer)
	}
}
