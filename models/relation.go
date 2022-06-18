package models

type Relation struct {
	RelationUID string  `json:"relation_uid"`
	ArticleUIDA string  `json:"article_uid_a"`
	ArticleUIDB string  `json:"article_uid_b"`
	Coefficient float64 `json:"coefficient"`
}
