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
	users = append(users, domain.User{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 5, Name: "test"}})
	users = append(users, domain.User{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 6, Name: "guest"}})
	json.NewEncoder(w).Encode(users)
}

func (h *VipSiteHandlers) GetRoles(w http.ResponseWriter, r *http.Request) {
	users := []domain.Role{}
	users = append(users, domain.Role{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 5, Name: "test"}})
	users = append(users, domain.Role{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 6, Name: "guest"}})
	json.NewEncoder(w).Encode(users)
}

func (h *VipSiteHandlers) GetPermissions(w http.ResponseWriter, r *http.Request) {
	users := []domain.Permission{}
	users = append(users, domain.Permission{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 5, Name: "permissions"}})
	users = append(users, domain.Permission{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 6, Name: "guest"}})
	json.NewEncoder(w).Encode(users)
}

func (h *VipSiteHandlers) GetGames(w http.ResponseWriter, r *http.Request) {
	users := []domain.Game{}
	users = append(users, domain.Game{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 5, Name: "games"}})
	users = append(users, domain.Game{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 6, Name: "guest"}})
	json.NewEncoder(w).Encode(users)
}

func (h *VipSiteHandlers) GetChannels(w http.ResponseWriter, r *http.Request) {
	users := []domain.Channel{}
	users = append(users, domain.Channel{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 5, Name: "channel"}})
	users = append(users, domain.Channel{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 6, Name: "guest"}})
	json.NewEncoder(w).Encode(users)
}

func (h *VipSiteHandlers) GetRules(w http.ResponseWriter, r *http.Request) {
	users := []domain.Rule{}
	users = append(users, domain.Rule{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 5, Name: "rule"}})
	users = append(users, domain.Rule{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 6, Name: "guest"}})
	json.NewEncoder(w).Encode(users)
}

func (h *VipSiteHandlers) GetWishTickets(w http.ResponseWriter, r *http.Request) {
	users := []domain.WishTicket{}
	users = append(users, domain.WishTicket{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 5, Name: "wishticket"}})
	users = append(users, domain.WishTicket{IdentifiedNamedObject: domain.IdentifiedNamedObject{Id: 6, Name: "guest"}})
	json.NewEncoder(w).Encode(users)
}
