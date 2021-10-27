package service

import "mixindev/model"

func ListMenus(offset, limit int) ([]model.MenuInfo, int64, error) {
	var infos []model.MenuInfo
	var m model.Menu

	list, count, err := m.ListMenus(offset, limit)

	if err != nil {
		return nil, count, err
	}

	for _, l := range list {
		infos = append(infos, model.MenuInfo{
			Id:     int(l.ID),
			Name:   l.Name,
			Path:   l.Path,
			Method: l.Method,
		})
	}
	return infos, count, nil
}
