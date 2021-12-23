package model

import "git.zx-tech.net/pengfeng/facebook/model"

type Businesses struct {
	Data   []Business   `json:"data"`
	Paging model.Paging `json:"paging"`
}

type Business struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}
type BusinessUsers struct {
	Data   []BusinessUser `json:"data"`
	Paging model.Paging   `json:"paging"`
}
type BusinessUser struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Role              string   `json:"role"`
	Title             string   `json:"title"`
	Email             string   `json:"email"`
	FinancePermission string   `json:"finance_permission"`
	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	PendingEmail      string   `json:"pending_email"`
	IpPermission      string   `json:"ip_permission"`
	Business          Business `json:"business"`
}
type BusinessAdAccount struct {
	Data   []AdCoount   `json:"data"`
	Paging model.Paging `json:"paging"`
}

type BusinessAssignedAdAccountMeta struct {
	Data   []AdCoount   `json:"data"`
	Paging model.Paging `json:"paging"`
}

//BusinessAdAccountMeta ...
type BusinessAdAccountMeta struct {
	Data   []AdCoount   `json:"data"`
	Paging model.Paging `json:"paging"`
}
