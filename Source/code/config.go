package sdresolume

import (
	"encoding/json"
	"fmt"
	"os"
)

// HostConfig は単一のResolume接続設定を表します
type HostConfig struct {
	Name     string `json:"name"`     // ホストの表示名
	Host     string `json:"host"`     // ホストアドレス
	Port     int    `json:"port"`     // ポート番号
	Username string `json:"username"` // ユーザー名（オプション）
	Password string `json:"password"` // パスワード（オプション）
}

// Config はプラグインの設定を表します
type Config struct {
	Hosts []HostConfig `json:"hosts"` // 複数のホスト設定
}

// LoadConfig は設定ファイルから設定を読み込みます
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// SaveConfig は設定をファイルに保存します
func (c *Config) SaveConfig(path string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
