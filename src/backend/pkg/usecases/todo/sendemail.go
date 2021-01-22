package todo

import (
	"context"
	"fmt"
	"gosvelteHex/src/backend/pkg/events"
)



func SendEmail(ctx context.Context, msg events.Msg) error {

	fmt.Println("send Email")

	return nil
}
