package dropbox

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// Files client for files and folders.
type Files struct {
	*Client
}

// NewFiles client.
func NewFiles(config *Config) *Files {
	return &Files{
		Client: &Client{
			Config: config,
		},
	}
}

// WriteMode determines what to do if the file already exists.
type WriteMode string

// Supported write modes.
const (
	WriteModeAdd       WriteMode = "add"
	WriteModeOverwrite           = "overwrite"
)

// Dimensions specifies the dimensions of a photo or video.
type Dimensions struct {
	Width  uint64 `json:"width"`
	Height uint64 `json:"height"`
}

// GPSCoordinates specifies the GPS coordinate of a photo or video.
type GPSCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// PhotoMetadata specifies metadata for a photo.
type PhotoMetadata struct {
	Dimensions *Dimensions     `json:"dimensions,omitempty"`
	Location   *GPSCoordinates `json:"location,omitempty"`
	TimeTaken  time.Time       `json:"time_taken,omitempty"`
}

// VideoMetadata specifies metadata for a video.
type VideoMetadata struct {
	Dimensions *Dimensions     `json:"dimensions,omitempty"`
	Location   *GPSCoordinates `json:"location,omitempty"`
	TimeTaken  time.Time       `json:"time_taken,omitempty"`
	Duration   uint64          `json:"duration,omitempty"`
}

// MediaMetadata provides metadata for a photo or video.
type MediaMetadata struct {
	Photo *PhotoMetadata `json:"photo,omitempty"`
	Video *VideoMetadata `json:"video,omitempty"`
}

// MediaInfo provides additional information for a photo or video file.
type MediaInfo struct {
	Pending  bool           `json:"pending"`
	Metadata *MediaMetadata `json:"metadata,omitempty"`
}

// FileSharingInfo for a file which is contained in a shared folder.
type FileSharingInfo struct {
	ReadOnly             bool   `json:"read_only"`
	ParentSharedFolderID string `json:"parent_shared_folder_id"`
	ModifiedBy           string `json:"modified_by,omitempty"`
}

// Metadata for a file or folder.
type Metadata struct {
	Tag            string           `json:".tag"`
	Name           string           `json:"name"`
	PathLower      string           `json:"path_lower"`
	PathDisplay    string           `json:"path_display"`
	ClientModified time.Time        `json:"client_modified"`
	ServerModified time.Time        `json:"server_modified"`
	Rev            string           `json:"rev"`
	Size           uint64           `json:"size"`
	ID             string           `json:"id"`
	MediaInfo      *MediaInfo       `json:"media_info,omitempty"`
	SharingInfo    *FileSharingInfo `json:"sharing_info,omitempty"`
	ContentHash    string           `json:"content_hash,omitempty"`
}

// GetMetadataInput request input.
type GetMetadataInput struct {
	Path             string `json:"path"`
	IncludeMediaInfo bool   `json:"include_media_info"`
}

// GetMetadataOutput request output.
type GetMetadataOutput struct {
	Metadata
}

