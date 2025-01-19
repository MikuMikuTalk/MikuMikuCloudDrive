package directory_types

import (
	"fmt"
	"time"
)

type CreateDirectoryRequest struct {
	Token    string  `header:"Authorization"` // 绑定到 Authorization 头
	Name     string  `json:"name"`
	ParentID *uint   `json:"parent_id"`
	Path     *string `json:"path"`
}

type CreateDirectoryResponse struct {
}

func (r CreateDirectoryResponse) String() string {
	return fmt.Sprintf("CreateDirectoryResponse{}")
}

type DeleteDirectoryRequest struct {
	Token       string `header:"Authorization"`
	DirectoryID uint   `json:"directory_id"`
}

type DeleteDirectoryResponse struct {
}

func (r DeleteDirectoryResponse) String() string {
	return fmt.Sprintf("DeleteDirectoryResponse{}")
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

func (i DirectoryItem) String() string {
	return fmt.Sprintf("DirectoryItem{ID: %d, Name: %s, Type: %s, Size: %d, CreatedAt: %s, UpdatedAt: %s, IsShared: %t, Permissions: %s}",
		i.ID, i.Name, i.Type, i.Size, i.CreatedAt.Format(time.RFC3339), i.UpdatedAt.Format(time.RFC3339), i.IsShared, i.Permissions)
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

func (i DirectoryInfo) String() string {
	return fmt.Sprintf("DirectoryInfo{ID: %d, Name: %s, Path: %s, CreatedAt: %s, UpdatedAt: %s, TotalFiles: %d, TotalSize: %d, IsRoot: %t, IsShared: %t, Permissions: %s}",
		i.ID, i.Name, i.Path, i.CreatedAt.Format(time.RFC3339), i.UpdatedAt.Format(time.RFC3339), i.TotalFiles, i.TotalSize, i.IsRoot, i.IsShared, i.Permissions)
}

type GetDirectoryInfoRequest struct {
	Token       string `header:"Authorization"`
	DirectoryID uint   `json:"directory_id"`
}

// GetDirectoryInfoResponse represents the full response for directory info
type GetDirectoryInfoResponse struct {
	DirectoryInfo DirectoryInfo   `json:"directory_info"`
	Contents      []DirectoryItem `json:"contents"`
}

func (r GetDirectoryInfoResponse) String() string {
	contentsStr := "["
	for i, item := range r.Contents {
		if i > 0 {
			contentsStr += ", "
		}
		contentsStr += item.String()
	}
	contentsStr += "]"

	return fmt.Sprintf("GetDirectoryInfoResponse{DirectoryInfo: %s, Contents: %s}",
		r.DirectoryInfo.String(), contentsStr)
}
