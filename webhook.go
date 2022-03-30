package sora

type AudioParams struct {
	BitRate   int    `json:"bit_rate,omitempty"`
	CodecType string `json:"codec_type,omitempty"`
}

type VideoParams struct {
	BitRate   uint   `json:"bit_rate,omitempty"`
	CodecType string `json:"codec_type,omitempty"`
}
