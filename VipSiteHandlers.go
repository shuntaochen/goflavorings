package main

import (
	"encoding/json"
	"flavorings/domain"
	"net/http"
	"time"
)

type VipSiteHandlers struct {
}

func (h *VipSiteHandlers) Login(w http.ResponseWriter, r *http.Request) {
	var d *domain.User = new(domain.User)
	mylog(d)
	d.Code = "hello world"
	d.CreatedOn = time.Now().UnixMilli()
	// json.NewEncoder(w).Encode(map[string]bool{"ok": false})
	json.NewEncoder(w).Encode(d)
}

func (h *VipSiteHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []domain.User{}
	users = append(users, domain.User{IdentifiedNamedObject: &domain.IdentifiedNamedObject{Id: 5, Name: "test"}})
	users = append(users, domain.User{IdentifiedNamedObject: &domain.IdentifiedNamedObject{Id: 6, Name: "guest"}})
	json.NewEncoder(w).Encode(users)
}
