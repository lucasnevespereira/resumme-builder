package json

import (
	jsoniter "github.com/json-iterator/go"
)

// Copy from https://github.com/gin-gonic/gin/blob/master/internal/json/jsoniter.go
var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gin/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gin/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gin/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gin/json package.
	NewEncoder = json.NewEncoder
	// MarshalToString is exported by gin/json package.
	MarshalToString = json.MarshalToString
)
