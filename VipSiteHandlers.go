package main

import (
	"encoding/json"
	"flavorings/domain"
	"net/http"
)

type VipSiteHandlers struct {
}

func (h *VipSiteHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var d *domain.User = new(domain.User)
	mylog(d)

	json.NewEncoder(w).Encode(map[string]bool{"ok": false})
}
