package main

type Post struct {
	Id      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostMap map[string]*Post

type ById []Post

func (xt ById) Len() int {return len(xt)}
func (xt ById) Swap(i, j int) {xt[i], xt[j] = xt[j], xt[i]}
func (xt ById) Less(i, j int) bool {return xt[i].Id < xt[j].Id}