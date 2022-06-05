package main

import (
	"encoding/json"
	"net/http"
)

type VipSiteHandlers struct {
}

func (h *VipSiteHandlers) Login(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": false})
}
