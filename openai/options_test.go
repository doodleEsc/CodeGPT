package openai

import (
	"testing"

	openai "github.com/sashabaranov/go-openai"
)

func Test_config_valid(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *config
		wantErr error
	}{
		{
			name: "valid config",
			cfg: newConfig(
				WithToken("test"),
				WithModel(openai.GPT3Dot5Turbo),
				WithProvider(OPENAI.String()),
			),
			wantErr: nil,
		},
		{
			name:    "missing token",
			cfg:     newConfig(),
			wantErr: errorsMissingToken,
		},
		{
			name: "missing model",
			cfg: newConfig(
				WithToken("test"),
				WithModel(""),
				WithProvider(OPENAI.String()),
			),
			wantErr: errorsMissingModel,
		},
		{
			name: "missing Azure deployment model",
			cfg: newConfig(
				WithToken("test"),
				WithModel(""),
				WithProvider(AZURE.String()),
			),
			wantErr: errorsMissingModel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.cfg
			if err := cfg.valid(); err != tt.wantErr {
				t.Errorf("config.valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
