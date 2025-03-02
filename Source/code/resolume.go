package sdresolume

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/FlowingSPDG/resolume-go"
	"github.com/FlowingSPDG/streamdeck"
	"golang.org/x/xerrors"
)

// Plugin はStreamDeckプラグインの状態を管理します
type Plugin struct {
	client *streamdeck.Client
}

func NewPlugin(client *streamdeck.Client) *Plugin {
	return &Plugin{
		client: client,
	}
}

// SelectColumnKeyDown はボタンが押された際に呼ばれます
func (p *Plugin) SelectColumnKeyDown(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	p.client.LogMessage(ctx, fmt.Sprintf("event: %v", event))
	var payload streamdeck.KeyDownPayload[SelectColumnConfig]
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		p.client.LogMessage(ctx, fmt.Sprintf("failed to unmarshal action settings: %s", err))
		return xerrors.Errorf("failed to unmarshal action settings: %w", err)
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	p.client.LogMessage(ctx, fmt.Sprintf("%s Host: %s, Port: %s, Colunm: %s", now, payload.Settings.Host, payload.Settings.Port, payload.Settings.Column))

	// クライアントを作成
	rc, err := resolume.NewClient(payload.Settings.Host, payload.Settings.Port)
	if err != nil {
		p.client.LogMessage(ctx, fmt.Sprintf("failed to create client: %s", err))
		return xerrors.Errorf("failed to create client: %w", err)
	}

	column, err := payload.Settings.Column.Int64()
	if err != nil {
		p.client.LogMessage(ctx, fmt.Sprintf("failed to convert column to int64: %s", err))
		return xerrors.Errorf("failed to convert column to int64: %w", err)
	}

	if err := rc.SelectColumn(ctx, column); err != nil {
		p.client.LogMessage(ctx, fmt.Sprintf("failed to select clip: %s", err))
		return xerrors.Errorf("failed to select clip: %w", err)
	}

	return nil
}

func (p *Plugin) SelectClipKeyDown(ctx context.Context, client *streamdeck.Client, event streamdeck.Event) error {
	p.client.LogMessage(ctx, fmt.Sprintf("event: %v", event))
	var payload streamdeck.KeyDownPayload[SelectLayerClipConfig]
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		p.client.LogMessage(ctx, fmt.Sprintf("failed to unmarshal action settings: %s", err))
		return xerrors.Errorf("failed to unmarshal action settings: %w", err)
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	p.client.LogMessage(ctx, fmt.Sprintf("%s Host: %s, Port: %s, Layer: %s, Clip: %s", now, payload.Settings.Host, payload.Settings.Port, payload.Settings.Layer, payload.Settings.Clip))

	// クライアントを作成
	rc, err := resolume.NewClient(payload.Settings.Host, payload.Settings.Port)
	if err != nil {
		p.client.LogMessage(ctx, fmt.Sprintf("failed to create client: %s", err))
		return xerrors.Errorf("failed to create client: %w", err)
	}

	layer, err := payload.Settings.Layer.Int64()
	if err != nil {
		p.client.LogMessage(ctx, fmt.Sprintf("failed to convert layer to int64: %s", err))
		return xerrors.Errorf("failed to convert layer to int64: %w", err)
	}

	clip, err := payload.Settings.Clip.Int64()
	if err != nil {
		p.client.LogMessage(ctx, fmt.Sprintf("failed to convert column to int64: %s", err))
		return xerrors.Errorf("failed to convert column to int64: %w", err)
	}

	if err := rc.SelectLayerClip(ctx, int(layer), int(clip)); err != nil {
		p.client.LogMessage(ctx, fmt.Sprintf("failed to select clip: %s", err))
		return xerrors.Errorf("failed to select clip: %w", err)
	}

	return nil
}

func (p *Plugin) Run(ctx context.Context) error {
	selectColumnAction := p.client.Action("dev.flowingspdg.resolume.selectcolumn")
	selectColumnAction.RegisterHandler(streamdeck.KeyDown, p.SelectColumnKeyDown)

	selectClipAction := p.client.Action("dev.flowingspdg.resolume.selectclip")
	selectClipAction.RegisterHandler(streamdeck.KeyDown, p.SelectClipKeyDown)

	return p.client.Run(ctx)
}
