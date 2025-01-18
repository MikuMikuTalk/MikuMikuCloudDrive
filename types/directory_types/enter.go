package directory_types

import (
	"time"
)

type CreateDirectoryRequest struct {
	Token    string `header:"Authorization"` // 绑定到 Authorization 头
	Name     string `json:"name"`
	ParentID uint   `json:"parent_id"`
}

type CreateDirectoryResponse struct {
}

type DeleteDirectoryRequest struct {
	Token       string `header:"Authorization"`
	DirectoryID uint   `json:"directory_id"`
}

type DeleteDirectoryResponse struct {
}

// DirectoryItem represents a single item in directory (file or subdirectory)
type DirectoryItem struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`           // "file" or "directory"
	Size        int64     `json:"size,omitempty"` // Only for files
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsShared    bool      `json:"is_shared"`
	Permissions string    `json:"permissions"` // e.g. "rwxr-xr-x"
}

// DirectoryInfo contains metadata about the directory
type DirectoryInfo struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	TotalFiles  int       `json:"total_files"`
	TotalSize   int64     `json:"total_size"`
	IsRoot      bool      `json:"is_root"`
	IsShared    bool      `json:"is_shared"`
	Permissions string    `json:"permissions"`
}

type GetDirectoryInfoRequest struct {
	// Token       string `header:"Authorization"`
	UserID      uint `json:"user_id"`
	DirectoryID uint `json:"directory_id"`
}

// GetDirectoryInfoResponse represents the full response for directory info
type GetDirectoryInfoResponse struct {
	DirectoryInfo DirectoryInfo   `json:"directory_info"`
	Contents      []DirectoryItem `json:"contents"`
}
