package user_models

import (
	"github.com/jackc/pgx/v5"
)

type Entity interface {
  Scan(row pgx.Row) error
  ScanRows(rows pgx.Rows) error
}
