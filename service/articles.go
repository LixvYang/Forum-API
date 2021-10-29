package service

import (
	"mixindev/model"
)

func ListArticles(offset, limit int) ([]model.ArticleInfo, int64, error) {
	var article *model.Article
	var category *model.Category
	articles, count, err := article.ListArticles(offset, limit)
	if err != nil {
		return nil, count, err
	}

	var articleList []model.ArticleInfo
	for _, article := range articles {
		user, uErr := model.GetUserById(article.UserId)
		category, cErr := category.GetCategoryById(article.CategoryId)
		tag, tErr := model.GetTagById(article.TagId)
		if uErr != nil || cErr != nil || tErr != nil {
			return nil, count, err
		}
		articleInfo := model.ArticleInfo{
			Id:           int(article.ID),
			Title:        article.Title,
			Content:      article.Content,
			CategoryId:   article.CategoryId,
			CategoryName: category.CategoryName,
			TagId:        article.TagId,
			TagName:      tag.TagName,
			UserId:       article.UserId,
			UserName:     user.Username,
			Avatar:       user.Avatar,
			CreatedAt:    article.CreatedAt,
			UpdatedAt:    article.UpdatedAt,
		}
		articleList = append(articleList, articleInfo)
	}
	return articleList, count, nil
}
