package sora

import (
	"encoding/json"
	"time"
)

type SessionWebhookRequest struct {
	Timestamp time.Time `json:"timestamp"`

	ID string `json:"id"`

	Type string `json:"type"`

	Label    string `json:"label"`
	Version  string `json:"version"`
	NodeName string `json:"node_name"`
}

type SessionCreatedWebhookRequest struct {
	SessionWebhookRequest

	Multistream bool `json:"multistream"`
	Spotlight   bool `json:"spotlight"`

	ChannelID string `json:"channel_id"`
	SessionID string `json:"session_id"`

	CreatedTime      int64     `json:"created_at"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
}

type SessionConnection struct {
	Role         string `json:"role"`
	ClientID     string `json:"client_id"`
	ConnectionID string `json:"connection_id"`

	Simulcast bool `json:"simulcast"`

	Audio          bool   `json:"audio"`
	AudioCodecType string `json:"audio_codec_type"`
	AudioBitRate   *int32 `json:"audio_bit_rate"`

	Video          bool   `json:"video"`
	VideoCodecType string `json:"video_codec_type"`
	VideoBitRate   int32  `json:"video_bit_rate"`

	CreatedTime   int64 `json:"created_time"`
	DestroyedTime int64 `json:"destroyed_time"`
}

type SessionDestroyedWebhookRequest struct {
	SessionCreatedWebhookRequest

	DestroyedTime      int64     `json:"destroyed_time"`
	DestroyedTimestamp time.Time `json:"destroyed_timestamp"`

	MaxConnections   int32 `json:"max_connections"`
	TotalConnections int32 `json:"total_connections"`

	Connections []SessionConnection `json:"connections"`

	SessionMetadata json.RawMessage `json:"session_metadata"`
}

type SessionVanishedWebhookRequest struct {
	SessionWebhookRequest
}
