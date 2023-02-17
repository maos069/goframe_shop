// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe_shop/internal/dao/internal"
)

// internalArticleInfoDao is internal type for wrapping internal DAO implements.
type internalArticleInfoDao = *internal.ArticleInfoDao

// articleInfoDao is the data access object for table article_info.
// You can define custom methods on it to extend its functionality as you wish.
type articleInfoDao struct {
	internalArticleInfoDao
}

var (
	// ArticleInfo is globally public accessible object for table article_info operations.
	ArticleInfo = articleInfoDao{
		internal.NewArticleInfoDao(),
	}
)

// Fill with you ideas below.