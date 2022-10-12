package service

import (
	"strconv"

	_const "github.com/c/server-monitoring/common/const"
	"github.com/c/server-monitoring/server/internal/v1/daos"
)

// MenusListView
type MenusListView struct {
	ID        int64            `json:"id"`
	ParentID  string           `json:"parent_id"`
	MenuName  string           `json:"menu_name"`
	Icons     string           `json:"icons"`
	Url       string           `json:"url"`
	FrontUrl  string           `json:"front_url"`
	Method    string           `json:"method"`
	CreatedAt string           `json:"created_at"`
	Children  []*MenusListView `json:"children"`
}

// MenuList
func MenuList() []*MenusListView {
	data := make([]*MenusListView, 0)
	menus, children := daos.MenuList()
	for _, v := range menus {
		mlv := &MenusListView{}
		mlv.ID = v.Id
		mlv.ParentID = v.ParentId
		mlv.MenuName = v.MenuName
		mlv.Icons = v.Icons
		mlv.Url = v.Url
		mlv.FrontUrl = v.FrontUrl
		mlv.Method = v.Method
		if !v.CreatedAt.IsZero() {
			mlv.CreatedAt = v.CreatedAt.Format(_const.Layout)
		}
		for _, c := range children {
			pid, _ := strconv.Atoi(c.ParentId)
			if int64(pid) == v.Id {
				ml := &MenusListView{}
				ml.ID = c.Id
				ml.ParentID = c.ParentId
				ml.MenuName = c.MenuName
				ml.Icons = c.Icons
				ml.Url = c.Url
				ml.FrontUrl = c.FrontUrl
				ml.Method = c.Method
				if !v.CreatedAt.IsZero() {
					ml.CreatedAt = c.CreatedAt.Format(_const.Layout)
				}
				mlv.Children = append(mlv.Children, ml)
			}
		}
		data = append(data, mlv)
	}
	return data
}
