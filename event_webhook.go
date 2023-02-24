package sora

import (
	"encoding/json"
	"time"
)

type EventWebhookType string

const (
	EventWebhookTypeConnectionCreated   EventWebhookType = "connection.destroyed"
	EventWebhookTypeConnectionUpdated   EventWebhookType = "connection.updated"
	EventWebhookTypeConnectionDestroyed EventWebhookType = "connection.destroyed"
	EventWebhookTypeConnectionFailed    EventWebhookType = "connection.failed"

	EventWebhookTypeRecordingStarted EventWebhookType = "recording.started"
	EventWebhookTypeRecordingReport  EventWebhookType = "recording.report"

	EventWebhookTypeArchiveStarted        EventWebhookType = "archive.started"
	EventWebhookTypeArchiveAvailable      EventWebhookType = "archive.available"
	EventWebhookTypeSplitArchiveAvailable EventWebhookType = "split-archive.available"
	EventWebhookTypeSplitArchiveEnd       EventWebhookType = "split-archive.end"

	EventWebhookTypeSpotlightFocused   EventWebhookType = "spotlight.focused"
	EventWebhookTypeSpotlightUnfocused EventWebhookType = "spotlight.unfocused"
)

type EventWebhookRequest struct {
	Timestamp time.Time `json:"timestamp"`

	ID string `json:"id"`

	Type EventWebhookType `json:"type"`

	Label    string `json:"label"`
	NodeName string `json:"node_name"`
	Version  string `json:"version"`

	Multistream bool `json:"multistream"`
	Simulcast   bool `json:"simulcast"`
	Spotlight   bool `json:"spotlight"`

	Role         string `json:"role"`
	ChannelID    string `json:"channel_id"`
	SessionID    string `json:"session_id"`
	ClientID     string `json:"client_id"`
	BundleID     string `json:"bundle_id"`
	ConnectionID string `json:"connection_id"`

	EventMetadata *json.RawMessage `json:"event_metadata"`

	Data json.RawMessage `json:"data"`
}

type ConnectionCreatedData struct {
	CreatedTime      int64     `json:"created_time"`
	CreatedTimestamp time.Time `json:"created_timestamp"`

	Audio          bool   `json:"audio"`
	AudioCodecType string `json:"audio_codec_type"`
	AudioBitRate   *int32 `json:"audio_bit_rate"`

	Video          bool   `json:"video"`
	VideoCodecType string `json:"video_codec_type"`
	VideoBitRate   int32  `json:"video_bit_rate"`

	ChannelConnections         int32 `json:"channel_connections"`
	ChannelSendrecvConnections int32 `json:"channel_sendrecv_connections"`
	ChannelSendonlyConnections int32 `json:"channel_sendonly_connections"`
	ChannelRecvonlyConnections int32 `json:"channel_recvonly_connections"`

	Minutes int64 `json:"minutes"`

	TotalSentBytes     int64 `json:"total_sent_bytes"`
	TotalReceivedBytes int64 `json:"total_received_bytes"`

	TurnTransportType string `json:"turn_transport_type"`
}

type ConnectionUpdatedData struct {
	ConnectionCreatedData
}

type ConnectionDestroyedData struct {
	ConnectionCreatedData

	DestroyedTime      int64     `json:"destroyed_time"`
	DestroyedTimestamp time.Time `json:"destroyed_timestamp"`

	TypeDisconnectReason *string `json:"type_disconnect_reason"`
	DisconnectAPIReason  *string `json:"disconnect_api_reason"`
	Reason               *string `json:"reason"`
}

type ConnectionFailedData struct {
	Message string `json:"message"`

	ChannelConnections         int32 `json:"channel_connections"`
	ChannelSendrecvConnections int32 `json:"channel_sendrecv_connections"`
	ChannelSendonlyConnections int32 `json:"channel_sendonly_connections"`
	ChannelRecvonlyConnections int32 `json:"channel_recvonly_connections"`

	TotalSentBytes     int64 `json:"total_sent_bytes"`
	TotalReceivedBytes int64 `json:"total_received_bytes"`
}
