package webhook

import (
	"encoding/json"
	"time"
)

type SessionRequest struct {
	Timestamp time.Time `json:"timestamp"`

	Type     string `json:"type"`
	Label    string `json:"label"`
	Version  string `json:"version"`
	NodeName string `json:"node_name"`
}

type SessionCreatedRequest struct {
	SessionRequest

	Multistream bool `json:"multistream"`

	ChannelID   string `json:"channel_id"`
	SessionID   string `json:"session_id"`
	CreatedTime int64  `json:"created_time"`
}

type SessionConnection struct {
	CreatedTime   int64 `json:"created_time"`
	DestroyedTime int64 `json:"destroyed_time"`

	Role         string `json:"role"`
	ClientID     string `json:"client_id"`
	ConnectionID string `json:"connection_id"`
	Simulcast    bool   `json:"simulcast"`

	Audio interface{} `json:"audio"`
	Video interface{} `json:"video"`
}

type SessionDestroyedRequest struct {
	SessionRequest
	Multistream        bool                `json:"multistream"`
	Spotlight          bool                `json:"spotlight"`
	ChannelID          string              `json:"channel_id"`
	SessionID          string              `json:"session_id"`
	CreatedTime        int64               `json:"created_time"`
	CreatedTimestamp   time.Time           `json:"created_timestamp"`
	DestroyedTime      int64               `json:"destroyed_time"`
	DestroyedTimestamp time.Time           `json:"destroyed_timestamp"`
	MaxConnections     int                 `json:"max_connections"`
	TotalConnections   int                 `json:"total_connections"`
	Connections        []SessionConnection `json:"connections"`
	SessionMetadata    json.RawMessage     `json:"session_metadata"`
}

type SessionVanishedRequest struct {
	SessionRequest
}
