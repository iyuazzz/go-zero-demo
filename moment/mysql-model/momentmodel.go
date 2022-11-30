package mysql_model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MomentModel = (*customMomentModel)(nil)

type (
	// MomentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMomentModel.
	MomentModel interface {
		momentModel
	}

	customMomentModel struct {
		*defaultMomentModel
	}
)

// NewMomentModel returns a model for the database table.
func NewMomentModel(conn sqlx.SqlConn) MomentModel {
	return &customMomentModel{
		defaultMomentModel: newMomentModel(conn),
	}
}
