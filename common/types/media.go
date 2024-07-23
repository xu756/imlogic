package types

// 图片类型
type ImageType struct {
	UID string `json:"uid"`
	URL string `json:"url"`
}

// 视频类型
type VideoType struct {
	UID      string `json:"uid"`
	URL      string `json:"url"`
	Duration int64  `json:"duration"`
}

// 文件类型
type FileType struct {
	UID string `json:"uid"`
	URL string `json:"url"`
}

// 音频类型
type AudioType struct {
	UID      string `json:"uid"`
	URL      string `json:"url"`
	Duration int64  `json:"duration"`
}
