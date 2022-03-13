package repository

import (
	"auth-service/app/model"
)

func FindAccessRoleByIdRole(idRole int) ([]model.Access, error) {
	db, err := openConnection()
	var accesses []model.Access
	db.Select("accesses.name").Joins("INNER JOIN access_role ON access_role.id_access = accesses.id").Where(
		"access_role.id_access = ?", idRole).Find(&accesses)
	return accesses, err
}
