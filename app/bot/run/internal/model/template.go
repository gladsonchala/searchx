package model

type TStart struct {
	ID       int64
	Username string
	Notice   string
	Chats    []string
	Version  string
}

type TSearchResults struct {
	Keyword string
	Results []*TSearchResult
	Took    int64
}

type TSearchResult struct {
	Seq        int
	ViewLink   string
	SenderName string
	SenderLink string
	ChatName   string
	Date       string
	Content    string
	GoLink     string
}

type TSearchView struct {
	MsgID      string
	ChatID     string
	ChatName   string
	SenderID   string
	SenderName string
	Date       string
	Content    string
}
