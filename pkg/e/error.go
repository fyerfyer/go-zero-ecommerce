package e

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(code codes.Code, err error, msg, context string) error {
	if err != nil {
		return status.Error(code, fmt.Sprintf("[%s]: msg%v", context, err))
	}
	return nil
}
