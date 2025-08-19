package utils

type Size struct {
	W int32 `json:"w"`
	H int32 `json:"h"`
}

type TextureReference struct {
	ID    string `json:"id"`
	Index int    `json:"index"`
}

type Collision string
