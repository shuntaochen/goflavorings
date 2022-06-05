package main

import (
	"encoding/json"
	"net/http"
	"flavorings/domain"
)

type VipSiteHandlers struct {
}

func (h *VipSiteHandlers) Login(w http.ResponseWriter, r *http.Request) {
	domain.User d:=New()
	
	json.NewEncoder(w).Encode(map[string]bool{"ok": false})
}
