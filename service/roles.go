package service

import "mixindev/model"

func ListRoles(offset, limit int) ([]model.RoleInfo, int64, error) {
	var infos []model.RoleInfo
	var r model.Role

	list, count, err := r.ListRole(offset, limit)
	if err != nil {
		return nil, count, err
	}
	for _, l := range list {
		infos = append(infos, model.RoleInfo{
			Id:   uint64(l.ID),
			Name: l.Name,
		})
	}
	return infos, count, nil
}
