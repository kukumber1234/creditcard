package service

import (
	"net/http"

	dal "tea/internal/dal"
	model "tea/models"
)

func PostMenuService(putMenu model.MenuItem) {
	// logic

	dal.PostMenuDal(putMenu)
}

func GetMenuService(data []byte, w http.ResponseWriter) {
	// logic check

	dal.GetMenuDal(data, w)
}
