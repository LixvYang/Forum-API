package service

import "mixindev/model"

func ListCategories(offset, limit int) ([]model.CategoryInfo, int64, error) {
	var c model.Category
	categories, count, err := c.ListCategories(offset, limit)
	if err != nil {
		return nil, count, err
	}

	var infos []model.CategoryInfo

	for _, category := range categories {
		categoryInfo := model.CategoryInfo{
			Id:           uint64(category.ID),
			CategoryName: category.CategoryName,
		}
		infos = append(infos, categoryInfo)
	}

	return infos, count, nil
}
