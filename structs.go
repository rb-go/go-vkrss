package main

type Item struct {
	ID          int          `json:"id"`
	FromID      int          `json:"from_id"`
	OwnerID     int          `json:"owner_id"`
	Date        int          `json:"date"`
	MarkedAsAds int          `json:"marked_as_ads"`
	PostType    string       `json:"post_type"`
	Text        string       `json:"text"`
	CanEdit     int          `json:"can_edit,omitempty"`
	CreatedBy   int          `json:"created_by,omitempty"`
	CanDelete   int          `json:"can_delete"`
	CanPin      int          `json:"can_pin"`
	Attachments []Attachment `json:"attachments,omitempty"`
	PostSource  struct {
		Type     string `json:"type"`
		Platform string `json:"platform"`
	} `json:"post_source"`
	Comments struct {
		Count         int  `json:"count"`
		GroupsCanPost bool `json:"groups_can_post"`
		CanPost       int  `json:"can_post"`
	} `json:"comments"`
	Likes struct {
		Count      int `json:"count"`
		UserLikes  int `json:"user_likes"`
		CanLike    int `json:"can_like"`
		CanPublish int `json:"can_publish"`
	} `json:"likes"`
	Reposts struct {
		Count        int `json:"count"`
		UserReposted int `json:"user_reposted"`
	} `json:"reposts"`
	Views struct {
		Count int `json:"count"`
	} `json:"views"`
	CopyHistory []Item  `json:"copy_history,omitempty"`
	Groups      []Group `json:"groups"`
}

type JSONBody struct {
	Error    map[string]interface{} `json:"error"`
	Response struct {
		Count int    `json:"count"`
		Items []Item `json:"items"`
	} `json:"response"`
}

type Group struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	IsClosed   int    `json:"is_closed"`
	Type       string `json:"type"`
	IsAdmin    int    `json:"is_admin"`
	AdminLevel int    `json:"admin_level"`
	IsMember   int    `json:"is_member"`
	Photo50    string `json:"photo_50"`
	Photo100   string `json:"photo_100"`
	Photo200   string `json:"photo_200"`
}

type Attachment struct {
	Type string `json:"type"`

	Video struct {
		ID          int    `json:"id"`
		OwnerID     int    `json:"owner_id"`
		Title       string `json:"title"`
		Duration    int    `json:"duration"`
		Description string `json:"description"`
		Date        int    `json:"date"`
		Comments    int    `json:"comments"`
		Views       int    `json:"views"`
		Photo130    string `json:"photo_130"`
		Photo320    string `json:"photo_320"`
		Photo640    string `json:"photo_640"`
		Photo800    string `json:"photo_800"`
		AccessKey   string `json:"access_key"`
		Platform    string `json:"platform"`
		CanAdd      int    `json:"can_add"`
	} `json:"video"`

	Photo struct {
		ID        int    `json:"id"`
		AlbumID   int    `json:"album_id"`
		OwnerID   int    `json:"owner_id"`
		UserID    int    `json:"user_id"`
		Photo75   string `json:"photo_75"`
		Photo130  string `json:"photo_130"`
		Photo604  string `json:"photo_604"`
		Photo807  string `json:"photo_807"`
		Photo1280 string `json:"photo_1280"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Text      string `json:"text"`
		Date      int    `json:"date"`
		PostID    int    `json:"post_id"`
		AccessKey string `json:"access_key"`
	} `json:"photo"`

	Doc struct {
		ID        int    `json:"id"`
		OwnerID   int    `json:"owner_id"`
		Title     string `json:"title"`
		Size      int    `json:"size"`
		Ext       string `json:"ext"`
		URL       string `json:"url"`
		Date      int    `json:"date"`
		Type      int    `json:"type"`
		AccessKey string `json:"access_key"`
	} `json:"doc"`

	Audio struct {
		ID       int    `json:"id"`
		OwnerID  int    `json:"owner_id"`
		Artist   string `json:"artist"`
		Title    string `json:"title"`
		Duration int    `json:"duration"`
		Date     int    `json:"date"`
		URL      string `json:"url"`
		GenreID  int    `json:"genre_id"`
		IsHQ     bool   `json:"is_hq"`
	} `json:"audio"`
}
