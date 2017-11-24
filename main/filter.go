package main

import "net/url"

type Filter interface {
	ForLearned()
	ForSortDate()
	ForType()
	ForGroup()
	ToString()
}

type LeoFilter struct {
	url_builder url.URL
}

func (leo *LeoFilter)ForLearned(){
	q := leo.url_builder.Query()
	q.Add("filter", "learned")
}

