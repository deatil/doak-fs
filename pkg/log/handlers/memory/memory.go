package memory

import (
    "sync"

    "github.com/deatil/doak-fs/pkg/log"
)

// Handler implementation.
type Handler struct {
    mu      sync.Mutex
    Entries []*log.Entry
}

// New handler.
func New() *Handler {
    return &Handler{}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
    h.mu.Lock()
    defer h.mu.Unlock()
    h.Entries = append(h.Entries, e)
    return nil
}
