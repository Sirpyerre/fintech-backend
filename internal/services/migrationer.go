package services

import (
	"context"
	"io"
)

type Migrationer interface {
	Migrate(ctx context.Context, cvsFile io.Reader) error
}
