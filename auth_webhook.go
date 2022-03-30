package sora

import "encoding/json"

type AuthWebhookRequest struct {
	Version  string `json:"version"`
	Label    string `json:"label"`
	NodeName string `json:"node_name"`

	Multistream bool `json:"multistream"`
	Simulcast   bool `json:"simulcast"`
	Spotlight   bool `json:"spotlight"`

	Role         string `json:"role"`
	ChannelID    string `json:"channel_id"`
	ClientID     string `json:"client_id"`
	ConnectionID string `json:"connection_id"`

	Metadata json.RawMessage `json:"metadata"`

	ChannelConnections         uint `json:"channel_connetions"`
	ChannelSendrecvConnections uint `json:"channel_sendrecv_connections"`
	ChannelSendonlyConnections uint `json:"channel_sendonly_connections"`
	ChannelRecvonlyConnections uint `json:"channel_recvonly_connections"`

	Audio interface{} `json:"audio"`
	Video interface{} `json:"video"`

	XForwardedFor string `json:"x_forwarded_for"`
}

type AcceptResponse struct {
	Allowed bool `json:"allowed"`

	Multistream bool `json:"multistream,omitempty"`
	Simulcast   bool `json:"simulcast,omitempty"`
	Spotlight   bool `json:"spotlight,omitempty"`

	EventMetadata json.RawMessage `json:"event_metadata,omitempty"`

	Audio interface{} `json:"audio,omitempty"`
	Video interface{} `json:"video,omitempty"`

	TurnTCPOnly bool `json:"turn_tcp_only,omitempty"`
	TurnTLSOnly bool `json:"turn_tls_only,omitempty"`
}

type RejectResponse struct {
	Allowed bool   `json:"allowed"`
	Reason  string `json:"reason"`
}
