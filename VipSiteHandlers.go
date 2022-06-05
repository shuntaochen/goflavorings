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
	d.Code = "hello world"
	// json.NewEncoder(w).Encode(map[string]bool{"ok": false})
	json.NewEncoder(w).Encode(d)
}
