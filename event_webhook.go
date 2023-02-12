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

type ArchiveStartedData struct {
	RecordingID string `json:"recording_id"`

	ChannelID    string `json:"channel_id"`
	SessionID    string `json:"session_id"`
	ClientID     string `json:"client_id"`
	BundleID     string `json:"bundle_id"`
	ConnectionID string `json:"connection_id"`

	Audio          bool   `json:"audio"`
	AudioCodecType string `json:"audio_codec_type"`
	AudioBitRate   *int32 `json:"audio_bit_rate"`

	Video          bool   `json:"video"`
	VideoCodecType string `json:"video_codec_type"`
	VideoBitRate   int32  `json:"video_bit_rate"`
	VideoHeight    int32  `json:"video_height"`
	VideoWidth     int32  `json:"video_width"`

	// unix-time
	StartTime       int64     `json:"start_time"`
	StartTimeOffset int64     `json:"start_time_offset"`
	StartTimestamp  time.Time `json:"start_timestamp"`

	// unix-time
	CreatedAt int64 `json:"created_at"`
}

type ArchiveAvailableData struct {
	ArchiveStartedData

	Filename string `json:"filename"`
	FilePath string `json:"file_path"`

	MetadataFilename string `json:"metadata_filename"`
	MetadataFilePath string `json:"metadata_file_path"`

	// unix-time
	StopTime       int64     `json:"stop_time"`
	StopTimeOffset int64     `json:"stop_time_offset"`
	StopTimestamp  time.Time `json:"stop_timestamp"`

	Size int64 `json:"size"`

	Stats json.RawMessage `json:"stats"`
}

type SplitArchiveAvailableData struct {
	ArchiveAvailableData
	SplitIndex string `json:"split_index"`
}

type SplitArchiveEndData struct {
	SplitLastIndex string `json:"split_last_index"`

	RecordingID string `json:"recording_id"`

	ChannelID    string `json:"channel_id"`
	SessionID    string `json:"session_id"`
	ClientID     string `json:"client_id"`
	BundleID     string `json:"bundle_id"`
	ConnectionID string `json:"connection_id"`

	Audio          bool   `json:"audio"`
	AudioCodecType string `json:"audio_codec_type"`
	AudioBitRate   *int32 `json:"audio_bit_rate"`

	Video          bool   `json:"video"`
	VideoCodecType string `json:"video_codec_type"`
	VideoBitRate   int32  `json:"video_bit_rate"`

	Filename string `json:"filename"`
	FilePath string `json:"file_path"`

	// unix-time
	StartTime       int64     `json:"start_time"`
	StartTimeOffset int64     `json:"start_time_offset"`
	StartTimestamp  time.Time `json:"start_timestamp"`

	// unix-time
	StopTime       int64     `json:"stop_time"`
	StopTimeOffset int64     `json:"stop_time_offset"`
	StopTimestamp  time.Time `json:"stop_timestamp"`
}

type ArchiveFailed struct {
	RecordingID string `json:"recording_id"`

	Audio          bool   `json:"audio"`
	AudioCodecType string `json:"audio_codec_type"`
	AudioBitRate   *int32 `json:"audio_bit_rate"`

	Video          bool   `json:"video"`
	VideoCodecType string `json:"video_codec_type"`
	VideoBitRate   int32  `json:"video_bit_rate"`

	Filename string `json:"filename"`
	FilePath string `json:"file_path"`

	// unix-time
	StartTime       int64     `json:"start_time"`
	StartTimeOffset int64     `json:"start_time_offset"`
	StartTimestamp  time.Time `json:"start_timestamp"`

	// unix-time
	StopTime       int64     `json:"stop_time"`
	StopTimeOffset int64     `json:"stop_time_offset"`
	StopTimestamp  time.Time `json:"stop_timestamp"`
}
