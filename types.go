package goinsta

import (
	"fmt"
	"net/http"
	"strconv"
)

type ConfigFile struct {
	User      string         `json:"username"`
	DeviceID  string         `json:"device_id"`
	UUID      string         `json:"uuid"`
	RankToken string         `json:"rank_token"`
	Token     string         `json:"token"`
	PhoneID   string         `json:"phone_id"`
	Cookies   []*http.Cookie `json:"cookies"`
}

// School is void structure (yet).
type School struct {
}

// PicURLInfo repre
type PicURLInfo struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type instaError struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	ErrorType string `json:"error_type"`
}

func instaToErr(ierr instaError) error {
	return fmt.Errorf("%s: %s (%s)", ierr.Status, ierr.Message, ierr.ErrorType)
}

type Nametag struct {
	Mode          int    `json:"mode"`
	Gradient      int    `json:"gradient"`
	Emoji         string `json:"emoji"`
	SelfieSticker int    `json:"selfie_sticker"`
}

type friendResp struct {
	Status     string     `json:"status"`
	Friendship Friendship `json:"friendship_status"`
}

type Location struct {
	Pk               int     `json:"pk"`
	Name             string  `json:"name"`
	Address          string  `json:"address"`
	City             string  `json:"city"`
	ShortName        string  `json:"short_name"`
	Lng              float64 `json:"lng"`
	Lat              float64 `json:"lat"`
	ExternalSource   string  `json:"external_source"`
	FacebookPlacesID int64   `json:"facebook_places_id"`
}

type SuggestedUsers struct {
	Type        int `json:"type"`
	Suggestions []struct {
		User            User          `json:"user"`
		Algorithm       string        `json:"algorithm"`
		SocialContext   string        `json:"social_context"`
		Icon            string        `json:"icon"`
		Caption         string        `json:"caption"`
		MediaIds        []interface{} `json:"media_ids"`
		ThumbnailUrls   []interface{} `json:"thumbnail_urls"`
		LargeUrls       []interface{} `json:"large_urls"`
		MediaInfos      []interface{} `json:"media_infos"`
		Value           float64       `json:"value"`
		IsNewSuggestion bool          `json:"is_new_suggestion"`
	} `json:"suggestions"`
	LandingSiteType  string `json:"landing_site_type"`
	Title            string `json:"title"`
	ViewAllText      string `json:"view_all_text"`
	LandingSiteTitle string `json:"landing_site_title"`
	NetegoType       string `json:"netego_type"`
	UpsellFbPos      string `json:"upsell_fb_pos"`
	AutoDvance       string `json:"auto_dvance"`
	ID               string `json:"id"`
	TrackingToken    string `json:"tracking_token"`
}

type Friendship struct {
	IncomingRequest bool `json:"incoming_request"`
	FollowedBy      bool `json:"followed_by"`
	OutgoingRequest bool `json:"outgoing_request"`
	Following       bool `json:"following"`
	Blocking        bool `json:"blocking"`
	IsPrivate       bool `json:"is_private"`
	Muting          bool `json:"muting"`
	IsMutingReel    bool `json:"is_muting_reel"`
}

