package goWebhook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Thumbnail struct {
	URL string `json:"url"`
}
type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type Embed struct {
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Color     int       `json:"color"`
	Timestamp string    `json:"timestamp"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Fields    []Fields  `json:"fields"`
	Footer    Footer    `json:"footer"`
}
type Webhook struct {
	Username  string  `json:"username"`
	AvatarURL string  `json:"avatar_url"`
	Embeds    []Embed `json:"embeds"`
}

// creates and returns the webhook struct

func CreateWebhook() Webhook {
	Wh := Webhook{
		Username:  "",
		AvatarURL: "",
		Embeds: []Embed{
			{
				Title:     "",
				URL:       "",
				Color:     16411130,
				Thumbnail: Thumbnail{URL: ""},
				Fields:    []Fields{},
			},
		},
	}

	return Wh
}

// add a username to the webhook

func (wh *Webhook) SetWebhookUsername(username string) {
	wh.Username = username
}

// add a avatar to the webhook

func (wh *Webhook) SetWebhookAvatarURL(avatarURL string) {
	wh.AvatarURL = avatarURL
}

// add a title to the webhook

func (wh *Webhook) SetTitle(title string) {
	wh.Embeds[0].Title = title
}

// add a url to the webhook

func (wh *Webhook) SetURL(URL string) {
	wh.Embeds[0].URL = URL
}

// set the color of the webhook
func (wh *Webhook) SetColor(color int) {
	wh.Embeds[0].Color = color
}

// adds a thumbnail

func (wh *Webhook) SetThumbnailURL(thumbnailURL string) {
	wh.Embeds[0].Thumbnail.URL = thumbnailURL
}

// simple function to add fields

func (wh *Webhook) AddField(title string, value string, inline bool) {

	newField := Fields{
		Name:   title,
		Value:  value,
		Inline: inline,
	}

	wh.Embeds[0].Fields = append(wh.Embeds[0].Fields, newField)

}

// final function encodes webhook data and then posts to webhook provided via function args
func (wh Webhook) SendWebhook(url string) *http.Response {
	client := &http.Client{}

	webhookData, err := json.Marshal(wh)

	if err != nil {
		panic("Eror encoding webhook data")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(webhookData))

	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		panic("Error creating webhook request")
	}

	webhookPost, err := client.Do(req)

	if err != nil {
		panic("Error posting webhook")
	}

	if webhookPost.StatusCode == 204 {
		return webhookPost
	}
	return webhookPost
}
