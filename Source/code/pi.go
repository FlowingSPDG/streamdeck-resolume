package sdresolume

import "encoding/json"

// SelectColumnConfig はSelectColumnのProperty Inspectorの設定を表します
type SelectColumnConfig struct {
	Host   string      `json:"host"`   // ホストアドレス
	Port   string      `json:"port"`   // ポート番号
	Column json.Number `json:"column"` // カラム
}

type SelectLayerClipConfig struct {
	Host  string      `json:"host"`  // ホストアドレス
	Port  string      `json:"port"`  // ポート番号
	Layer json.Number `json:"layer"` // レイヤー
	Clip  json.Number `json:"clip"`  // クリップ
}