type SavedMedia struct {
	Items []struct {
		Media Item `json:"media"`
	} `json:"items"`
	NumResults          int    `json:"num_results"`
	MoreAvailable       bool   `json:"more_available"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Status              string `json:"status"`
}

// Images are different quality images
type Images struct {
	Versions []Candidate `json:"candidates"`
}

type Candidate struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

type Comment struct {
	ID                             int64     `json:"pk"`
	idstr                          string    `json:"-"`
	Text                           string    `json:"text"`
	Type                           int       `json:"type"`
	User                           User      `json:"user"`
	UserID                         int64     `json:"user_id"`
	BitFlags                       int       `json:"bit_flags"`
	ChildCommentCount              int       `json:"child_comment_count"`
	CommentIndex                   int       `json:"comment_index"`
	CommentLikeCount               int       `json:"comment_like_count"`
	ContentType                    string    `json:"content_type"`
	CreatedAt                      int       `json:"created_at"`
	CreatedAtUtc                   int       `json:"created_at_utc"`
	DidReportAsSpam                bool      `json:"did_report_as_spam"`
	HasLikedComment                bool      `json:"has_liked_comment"`
	InlineComposerDisplayCondition string    `json:"inline_composer_display_condition"`
	OtherPreviewUsers              []User    `json:"other_preview_users"`
	PreviewChildComments           []Comment `json:"preview_child_comments"`
	Status                         string    `json:"status"`
}

func (c Comment) getid() string {
	switch {
	case c.ID == 0:
		return c.idstr
	case c.idstr == "":
		return strconv.FormatInt(c.ID, 10)
	}
	return ""
}

type Tag struct {
	In []struct {
		User                  User        `json:"user"`
		Position              []float64   `json:"position"`
		StartTimeInVideoInSec interface{} `json:"start_time_in_video_in_sec"`
		DurationInVideoInSec  interface{} `json:"duration_in_video_in_sec"`
	} `json:"in"`
}

// Caption is media caption
type Caption struct {
	ID              int64  `json:"pk"`
	UserID          int    `json:"user_id"`
	Text            string `json:"text"`
	Type            int    `json:"type"`
	CreatedAt       int    `json:"created_at"`
	CreatedAtUtc    int    `json:"created_at_utc"`
	ContentType     string `json:"content_type"`
	Status          string `json:"status"`
	BitFlags        int    `json:"bit_flags"`
	User            User   `json:"user"`
	DidReportAsSpam bool   `json:"did_report_as_spam"`
	MediaID         int64  `json:"media_id"`
	HasTranslation  bool   `json:"has_translation"`
}

type Mentions struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        int     `json:"z"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation float64 `json:"rotation"`
	IsPinned int     `json:"is_pinned"`
	User     User    `json:"user"`
}

