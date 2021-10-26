package article

import "mixindev/model"

// check article felid error?
func ValidateCreateArticle(userId, categoryId, tagId int) (valid bool) {
	_, err := model.GetUserById(userId)
	if err != nil {
		return false
	}

	_, err = model.GetCategoryById(categoryId)
	if err != nil {
		return false
	}

	_, err = model.GetTagById(tagId)

	
	return err == nil
}
