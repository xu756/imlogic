package xerr

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func ServerErrorHandler(ctx context.Context, err error) error {
	if errors.Is(err, kerrors.ErrBiz) {
		err = errors.Unwrap(err)
	}
	return err
}

func ClientErrorHandler(ctx context.Context, err error) error {
	return err
}
