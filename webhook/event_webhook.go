package webhook

import (
	"encoding/json"
	"time"
)

type EventRequest struct {
	Type string `json:"type"`

	ID string `json:"id"`

	Label    string `json:"label"`
	NodeName string `json:"node_name"`
	Version  string `json:"version"`

	Timestamp time.Time `json:"timestamp"`

	Multistream bool `json:"multistream"`
	Simulcast   bool `json:"simulcast"`
	Spotlight   bool `json:"spotlight"`

	Role         string `json:"role"`
	ChannelID    string `json:"channel_id"`
	SessionID    string `json:"session_id"`
	ClientID     string `json:"client_id"`
	BundleID     string `json:"bundle_id"`
	ConnectionID string `json:"connection_id"`

	Data          json.RawMessage `json:"data"`
	EventMetadata json.RawMessage `json:"event_metadata"`

	// undoc params
	Bandwidth Bandwidth `json:"bandwidth"`
}

// undoc params
type Bandwidth struct {
	ReceivedBytes int64 `json:"received_bytes"`
	SentBytes     int64 `json:"sent_bytes"`
}

type ConnectionData struct {
	CreatedTime      time.Time `json:"created_time"`
	CreatedTimestamp time.Time `json:"created_timestamp"`

	TotalReceivedBytes int64 `json:"total_received_bytes"`
	TotalSentBytes     int64 `json:"total_sent_bytes"`

	Audio AudioParams `json:"audio"`
	Video VideoParams `json:"video"`

	Minutes           int64  `json:"minutes"`
	TurnTransportType string `json:"turn_transport_type"`

	ChannelConnections         int64 `json:"channel_connections"`
	ChannelSenonlyConnections  int64 `json:"channel_sendonly_connections"`
	ChannelRecvonlyConnections int64 `json:"channel_recvonly_connections"`
	ChannelSendrecvConnections int64 `json:"channel_sendrecv_connections"`
}

type ConnectionCreatedData struct {
	ConnectionData
}

type ConnectionUpdatedData struct {
	ConnectionData
}

type ConnectionDestroyedData struct {
	ConnectionData

	DestroyedTimestamp   time.Time `json:"destroyed_timestamp"`
	TypeDisconnectReason string    `json:"type_disconnect_reason"`
}

type ArchiveAvailableData struct {
	ID           string `json:"id"`
	ConnectionID string `json:"connection_id"`
	Filename     string `json:"filename"`
	FilePath     string `json:"file_path"`
}
