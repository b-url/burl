package log

import (
	"context"
	"log/slog"
)

type contextKey int

const (
	RequestIDContextKey contextKey = 0
)

// Handler is a struct that embeds slog.Handler to provide additional
// functionality or customization for handling logs in the application.
// This Handler enables request ID functionality by adding a request_id
// to the log records if available in the context.
type Handler struct {
	slog.Handler
}

func NewLogHandler(h slog.Handler) *Handler {
	return &Handler{Handler: h}
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(RequestIDContextKey).(string); ok {
		r.Add("request_id", slog.StringValue(traceID))
	}

	return h.Handler.Handle(ctx, r)
}

func (h *Handler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h.clone()
}

func (h *Handler) clone() *Handler {
	clone := *h
	return &clone
}