// Video are different quality videos
type Video struct {
	Type   int    `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
	ID     string `json:"id"`
}

type timeStoryResp struct {
	Status string       `json:"status"`
	Media  []StoryMedia `json:"tray"`
}

// Tray is a set of story media received from timeline calls.
type Tray struct {
	Stories []StoryMedia `json:"tray"`
	Lives   struct {
		LiveItems []LiveItems `json:"post_live_items"`
	} `json:"post_live"`
	StoryRankingToken    string      `json:"story_ranking_token"`
	Broadcasts           []Broadcast `json:"broadcasts"`
	FaceFilterNuxVersion int         `json:"face_filter_nux_version"`
	HasNewNuxStory       bool        `json:"has_new_nux_story"`
	Status               string      `json:"status"`
}

func (tray *Tray) set(inst *Instagram, url string) {
	for i := range tray.Stories {
		tray.Stories[i].inst = inst
		tray.Stories[i].endpoint = url
		tray.Stories[i].setValues()
	}
	for i := range tray.Lives.LiveItems {
		tray.Lives.LiveItems[i].User.inst = inst
		for j := range tray.Lives.LiveItems[i].Broadcasts {
			tray.Lives.LiveItems[i].Broadcasts[j].BroadcastOwner.inst = inst
		}
	}
	for i := range tray.Broadcasts {
		tray.Broadcasts[i].BroadcastOwner.inst = inst
	}
}

// LiveItems are Live media items
type LiveItems struct {
	ID                  string      `json:"pk"`
	User                User        `json:"user"`
	Broadcasts          []Broadcast `json:"broadcasts"`
	LastSeenBroadcastTs int         `json:"last_seen_broadcast_ts"`
	RankedPosition      int         `json:"ranked_position"`
	SeenRankedPosition  int         `json:"seen_ranked_position"`
	Muted               bool        `json:"muted"`
	CanReply            bool        `json:"can_reply"`
	CanReshare          bool        `json:"can_reshare"`
}

// Broadcast is live videos.
type Broadcast struct {
	ID                   int64  `json:"id"`
	BroadcastStatus      string `json:"broadcast_status"`
	DashManifest         string `json:"dash_manifest"`
	ExpireAt             int    `json:"expire_at"`
	EncodingTag          string `json:"encoding_tag"`
	InternalOnly         bool   `json:"internal_only"`
	NumberOfQualities    int    `json:"number_of_qualities"`
	CoverFrameURL        string `json:"cover_frame_url"`
	BroadcastOwner       User   `json:"broadcast_owner"`
	PublishedTime        int    `json:"published_time"`
	MediaID              string `json:"media_id"`
	BroadcastMessage     string `json:"broadcast_message"`
	OrganicTrackingToken string `json:"organic_tracking_token"`
}

type BlockedUser struct {
	UserID        int64  `json:"user_id"`
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	ProfilePicURL string `json:"profile_pic_url"`
	BlockAt       int    `json:"block_at"`
}

// Unblock unblocks blocked user.
func (b *BlockedUser) Unblock() error {
	u := User{ID: b.UserID}
	return u.Unblock()
}

type blockedListResp struct {
	BlockedList []BlockedUser `json:"blocked_list"`
	PageSize    int           `json:"page_size"`
	Status      string        `json:"status"`
}

type InboxItem struct {
	ItemID     string `json:"item_id"`
	UserID     int64  `json:"user_id"`
	Timestamp  int64  `json:"timestamp"`
	ItemType   string `json:"item_type"`
	RavenMedia struct {
		MediaType int `json:"media_type"`
	} `json:"raven_media"`
	ClientContext              string        `json:"client_context"`
	SeenUserIds                []interface{} `json:"seen_user_ids"`
	ReplyChainCount            int           `json:"reply_chain_count"`
	ExpiringMediaActionSummary struct {
		Type      string `json:"type"`
		Timestamp int64  `json:"timestamp"`
		Count     int    `json:"count"`
	} `json:"expiring_media_action_summary"`
	ViewMode string `json:"view_mode"`
	Like     string `json:"like"`
}

type Thread struct {
	ThreadID                  string      `json:"thread_id"`
	ThreadV2ID                int64       `json:"thread_v2_id"`
	Users                     []User      `json:"users"`
	LeftUsers                 []User      `json:"left_users"`
	Items                     []InboxItem `json:"items"`
	LastActivityAt            int64       `json:"last_activity_at"`
	Muted                     bool        `json:"muted"`
	IsPin                     bool        `json:"is_pin"`
	Named                     bool        `json:"named"`
	Canonical                 bool        `json:"canonical"`
	Pending                   bool        `json:"pending"`
	ValuedRequest             bool        `json:"valued_request"`
	ThreadType                string      `json:"thread_type"`
	ViewerID                  int64       `json:"viewer_id"`
	ThreadTitle               string      `json:"thread_title"`
	PendingScore              int64       `json:"pending_score"`
	ReshareSendCount          int         `json:"reshare_send_count"`
	ReshareReceiveCount       int         `json:"reshare_receive_count"`
	ExpiringMediaSendCount    int         `json:"expiring_media_send_count"`
	ExpiringMediaReceiveCount int         `json:"expiring_media_receive_count"`
	Inviter                   User        `json:"inviter"`
	HasOlder                  bool        `json:"has_older"`
	HasNewer                  bool        `json:"has_newer"`
	LastSeenAt                struct {
		Num7629421016 struct {
			Timestamp string `json:"timestamp"`
			ItemID    string `json:"item_id"`
		} `json:"7629421016"`
	} `json:"last_seen_at"`
	NewestCursor      string `json:"newest_cursor"`
	OldestCursor      string `json:"oldest_cursor"`
	IsSpam            bool   `json:"is_spam"`
	LastPermanentItem Item   `json:"last_permanent_item"`
}
