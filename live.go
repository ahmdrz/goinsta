package goinsta

import (
	"encoding/json"
	"fmt"
)

type LiveCreateResponse struct {
	BroadcastID                                   int    `json:"broadcast_id"`
	UploadUrl                                     string `json:"upload_url"`
	// MaxTimeInSeconds                              int `json:"max_time_in_seconds"`
	// SpeedTestUiTimeout                            int `json:"speed_test_ui_timeout"`
	// StreamNetworkSpeedTestPayloadChunkSizeInBytes int `json:"stream_network_speed_test_payload_chunk_size_in_bytes"`
	// StreamNetworkSpeedTestPayloadSizeInBytes      int `json:"stream_network_speed_test_payload_size_in_bytes"`
	// StreamNetworkSpeedTestPayloadTimeoutInSeconds int `json:"stream_network_speed_test_payload_timeout_in_seconds"`
	// SpeedTestMinimumBandwidthThreshold            int `json:"speed_test_minimum_bandwidth_threshold"`
	// SpeedTestRetryMaxCount                        int `json:"speed_test_retry_max_count"`
	// SpeedTestRetryTimeDelay                       int `json:"speed_test_retry_time_delay"`
	// DisableSpeedTest                              int `json:"disable_speed_test"`
	// StreamVideoAllowBFrames                       int `json:"stream_video_allow_b_frames"`
	// StreamVideoWidth                              int `json:"stream_video_width"`
	// StreamVideoBitRate                            int `json:"stream_video_bit_rate"`
	// StreamVideoFps                                int `json:"stream_video_fps"`
	// StreamAudioBitRate                            int `json:"stream_audio_bit_rate"`
	// StreamAudioSampleRate                         int `json:"stream_audio_sample_rate"`
	// StreamAudioChannels                           int `json:"stream_audio_channels"`
	// HeartbeatInterval                             int `json:"heartbeat_interval"`
	// BroadcasterUpdateFrequency                    int `json:"broadcaster_update_frequency"`
	// StreamVideoAdaptiveBitrateConfig              int `json:"stream_video_adaptive_bitrate_config"`
	// StreamNetworkConnectionRetryCount             int `json:"stream_network_connection_retry_count"`
	// StreamNetworkConnectionRetryDelayInSeconds    int `json:"stream_network_connection_retry_delay_in_seconds"`
	// ConnectWith1rtt                               int `json:"connect_with_1rtt"`
	// AvcRtmpPayload                                int `json:"avc_rtmp_payload"`
	// AllowResolutionChange                         int `json:"allow_resolution_change"`
}

type LiveStartResponse struct {
	MediaID string `json:"media_id"`
}

type Live struct {
	inst *Instagram
}

func newLive(inst *Instagram) *Live {
	return &Live{inst: inst}
}

func (live *Live) Create(width int, height int) (*LiveCreateResponse, error) {
	insta := live.inst
	response := &LiveCreateResponse{}

	data, err := insta.prepareData(
		map[string]interface{}{
			"preview_width":  width,
			"preview_height": height,
			"broadcast_type": "RTMP_SWAP_ENABLED",
			"internal_only":  0,
		},
	)
	if err != nil {
		return response, err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlLiveCreate,
			Query:    generateSignature(data),
			IsPost:   false,
		},
	)

	if err != nil {
		return response, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, response)

	return response, err
}

func (live *Live) Start(broadcastId string) (*LiveStartResponse, error) {
	insta := live.inst
	response := &LiveStartResponse{}

	data, err := insta.prepareData()
	if err != nil {
		return response, err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlLiveStart, broadcastId),
			Query:    generateSignature(data),
			IsPost:   false,
		},
	)

	err = json.Unmarshal(body, response)

	return response, err
}

func (live *Live) End(broadcastId string) error {
	insta := live.inst

	data, err := insta.prepareData(
		map[string]interface{}{
			"end_after_copyright_warning": false,
		},
	)
	if err != nil {
		return err
	}

	_, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlLiveEnd, broadcastId),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)

	return err
}
