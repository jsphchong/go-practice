package main

import "time"

const TimeFormat = "03:04 PM on January 02, 2006"

type Post struct {
	Id            string    `json:"id"`
	Time          time.Time `json:"date"`
	FormattedTime string    `json:"formattedTime"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
}

type PostMap map[string]*Post

type ById []Post

func (xt ById) Len() int           { return len(xt) }
func (xt ById) Swap(i, j int)      { xt[i], xt[j] = xt[j], xt[i] }
func (xt ById) Less(i, j int) bool { return xt[i].Id < xt[j].Id }
