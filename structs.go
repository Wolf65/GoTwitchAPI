package gtapi

import (
	"net/http"
	"time"
)
//Session contains field permanent date
type Session struct {
	Client   *http.Client
	URL      string
	ClientID string
}
//GetStreamByUser contains return fields
type GetStreamByUserOutput struct {
	Stream *StreamByUser         `json:"stream"`
	Links  *LinksGetStreamByUser `json:"_links"`
}
//StreamByUser contains return fields
type StreamByUser struct {
	ID          int64             `json:"_id"`
	Game        string            `json:"game"`
	Viewers     int               `json:"viewers"`
	VideoHeight int               `json:"video_height"`
	AverageFps  float64           `json:"average_fps"`
	Delay       int               `json:"delay"`
	CreatedAt   time.Time         `json:"created_at"`
	IsPlaylist  bool              `json:"is_playlist"`
	StreamType  string            `json:"stream_type"`
	Preview     *Preview           `json:"preview"`
	Channel     *ChannelByUser     `json:"channel"`
	Links       *LinksStreamByUser `json:"_links"`
}
//Preview contains url 
type Preview struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Template string `json:"template"`
}
//ChannelByUser contains fields return call StreamByUser
type ChannelByUser struct {
	Mature                       bool               `json:"mature"`
	Partner                      bool               `json:"partner"`
	Status                       string             `json:"status"`
	BroadcasterLanguage          string             `json:"broadcaster_language"`
	DisplayName                  string             `json:"display_name"`
	Game                         string             `json:"game"`
	Language                     string             `json:"language"`
	ID                           int                `json:"_id"`
	Name                         string             `json:"name"`
	CreatedAt                    time.Time          `json:"created_at"`
	UpdatedAt                    time.Time          `json:"updated_at"`
	Delay                        int                `json:"delay"`
	Logo                         string             `json:"logo"`
	Banner                       string             `json:"banner"`
	VideoBanner                  string             `json:"video_banner"`
	Background                   string             `json:"background"`
	ProfileBanner                string             `json:"profile_banner"`
	ProfileBannerBackgroundColor string             `json:"profile_banner_background_color"`
	URL                          string             `json:"url"`
	Views                        int                `json:"views"`
	Followers                    int                `json:"followers"`
	Links                        *LinksChannelByUser `json:"_links"`
}
//LinksChannelByUser contains fields return call StreamByUser in ChannelByUser
type LinksChannelByUser struct {
	Self          string `json:"self"`
	Follows       string `json:"follows"`
	Commercial    string `json:"commercial"`
	StreamKey     string `json:"stream_key"`
	Chat          string `json:"chat"`
	Features      string `json:"features"`
	Subscriptions string `json:"subscriptions"`
	Editors       string `json:"editors"`
	Teams         string `json:"teams"`
	Videos        string `json:"videos"`
}
//LinksStreamByUser contains fields return call GetStreamByUser in StreamByUser
type LinksStreamByUser struct {
	Self string `json:"self"`
}
//LinksGetStreamByUser contains fields return call GetStreamByUser 
type LinksGetStreamByUser struct {
	Self    string `json:"self"`
	Channel string `json:"channel"`
}
//GetTopStreamOutput contains return fields 
type GetTopStreamOutput struct {
	Total   int         `json:"_total,omitempty"`
	Streams *[]TopStream `json:"streams,omitempty"`
}
//TopStream contains return fields
type TopStream struct {
	ID                int64      `json:"_id"`
	Game              string     `json:"game"`
	BroadcastPlatform string     `json:"broadcast_platform"`
	CommunityID       string     `json:"community_id"`
	CommunityIds      []string   `json:"community_ids"`
	Viewers           int        `json:"viewers"`
	VideoHeight       int        `json:"video_height"`
	AverageFps        float64    `json:"average_fps"`
	Delay             int        `json:"delay"`
	CreatedAt         time.Time  `json:"created_at"`
	IsPlaylist        bool       `json:"is_playlist"`
	StreamType        string     `json:"stream_type"`
	Preview           *Preview    `json:"preview"`
	Channel           *TopChannel `json:"channel"`
}
//TopChannel contains fields return call TopStream
type TopChannel struct {
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
	DisplayName                  string    `json:"display_name"`
	Game                         string    `json:"game"`
	Language                     string    `json:"language"`
	ID                           int       `json:"_id"`
	Name                         string    `json:"name"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	Partner                      bool      `json:"partner"`
	Logo                         string    `json:"logo"`
	VideoBanner                  string    `json:"video_banner"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	URL                          string    `json:"url"`
	Views                        int       `json:"views"`
	Followers                    int       `json:"followers"`
	BroadcasterType              string    `json:"broadcaster_type"`
	Description                  string    `json:"description"`
}

/*
type Channel struct {
	Mature                       bool      `json:"mature"`
	Partner                      bool      `json:"partner"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
	DisplayName                  string    `json:"display_name"`
	Game                         string    `json:"game"`
	Language                     string    `json:"language"`
	ID                           int       `json:"_id"`
	Name                         string    `json:"name"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	Delay                        int       `json:"delay"`
	Logo                         string    `json:"logo"`
	Banner                       string    `json:"banner"`
	VideoBanner                  string    `json:"video_banner"`
	Background                   string    `json:"background"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	URL                          string    `json:"url"`
	Views                        int       `json:"views"`
	Followers                    int       `json:"followers"`
	Links                        struct {
		Self          string `json:"self"`
		Follows       string `json:"follows"`
		Commercial    string `json:"commercial"`
		StreamKey     string `json:"stream_key"`
		Chat          string `json:"chat"`
		Features      string `json:"features"`
		Subscriptions string `json:"subscriptions"`
		Editors       string `json:"editors"`
		Teams         string `json:"teams"`
		Videos        string `json:"videos"`
	} `json:"_links"`
}

type NetError struct {
	StatusCode int
	Message    string
}

type UserGroup struct {
	Total int    `json:"_total"`
	Users []User `json:"users"`
}

type User struct {
	ID          string `json:"_id"`
	Bio         string `json:"bio"`
	DateCreated string `json:"created_at"`
	DisplayName string `json:"display_name"`
	Logo        string `json:"logo"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	DateUpdated string `json:"updated_at"`
}


*/
