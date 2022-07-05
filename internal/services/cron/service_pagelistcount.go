package cron

import (
	"github.com/recallfuture/blog-api/internal/pkg/core"
	"github.com/recallfuture/blog-api/internal/repository/mysql"
	"github.com/recallfuture/blog-api/internal/repository/mysql/cron_task"
)

func (s *service) PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error) {
	qb := cron_task.NewQueryBuilder()

	if searchData.Name != "" {
		qb.WhereName(mysql.EqualPredicate, searchData.Name)
	}

	if searchData.Protocol != 0 {
		qb.WhereProtocol(mysql.EqualPredicate, searchData.Protocol)
	}

	if searchData.IsUsed != 0 {
		qb.WhereIsUsed(mysql.EqualPredicate, searchData.IsUsed)
	}

	total, err = qb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	return
}