// GetMetadata returns the metadata for a file or folder.
func (c *Files) GetMetadata(in *GetMetadataInput) (out *GetMetadataOutput, err error) {
	body, err := c.call("/files/get_metadata", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// CreateFolderInput request input.
type CreateFolderInput struct {
	Path string `json:"path"`
}

// CreateFolderOutput request output.
type CreateFolderOutput struct {
	Name      string `json:"name"`
	PathLower string `json:"path_lower"`
	ID        string `json:"id"`
}

// CreateFolder creates a folder.
func (c *Files) CreateFolder(in *CreateFolderInput) (out *CreateFolderOutput, err error) {
	body, err := c.call("/files/create_folder", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// DeleteInput request input.
type DeleteInput struct {
	Path string `json:"path"`
}

// DeleteOutput request output.
type DeleteOutput struct {
	Metadata
}

// Delete a file or folder and its contents.
func (c *Files) Delete(in *DeleteInput) (out *DeleteOutput, err error) {
	body, err := c.call("/files/delete", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// PermanentlyDeleteInput request input.
type PermanentlyDeleteInput struct {
	Path string `json:"path"`
}

// PermanentlyDelete a file or folder and its contents.
func (c *Files) PermanentlyDelete(in *PermanentlyDeleteInput) (err error) {
	body, err := c.call("/files/delete", in)
	if err != nil {
		return
	}
	defer body.Close()

	return
}

// CopyInput request input.
type CopyInput struct {
	FromPath string `json:"from_path"`
	ToPath   string `json:"to_path"`
}

// CopyOutput request output.
type CopyOutput struct {
	Metadata
}

// Copy a file or folder to a different location.
func (c *Files) Copy(in *CopyInput) (out *CopyOutput, err error) {
	body, err := c.call("/files/copy", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// MoveInput request input.
type MoveInput struct {
	FromPath string `json:"from_path"`
	ToPath   string `json:"to_path"`
}

// MoveOutput request output.
type MoveOutput struct {
	Metadata
}

// Move a file or folder to a different location.
func (c *Files) Move(in *MoveInput) (out *MoveOutput, err error) {
	body, err := c.call("/files/move", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// RestoreInput request input.
type RestoreInput struct {
	Path string `json:"path"`
	Rev  string `json:"rev"`
}

// RestoreOutput request output.
type RestoreOutput struct {
	Metadata
}

// Restore a file to a specific revision.
func (c *Files) Restore(in *RestoreInput) (out *RestoreOutput, err error) {
	body, err := c.call("/files/restore", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// ListFolderInput request input.
type ListFolderInput struct {
	Path             string `json:"path"`
	Recursive        bool   `json:"recursive"`
	IncludeMediaInfo bool   `json:"include_media_info"`
	IncludeDeleted   bool   `json:"include_deleted"`
}

// ListFolderOutput request output.
type ListFolderOutput struct {
	Cursor  string `json:"cursor"`
	HasMore bool   `json:"has_more"`
	Entries []*Metadata
}

// ListFolder returns the metadata for a file or folder.
func (c *Files) ListFolder(in *ListFolderInput) (out *ListFolderOutput, err error) {
	in.Path = normalizePath(in.Path)

	body, err := c.call("/files/list_folder", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// ListFolderContinueInput request input.
type ListFolderContinueInput struct {
	Cursor string `json:"cursor"`
}

// ListFolderContinue pagenates using the cursor from ListFolder.
func (c *Files) ListFolderContinue(in *ListFolderContinueInput) (out *ListFolderOutput, err error) {
	body, err := c.call("/files/list_folder/continue", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// SearchMode determines how a search is performed.
type SearchMode string

// Supported search modes.
const (
	SearchModeFilename           SearchMode = "filename"
	SearchModeFilenameAndContent            = "filename_and_content"
	SearchModeDeletedFilename               = "deleted_filename"
)

// SearchMatchType represents the type of match made.
type SearchMatchType string

// Supported search match types.
const (
	SearchMatchFilename SearchMatchType = "filename"
	SearchMatchContent                  = "content"
	SearchMatchBoth                     = "both"
)

// SearchMatch represents a matched file or folder.
type SearchMatch struct {
	MatchType struct {
		Tag SearchMatchType `json:".tag"`
	} `json:"match_type"`
	Metadata *Metadata `json:"metadata"`
}

// SearchInput request input.
type SearchInput struct {
	Path       string     `json:"path"`
	Query      string     `json:"query"`
	Start      uint64     `json:"start,omitempty"`
	MaxResults uint64     `json:"max_results,omitempty"`
	Mode       SearchMode `json:"mode"`
}

// SearchOutput request output.
type SearchOutput struct {
	Matches []*SearchMatch `json:"matches"`
	More    bool           `json:"more"`
	Start   uint64         `json:"start"`
}

// Search for files and folders.
func (c *Files) Search(in *SearchInput) (out *SearchOutput, err error) {
	in.Path = normalizePath(in.Path)

	if in.Mode == "" {
		in.Mode = SearchModeFilename
	}

	body, err := c.call("/files/search", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// UploadInput request input.
type UploadInput struct {
	Path           string    `json:"path"`
	Mode           WriteMode `json:"mode"`
	AutoRename     bool      `json:"autorename"`
	Mute           bool      `json:"mute"`
	ClientModified string    `json:"client_modified,omitempty"`
	Reader         io.Reader `json:"-"`
}

// UploadOutput request output.
type UploadOutput struct {
	Metadata
}

// Upload a file smaller than 150MB.
func (c *Files) Upload(in *UploadInput) (out *UploadOutput, err error) {
	body, _, err := c.download("/files/upload", in, in.Reader)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// DownloadInput request input.
type DownloadInput struct {
	Path string `json:"path"`
}

// DownloadOutput request output.
type DownloadOutput struct {
	Body   io.ReadCloser
	Length int64
}

// Download a file.
func (c *Files) Download(in *DownloadInput) (out *DownloadOutput, err error) {
	body, l, err := c.download("/files/download", in, nil)
	if err != nil {
		return
	}

	out = &DownloadOutput{body, l}
	return
}

// ThumbnailFormat determines the format of the thumbnail.
type ThumbnailFormat string

const (
	// GetThumbnailFormatJPEG specifies a JPG thumbnail
	GetThumbnailFormatJPEG ThumbnailFormat = "jpeg"
	// GetThumbnailFormatPNG specifies a PNG thumbnail
	GetThumbnailFormatPNG = "png"
)

// ThumbnailSize determines the size of the thumbnail.
type ThumbnailSize string

const (
	// GetThumbnailSizeW32H32 specifies a size of 32 by 32 px
	GetThumbnailSizeW32H32 ThumbnailSize = "w32h32"
	// GetThumbnailSizeW64H64 specifies a size of 64 by 64 px
	GetThumbnailSizeW64H64 = "w64h64"
	// GetThumbnailSizeW128H128 specifies a size of 128 by 128 px
	GetThumbnailSizeW128H128 = "w128h128"
	// GetThumbnailSizeW640H480 specifies a size of 640 by 480 px
	GetThumbnailSizeW640H480 = "w640h480"
	// GetThumbnailSizeW1024H768 specifies a size of 1024 by 768 px
	GetThumbnailSizeW1024H768 = "w1024h768"
)

// GetThumbnailInput request input.
type GetThumbnailInput struct {
	Path   string          `json:"path"`
	Format ThumbnailFormat `json:"format"`
	Size   ThumbnailSize   `json:"size"`
}

// GetThumbnailOutput request output.
type GetThumbnailOutput struct {
	Body   io.ReadCloser
	Length int64
}

// GetThumbnail a thumbnail for a file. Currently thumbnails are only generated for the
// files with the following extensions: png, jpeg, png, tiff, tif, gif and bmp.
func (c *Files) GetThumbnail(in *GetThumbnailInput) (out *GetThumbnailOutput, err error) {
	body, l, err := c.download("/files/get_thumbnail", in, nil)
	if err != nil {
		return
	}

	out = &GetThumbnailOutput{body, l}
	return
}

// GetPreviewInput request input.
type GetPreviewInput struct {
	Path string `json:"path"`
}

// GetPreviewOutput request output.
type GetPreviewOutput struct {
	Body   io.ReadCloser
	Length int64
}

// GetPreview a preview for a file. Currently previews are only generated for the
// files with the following extensions: .doc, .docx, .docm, .ppt, .pps, .ppsx,
// .ppsm, .pptx, .pptm, .xls, .xlsx, .xlsm, .rtf
func (c *Files) GetPreview(in *GetPreviewInput) (out *GetPreviewOutput, err error) {
	body, l, err := c.download("/files/get_preview", in, nil)
	if err != nil {
		return
	}

	out = &GetPreviewOutput{body, l}
	return
}

// ListRevisionsInput request input.
type ListRevisionsInput struct {
	Path  string `json:"path"`
	Limit uint64 `json:"limit,omitempty"`
}

// ListRevisionsOutput request output.
type ListRevisionsOutput struct {
	IsDeleted bool
	Entries   []*Metadata
}

// ListRevisions gets the revisions of the specified file.
func (c *Files) ListRevisions(in *ListRevisionsInput) (out *ListRevisionsOutput, err error) {
	body, err := c.call("/files/list_revisions", in)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// Normalize path so people can use "/" as they expect.
func normalizePath(s string) string {
	if s == "/" {
		return ""
	}
	return s
}

const hashBlockSize = 4 * 1024 * 1024

// ContentHash returns the Dropbox content_hash for a io.Reader.
// See https://www.dropbox.com/developers/reference/content-hash
func ContentHash(r io.Reader) (string, error) {
	buf := make([]byte, hashBlockSize)
	resultHash := sha256.New()
	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	if n > 0 {
		bufHash := sha256.Sum256(buf[:n])
		resultHash.Write(bufHash[:])
	}
	for n == hashBlockSize && err == nil {
		n, err = r.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n > 0 {
			bufHash := sha256.Sum256(buf[:n])
			resultHash.Write(bufHash[:])
		}
	}
	return fmt.Sprintf("%x", resultHash.Sum(nil)), nil
}

// FileContentHash returns the Dropbox content_hash for a local file.
// See https://www.dropbox.com/developers/reference/content-hash
func FileContentHash(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return ContentHash(f)
}
