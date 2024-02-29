package state

import (
	"urbskali/ssserver/models"

	"github.com/jackc/pgx/v5"
)

var Config models.Config
var DB *pgx.Conn
