package sora

import "encoding/json"

type EventWebhookRequest struct {
	Type          string          `json:"type"`
	Label         string          `json:"label"`
	NodeName      string          `json:"node_name"`
	Version       string          `json:"version"`
	Timestamp     string          `json:"timestamp"`
	Multistream   bool            `json:"multistream"`
	Simulcast     bool            `json:"simulcast"`
	Spotlight     bool            `json:"spotlight"`
	Role          string          `json:"role"`
	ChannelID     string          `json:"channel_id"`
	SessionID     string          `json:"session_id"`
	ClientID      string          `json:"client_id"`
	ConnectionID  string          `json:"connection_id"`
	Data          json.RawMessage `json:"data"`
	EventMetadata json.RawMessage `json:"event_metadata"`
}

type ConnectionUpdatedData struct {
	Minutes uint `json:"minutes"`
}

type ConnectionDestroyedData struct {
	ChannelConnections uint   `json:"channel_connections"`
	TurnTransportType  string `json:"turn_transport_type"`
	Minutes            uint   `json:"minutes"`
	TotalReceivedBytes uint   `json:"total_received_bytes"`
	TotalSentBytes     uint   `json:"total_sent_bytes"`
}

type ArchiveAvailableData struct {
	ID           string `json:"id"`
	ConnectionID string `json:"connection_id"`
	Filename     string `json:"filename"`
	FilePath     string `json:"file_path"`
}
