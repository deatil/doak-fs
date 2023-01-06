package discard

import (
    "github.com/deatil/doak-fs/pkg/log"
)

// Default handler.
var Default = New()

// Handler implementation.
type Handler struct{}

// New handler.
func New() *Handler {
    return &Handler{}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
    return nil
}
