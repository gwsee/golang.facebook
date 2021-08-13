package model

import "git.zx-tech.net/pengfeng/facebook/model"

//AdCreative 素材
type AdCreative struct {
	EffectiveObjectStoryID string          `json:"effective_object_story_id"`
	ObjectStoryID          string          `json:"object_story_id"`
	ImageUrl               string          `json:"image_url"`
	ID                     string          `json:"id"`
	ObjectStorySpec        ObjectStorySpec `json:"object_story_spec"`
}

//AdCreatives 素材-列表
type AdCreatives struct {
	Data   []AdCreative `json:"data"`
	Paging model.Paging `json:"paging"`
}

type ObjectStorySpec struct {
	PageID    string    `json:"page_id"`
	LinkData  LinkData  `json:"link_data"`
	PhotoData PhotoData `json:"photo_data"`
	VideoData VideoData `json:"video_data"`
}
type LinkData struct {
	Link            string       `json:"link"`
	Message         string       `json:"message"`
	AttachmentStyle string       `json:"attachment_style"`
	ImageHash       string       `json:"image_hash"`
	CallToAction    CallToAction `json:"call_to_action"`
}
type PhotoData struct {
	Caption      string       `json:"caption"`
	ImageHash    string       `json:"image_hash"`
	CallToAction CallToAction `json:"call_to_action"`
}
type VideoData struct {
	VideoID              string       `json:"video_id"`
	Title                string       `json:"title"`
	Message              string       `json:"message"`
	CallToAction         CallToAction `json:"call_to_action"`
	ImageHash            string       `json:"image_hash"`
	VideoThumbnailSource string       `json:"video_thumbnail_source"`
}
type CallToAction struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}
type Value struct {
	Link       string `json:"link"`
	LinkFormat string `json:"link_format"`
}
