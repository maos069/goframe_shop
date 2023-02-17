// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe_shop/internal/dao/internal"
)

// internalRotationInfoDao is internal type for wrapping internal DAO implements.
type internalRotationInfoDao = *internal.RotationInfoDao

// rotationInfoDao is the data access object for table rotation_info.
// You can define custom methods on it to extend its functionality as you wish.
type rotationInfoDao struct {
	internalRotationInfoDao
}

var (
	// RotationInfo is globally public accessible object for table rotation_info operations.
	RotationInfo = rotationInfoDao{
		internal.NewRotationInfoDao(),
	}
)

// Fill with you ideas below.
