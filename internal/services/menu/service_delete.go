package menu

import (
	"github.com/recallfuture/blog-api/internal/pkg/core"
	"github.com/recallfuture/blog-api/internal/repository/mysql"
	"github.com/recallfuture/blog-api/internal/repository/mysql/menu"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": ctx.SessionUserInfo().UserName,
	}

	qb := menu.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
