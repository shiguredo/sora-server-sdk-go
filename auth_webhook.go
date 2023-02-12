package sora

import (
	"encoding/json"
	"time"
)

type AuthWebhookRequest struct {
	Timestamp time.Time `json:"timestamp"`

	ID string `json:"id"`

	Version  string `json:"version"`
	Label    string `json:"label"`
	NodeName string `json:"node_name"`

	Multistream  bool   `json:"multistream"`
	Simulcast    bool   `json:"simulcast"`
	SimulcastRid string `json:"simulcast_rid"`
	Spotlight    bool   `json:"spotlight"`

	Audio          bool   `json:"audio"`
	AudioCodecType string `json:"audio_codec_type"`
	// AudioBitRate は入ってこない場合があるので nil を許容する
	AudioBitRate   *int   `json:"audio_bit_rate"`
	Video          bool   `json:"video"`
	VideoCodecType string `json:"video_codec_type"`
	VideoBitRate   int    `json:"video_bit_rate"`

	DataChannelSignaling      bool                `json:"data_channel_signaling"`
	IgnoreDisconnectWebSocket bool                `json:"ignore_disconnect_websocket"`
	DataChannels              []DataChannelParams `json:"data_channels"`

	Role         string `json:"role"`
	ChannelID    string `json:"channel_id"`
	ClientID     string `json:"client_id"`
	BundleID     string `json:"bundle_id"`
	ConnectionID string `json:"connection_id"`

	Metadata     json.RawMessage `json:"metadata"`
	AuthMetadata json.RawMessage `json:"auth_metadata"`

	E2EE bool `json:"e2ee"`

	ChannelConnections         uint `json:"channel_connetions"`
	ChannelSendrecvConnections uint `json:"channel_sendrecv_connections"`
	ChannelSendonlyConnections uint `json:"channel_sendonly_connections"`
	ChannelRecvonlyConnections uint `json:"channel_recvonly_connections"`

	SoraClient SoraClient `json:"sora_client"`
}

type SoraClient struct {
	Type        string `json:"type"`
	Raw         string `json:"raw"`
	Version     string `json:"version"`
	CommitShort string `json:"commit_short"`
	Environment string `json:"environment"`
	Libwebrtc   string `json:"libwebrtc"`
}

type AuthWebhookSuccessResponse struct {
	Allowed bool   `json:"allowed"`
	Reason  string `json:"reason,omitempty"`

	Multistream bool `json:"multistream,omitempty"`

	Simulcast          bool                `json:"simulcast,omitempty"`
	SimulcastRid       string              `json:"simulcast_rid,omitempty"`
	SimulcastEncodings []SimulcastEncoding `json:"simulcast_encodings,omitempty"`

	Spotlight          bool                `json:"spotlight,omitempty"`
	SpotlightNumber    uint                `json:"spotlight_number,omitempty"`
	SpotlightEncodings []SimulcastEncoding `json:"spotlight_encodings,omitempty"`

	Audio          bool   `json:"audio,omitempty"`
	AudioCodecType string `json:"audio_codec_type,omitempty"`
	AudioBitRate   int    `json:"audio_bit_rate,omitempty"`
	Video          bool   `json:"video,omitempty"`
	VideoCodecType string `json:"video_codec_type,omitempty"`
	VideoBitRate   int    `json:"video_bit_rate,omitempty"`

	DataChannelSignaling      bool                `json:"data_channel_signaling,omitempty"`
	IgnoreDisconnectWebSocket bool                `json:"ignore_disconnect_websocket,omitempty"`
	DataChannels              []DataChannelParams `json:"data_channels,omitempty"`

	ClientID string `json:"client_id,omitempty"`
	BundleID string `json:"bundle_id,omitempty"`

	SignalingNotify bool `json:"signaling_notify,omitempty"`

	IPv4Address string `json:"ipv4_address,omitempty"`
	IPv6Address string `json:"ipv6_address,omitempty"`

	TurnFqdn    string `json:"turn_fqdn,omitempty"`
	TurnTlsFqdn string `json:"turn_tls_fqdn,omitempty"`
	TurnTcpOnly bool   `json:"turn_tcp_only,omitempty"`
	TurnTlsOnly bool   `json:"turn_tls_only,omitempty"`

	H264ProfileLevelID string `json:"h264_profile_level_id,omitempty"`

	UserAgent bool `json:"user_agent,omitempty"`

	Metadata      json.RawMessage `json:"metadata,omitempty"`
	EventMetadata json.RawMessage `json:"event_metadata,omitempty"`
}

type SimulcastEncoding struct {
	Rid                   string  `json:"rid"`
	Active                bool    `json:"active,omitempty"`
	ScaleResolutionDownBy float64 `json:"scaleResolutionDownBy,omitempty"`
	MaxBitrate            float64 `json:"maxBitrate,omitempty"`
	MaxFramerate          float64 `json:"maxFramerate,omitempty"`
	AdaptivePtime         bool    `json:"adaptivePtime"`
}

type AuthWebhookErrorResponse struct {
	Allowed bool   `json:"allowed"`
	Reason  string `json:"reason"`
}

type DataChannelParams struct {
	Label             string `json:"label,omitempty"`
	Direction         string `json:"direction,omitempty"`
	Compress          bool   `json:"compress"`
	Ordered           bool   `json:"ordered"`
	MaxRetransmit     int    `json:"max_retransmit,omitempty"`
	MaxPacketLifeTime int    `json:"max_packet_life_time,omitempty"`
}
