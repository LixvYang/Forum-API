package article

import "mixindev/model"

// check article felid error?
func ValidateCreateArticle(userId, categoryId, tagId int) (valid bool) {
	var user  *model.User

	_, err := user.GetUserById(userId)
	if err != nil {
		return false
	}
	var category *model.Category

	_, err = category.GetCategoryById(categoryId)
	if err != nil {
		return false
	}

	var tag *model.Tag
	_, err = tag.GetTagById(tagId)

	
	return err == nil
}
