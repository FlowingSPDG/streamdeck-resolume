package sdresolume

import "encoding/json"

// Config は単一のProperty Inspectorの設定を表します
type SelectClipConfig struct {
	Host   string      `json:"host"`   // ホストアドレス
	Port   string      `json:"port"`   // ポート番号
	Column json.Number `json:"column"` // カラム
}
