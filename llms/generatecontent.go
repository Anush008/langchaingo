package llms

import (
	"encoding/json"

	"github.com/tmc/langchaingo/schema"
)

// MessageContent is the content of a message sent to a LLM. It has a role and a
// sequence of parts. For example, it can represent one message in a chat
// session sent by the user, in which case Role will be
// schema.ChatMessageTypeHuman and Parts will be the sequence of items sent in
// this specific message.
type MessageContent struct {
	Role  schema.ChatMessageType
	Parts []ContentPart
}

// ContentPart is an interface all parts of content have to implement.
type ContentPart interface {
	isPart()
}

// TextContent is content with some text.
type TextContent struct {
	Text string
}

func (tc TextContent) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		"type": "text",
		"text": tc.Text,
	}
	return json.Marshal(m)
}

func (TextContent) isPart() {}

// ImageURLContent is content with an URL pointing to an image.
type ImageURLContent struct {
	URL string
}

func (iuc ImageURLContent) MarshalJSON() ([]byte, error) {
	m := map[string]any{
		"type": "image_url",
		"image_url": map[string]string{
			"url": iuc.URL,
		},
	}
	return json.Marshal(m)
}

func (ImageURLContent) isPart() {}

// BinaryContent is content holding some binary data with a MIME type.
type BinaryContent struct {
	MIMEType string
	Data     []byte
}

func (BinaryContent) isPart() {}

// ContentResponse is the response returned by a GenerateContent call.
// It can potentially return multiple content choices.
type ContentResponse struct {
	Choices []*ContentChoice
}

// ContentChoice is one of the response choices returned by GenerateContent
// calls.
type ContentChoice struct {
	// Content is the textual content of a response
	Content string

	// StopReason is the reason the model stopped generating output.
	StopReason string

	// GenerationInfo is arbitrary information the model adds to the response.
	GenerationInfo map[string]any

	// FuncCall is non-nil when the model asks to invoke a function/tool.
	FuncCall *schema.FunctionCall
}

// TextParts is a helper function to create a MessageContent with a role and a
// list of text parts.
func TextParts(role schema.ChatMessageType, parts ...string) MessageContent {
	result := MessageContent{
		Role:  role,
		Parts: []ContentPart{},
	}
	for _, part := range parts {
		result.Parts = append(result.Parts, TextContent{
			Text: part,
		})
	}
	return result
}
