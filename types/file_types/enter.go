package file_types

type DownloadFileRequest struct {
	FileID string `param:"file_id"` // 文件唯一标识
}

type DownloadFileResponse struct {
	FileID   string `json:"file_id"`   // 文件唯一标识
	Url      string `json:"url"`       // 文件下载地址
	FileName string `json:"file_name"` // 文件名称
	FileHash string `json:"file_hash"` // 文件哈希值
}
