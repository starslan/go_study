package file

type Event struct {
	ID   string      `json:"id"`
	Data URLUserItem `json:"data"`
}

type URLUserItem struct {
	URL    string `json:"url"`
	UserId uint32 `json:"user-id"`
}
