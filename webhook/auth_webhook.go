package webhook

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

	DataChannelSignaling      bool `json:"data_channel_signaling"`
	IgnoreDisconnectWebSocket bool `json:"ignore_disconnect_websocket"`

	DataChannels []DataChannelParams `json:"data_channels"`

	ChannelConnections         uint `json:"channel_connetions"`
	ChannelSendrecvConnections uint `json:"channel_sendrecv_connections"`
	ChannelSendonlyConnections uint `json:"channel_sendonly_connections"`
	ChannelRecvonlyConnections uint `json:"channel_recvonly_connections"`

	Audio AudioParams `json:"audio"`
	Video VideoParams `json:"video"`

	E2EE bool `json:"e2ee"`

	SoraClient SoraClient `json:"sora_client"`

	// undoc
	XForwardedFor string `json:"x_forwarded_for"`
}

type DataChannelParams struct {
	Label             string `json:"label"`
	Direction         string `json:"direction"`
	Protocol          string `json:"protocol"`
	Ordered           bool   `json:"ordered"`
	Compress          bool   `json:"compress"`
	MaxPacketLifeTime int    `json:"max_packet_life_time"`
	MaxRetransmits    int    `json:"max_retransmits"`
}

type SoraClient struct {
	Environment string `json:"environment"`
	Raw         string `json:"raw"`
	Type        string `json:"type"`
	Version     string `json:"version"`
	CommitShort string `json:"commit_short"`
	Libwebrtc   string `json:"libwebrtc"`
}

type SuccessResponse struct {
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

type ErrorResponse struct {
	Allowed bool   `json:"allowed"`
	Reason  string `json:"reason"`
}
