package service

import "mixindev/model"

func ListTags(offset, limit int) ([]model.TagInfo, int64, error) {
	var c model.Tag
	tags, count, err := c.ListTags(offset, limit)
	if err != nil {
		return nil, count, err
	}

	var infos []model.TagInfo

	for _, tag := range tags {
		tagInfo := model.TagInfo{
			Id:      int(tag.ID),
			TagName: tag.TagName,
		}
		infos = append(infos, tagInfo)
	}

	return infos, count, nil
}
