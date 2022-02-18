package photo

type PhotoLoadDto struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	OwnerId      int    `json:"owner"`
	AlbumId      int    `json:"albumId"`
}
