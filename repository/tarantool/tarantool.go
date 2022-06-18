package tarantool

import (
	"github.com/tarantool/go-tarantool"
	_ "github.com/tarantool/go-tarantool/uuid"

	"github.com/charCharacter/history/relations/models"
)

// Repo реализует api.Model.
type Repo struct {
	c *tarantool.Connection
}

// New инициализирует экземпляр Repo.
func New(addr, user, pass string) (Repo, error) {
	c, err := tarantool.Connect(addr, tarantool.Opts{
		User: user,
		Pass: pass,
	})
	if err != nil {
		return Repo{}, err
	}
	return Repo{c: c}, nil
}

func (m Repo) ArticleCreate(a *models.Article) (*models.Article, error) {
	_, err := m.c.Call("article_create", []interface{}{(*article)(a)})
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (m Repo) RelationsCreate(r *models.ArticleKeywords) (*models.ArticleKeywords, error) {
	_, err := m.c.Call("article_calc_relation", []interface{}{(*articleKeywords)(r)})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (m Repo) ArticleGet(uid string) (*models.Article, []*models.Keywords, error) {
	var article *article
	var keywords []*keywords
	err := m.c.CallTyped("article_get", []interface{}{uid}, &article)
	if err != nil {
		return nil, nil, err
	}
	result := (*models.Article)(article)

	err = m.c.CallTyped("keywords_get", []interface{}{uid}, &keywords)

	if err != nil {
		return nil, nil, err
	}

	result2 := make([]*models.Keywords, len(keywords))
	for i, e := range keywords {
		result2[i] = (*models.Keywords)(e)
	}

	return result, result2, nil
}

func (m Repo) ArticleList() ([]*models.Article, error) {
	var article []*article
	err := m.c.CallTyped("articles_all", []interface{}{}, &article)
	if err != nil {
		return nil, err
	}

	r := make([]*models.Article, len(article))
	for i, e := range article {
		r[i] = (*models.Article)(e)
	}

	return r, nil
}

func (m Repo) RelationsList() ([]*models.Relation, error) {
	var relations []*relation
	err := m.c.CallTyped("relations_all", []interface{}{}, &relations)
	if err != nil {
		return nil, err
	}

	r := make([]*models.Relation, len(relations))
	for i, e := range relations {
		r[i] = (*models.Relation)(e)
	}

	return r, nil
}
