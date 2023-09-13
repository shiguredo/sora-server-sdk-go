package sora

import (
	"encoding/json"
	"time"
)

// 2024.2.0 で廃止されます
type EventRecordingStartedWebhookRequest struct {
}

type EventRecordingStartedWebhookResponse struct {
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
