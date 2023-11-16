package biz

type Meta struct {
	Total int64
	CurrentPage int64
	PerPage int64
	HasNext int64
}

type Data struct {
	Meta Meta
	Items interface{}
}
