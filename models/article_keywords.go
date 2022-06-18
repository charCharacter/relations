package models

type ArticleKeywords struct {
	ArticleUID string   `json:"article_uid"`
	Keywords   []string `json:"keywords"`
}
type Keywords struct {
	Uid        string `json:"uid"`
	ArticleUid string `json:"article_uid"`
	Keyword    string `json:"keywords"`
}
