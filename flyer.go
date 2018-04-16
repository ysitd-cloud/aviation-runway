package tower

import (
	"context"
	"net/http"
)

type Flyer interface {
	Initial(ctx context.Context) error
	GetHandler() http.Handler
}

