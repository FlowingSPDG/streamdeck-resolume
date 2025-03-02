package sdresolume

import (
	"context"
	"encoding/json"

	"github.com/FlowingSPDG/resolume-go"
	"github.com/FlowingSPDG/streamdeck"
	"golang.org/x/xerrors"
)

// Plugin はStreamDeckプラグインの状態を管理します
type Plugin struct {
	client *streamdeck.Client
	config *Config
}

func NewPlugin(client *streamdeck.Client, config *Config) *Plugin {
	return &Plugin{
		client: client,
		config: config,
	}
}

// PluginAction はボタンのアクション設定を表します
type PluginAction struct {
	Host    string `json:"host"`    // IPアドレス
	Port    string `json:"port"`    // ポート番号
	Command string `json:"command"` // 送信するコマンド
}

// KeyDown はボタンが押された際に呼ばれます
func (p *Plugin) KeyDown(ctx context.Context, event streamdeck.Event) error {
	var actionSettings PluginAction
	if err := json.Unmarshal(event.Payload, &actionSettings); err != nil {
		return xerrors.Errorf("failed to unmarshal action settings: %w", err)
	}

	// クライアントを作成
	client, err := resolume.NewClient(actionSettings.Host, actionSettings.Port)
	if err != nil {
		return xerrors.Errorf("failed to create client: %w", err)
	}

	if err := client.SelectClipByID(1); err != nil {
		return xerrors.Errorf("failed to select clip: %w", err)
	}

	return nil
}

func (p *Plugin) Run(ctx context.Context) error {
	return p.client.Run(ctx)
}
