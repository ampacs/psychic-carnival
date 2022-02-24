package message

// Message represents a simple registration with a picture ID and a weight
type Message struct {
	PictureId string `json:"picture_id"`
	Weight    int    `json:"weight"`
}
