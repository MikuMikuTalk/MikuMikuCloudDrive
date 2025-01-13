package chunk_process_types

import "mime/multipart"

type commonModel struct {
	FileName    string `json:"fileName"`
	FileMD5     string `json:"fileMD5"`
	TotalChunks int    `json:"totalChunks"`
}
type GetUploadedChunksRequest struct {
	commonModel
}
type GetUploadedChunksResponse struct {
	ChunksArray []int `json:"chunksArray"`
}

type MergeChunksRequest struct {
	commonModel
}

type MergeChunksResponse struct{}

type ChunkUploadRequest struct {
	File        *multipart.FileHeader `form:"file" binding:"required"`         // 文件字段
	ChunkIndex  int                   `form:"chunk_index" binding:"required"`  // 切片索引
	TotalChunks int                   `form:"total_chunks" binding:"required"` // 总分片数量
	FileMD5     string                `form:"file_md5" binding:"required"`     //文件md5值
}
type ChunkUploadedResponse struct{}
