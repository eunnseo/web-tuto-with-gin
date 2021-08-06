package model

import (
	"errors"
	"time"
)

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	// article repository
	articleList = []Article{
		Article{
			ID:        0,
			Title:     "How to Fight Fraud with Artificial Intelligence and Intelligent Analytics",
			Content:   "Could artificial intelligence have been used to prevent the high-profile Target breach? The concept is not so far-fetched. Organizations such as Mastercard and RBS WorldPay have long relied on artificial intelligence to detect fraudulent transaction patterns and prevent card.",
			CreatedAt: time.Now(),
		},
		Article{
			ID:        1,
			Title:     "Big Payment Data: Leveraging Transactional Data to Ensure an Enterprise Approach to Risk Management",
			Content:   "60% of organizations were exposed to actual or attempted fraud loss last year. As fraud and risk increases year over year, the amount of data being collected increases as well.",
			CreatedAt: time.Now(),
		},
	}
	id int
)

func init() {
	// init article's unique id
	id = len(articleList)
}

// 메모리에 저장된 모든 Article들을 리턴
func GetAllArticles() []Article {
	return articleList
}

// 메모리에 저장된 Article 리스트에서 Id에 해당하는 Article을 검색한 뒤에 해당 값을 리턴
func GetArticleByID(id int) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

// 새로운 Article 인스턴스를 생성한 후에 articleList에 추가
func CreateNewArticle(title, content string) (*Article, error) {
	a := Article{
		ID:    id,
		Title: title, Content: content,
		CreatedAt: time.Now(),
	}
	id++

	articleList = append(articleList, a)

	return &a, nil
}

// articleList에서 id에 해당하는 Article을 검색한 뒤에 해당 글을 리스트에서 삭제
func DeleteArticleByID(id int) error {
	for i, a := range articleList {
		if a.ID == id {
			articleList = append(articleList[:i], articleList[i+1:]...)
			return nil
		}
	}
	return errors.New("Article not found")
}
