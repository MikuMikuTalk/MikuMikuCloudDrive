package directory_types

import (
	"time"
)

type CreateDirectoryRequest struct {
	Name     string `json:"name"`
	ParentID *uint  `json:"parent_id"`
}

type CreateDirectoryResponse struct {
}

type DeleteDirectoryRequest struct {
	DirectoryID uint `json:"directory_id"`
}

type DeleteDirectoryResponse struct {
}

// DirectoryItem represents a single item in directory (file or subdirectory)
type DirectoryItem struct {
	ID          uint      `json:"id" testlog:"ID"`
	Name        string    `json:"name" testlog:"名称"`
	Type        string    `json:"type" testlog:"类型"`           // "file" or "directory"
	Size        int64     `json:"size,omitempty" testlog:"大小"` // Only for files
	CreatedAt   time.Time `json:"created_at" testlog:"创建时间"`
	UpdatedAt   time.Time `json:"updated_at" testlog:"更新时间"`
	IsShared    bool      `json:"is_shared" testlog:"是否共享"`
	Permissions string    `json:"permissions" testlog:"权限"` // e.g. "rwxr-xr-x"
}

// DirectoryInfo contains metadata about the directory
type DirectoryInfo struct {
	ID          uint      `json:"id" testlog:"目录ID"`
	Name        string    `json:"name" testlog:"目录名称"`
	Path        string    `json:"path" testlog:"目录路径"`
	CreatedAt   time.Time `json:"created_at" testlog:"创建时间"`
	UpdatedAt   time.Time `json:"updated_at" testlog:"更新时间"`
	TotalFiles  int       `json:"total_files" testlog:"文件总数"`
	TotalSize   int64     `json:"total_size" testlog:"总大小"`
	IsRoot      bool      `json:"is_root" testlog:"是否为根目录"`
	IsShared    bool      `json:"is_shared" testlog:"是否共享"`
	Permissions string    `json:"permissions" testlog:"权限"`
}

type GetDirectoryInfoRequest struct {
	DirectoryID string `json:"directory_id"`
}

// GetDirectoryInfoResponse represents the full response for directory info
type GetDirectoryInfoResponse struct {
	DirectoryInfo DirectoryInfo   `json:"directory_info"`
	Contents      []DirectoryItem `json:"contents"`
}

type GetDirectoryListRequest struct {
}

type GetDirectoryListResponse struct {
	Directories []DirectoryInfo `json:"directories"`
}
