package multi

import (
    "github.com/deatil/doak-fs/pkg/log"
)

// Handler implementation.
type Handler struct {
    Handlers []log.Handler
}

// New handler.
func New(h ...log.Handler) *Handler {
    return &Handler{
        Handlers: h,
    }
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
    for _, handler := range h.Handlers {
        // TODO(tj): maybe just write to stderr here, definitely not ideal
        // to miss out logging to a more critical handler if something
        // goes wrong
        if err := handler.HandleLog(e); err != nil {
            return err
        }
    }

    return nil
}
