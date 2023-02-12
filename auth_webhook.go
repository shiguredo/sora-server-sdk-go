package sora

import (
	"encoding/json"
	"time"
)

type SimulcastRid string

const (
	SimulcastRid0 SimulcastRid = "r0"
	SimulcastRid1 SimulcastRid = "r1"
	SimulcastRid2 SimulcastRid = "r2"
)

type Role string

const (
	RoleSendrecv Role = "sendrecv"
	RoleSendonly Role = "sendonly"
	RoleRecvonly Role = "recvonly"
)

type AudioCodecType string

const (
	AudioCodecTypeOpus      AudioCodecType = "OPUS"
	AudioCodecTypeMultiopus AudioCodecType = "MULTIOPUS"
	AudioCodecTypeLyra      AudioCodecType = "LYRA"
)

type VideoCodecType string

const (
	VideoCodecTypeVP8  VideoCodecType = "VP8"
	VideoCodecTypeVP9  VideoCodecType = "VP9"
	VideoCodecTypeAV1  VideoCodecType = "AV1"
	VideoCodecTypeH264 VideoCodecType = "H264"
	VideoCodecTypeH265 VideoCodecType = "H265"
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

	Audio          bool            `json:"audio"`
	AudioCodecType *AudioCodecType `json:"audio_codec_type"`
	AudioBitRate   *int32          `json:"audio_bit_rate"`
	Video          bool            `json:"video"`
	VideoCodecType *VideoCodecType `json:"video_codec_type"`
	VideoBitRate   *int32          `json:"video_bit_rate"`

	DataChannelSignaling      bool                `json:"data_channel_signaling"`
	IgnoreDisconnectWebSocket bool                `json:"ignore_disconnect_websocket"`
	DataChannels              []DataChannelParams `json:"data_channels"`

	Role         Role   `json:"role"`
	ChannelID    string `json:"channel_id"`
	ClientID     string `json:"client_id"`
	BundleID     string `json:"bundle_id"`
	ConnectionID string `json:"connection_id"`

	Metadata     json.RawMessage `json:"metadata"`
	AuthMetadata json.RawMessage `json:"auth_metadata"`

	E2EE bool `json:"e2ee"`

	ChannelConnections         int32 `json:"channel_connections"`
	ChannelSendrecvConnections int32 `json:"channel_sendrecv_connections"`
	ChannelSendonlyConnections int32 `json:"channel_sendonly_connections"`
	ChannelRecvonlyConnections int32 `json:"channel_recvonly_connections"`

	SoraClient SoraClient `json:"sora_client"`

	AudioStreamingLanguageCode string `json:"audio_streaming_language_code"`
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
	SpotlightNumber    int32               `json:"spotlight_number,omitempty"`
	SpotlightEncodings []SimulcastEncoding `json:"spotlight_encodings,omitempty"`

	Audio          bool            `json:"audio,omitempty"`
	AudioCodecType *AudioCodecType `json:"audio_codec_type,omitempty"`
	AudioBitRate   *int32          `json:"audio_bit_rate,omitempty"`
	Video          bool            `json:"video,omitempty"`
	VideoCodecType *VideoCodecType `json:"video_codec_type,omitempty"`
	VideoBitRate   *int32          `json:"video_bit_rate,omitempty"`

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
	AdaptivePtime         bool    `json:"adaptivePtime,omitempty"`
	ScalabilityMode       string  `json:"scalabilityMode,omitempty"`
}

type AuthWebhookFailureResponse struct {
	Allowed bool   `json:"allowed"`
	Reason  string `json:"reason"`
}

type DataChannelParams struct {
	Label             string `json:"label,omitempty"`
	Direction         string `json:"direction,omitempty"`
	Compress          bool   `json:"compress"`
	Ordered           bool   `json:"ordered"`
	MaxRetransmit     int32  `json:"max_retransmit,omitempty"`
	MaxPacketLifeTime int32  `json:"max_packet_life_time,omitempty"`
}
