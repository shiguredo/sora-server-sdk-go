package sora

import (
	"encoding/json"
	"time"
)

const (
	EVENT_TYPE_CONNECTION_CREATED   = "connection.created"
	EVENT_TYPE_CONNECTION_UPDATED   = "connection.updated"
	EVENT_TYPE_CONNECTION_DESTROYED = "connection.destroyed"
	EVENT_TYPE_CONNECTION_FAILED    = "connection.failed"

	EVENT_TYPE_RECORDING_STARTED = "recording.started"
	EVENT_TYPE_RECORDING_REPORT  = "recording.report"

	EVENT_TYPE_ARCHIVE_STARTED         = "archive.started"
	EVENT_TYPE_ARCHIVE_AVAILABLE       = "archive.available"
	EVENT_TYPE_SPLIT_ARCHIVE_AVAILABLE = "split-archive.available"
	EVENT_TYPE_SPLIT_ARCHIVE_END       = "split-archive.end"

	EVENT_TYPE_SPOTLIGHT_FOCUSED   = "spotlight.focused"
	EVENT_TYPE_SPOTLIGHT_UNFOCUSED = "spotlight.unfocused"
)

type EventWebhookRequest struct {
	Timestamp time.Time `json:"timestamp"`

	ID string `json:"id"`

	Type string `json:"type"`

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

	EventMetadata json.RawMessage `json:"event_metadata"`

	Data json.RawMessage `json:"data"`
}

type ConnectionCreatedData struct {
	ChannelConnections int16 `json:"channel_connections"`
	Minutes            uint  `json:"minutes"`
}

type ConnectionUpdatedData struct {
	ChannelConnections int16 `json:"channel_connections"`
	Minutes            uint  `json:"minutes"`
}

type ConnectionDestroyedData struct {
	ChannelConnections int16  `json:"channel_connections"`
	TurnTransportType  string `json:"turn_transport_type"`
	Minutes            int64  `json:"minutes"`
	TotalReceivedBytes int64  `json:"total_received_bytes"`
	TotalSentBytes     int64  `json:"total_sent_bytes"`
}

type ArchiveAvailableData struct {
	ID           string `json:"id"`
	ConnectionID string `json:"connection_id"`
	Filename     string `json:"filename"`
	FilePath     string `json:"file_path"`
}
