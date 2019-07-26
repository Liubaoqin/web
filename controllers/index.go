package controllers

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	this.TplName = "index.html"
}

func (this *IndexController) About() {
	this.TplName = "about.html"
}

func (this *IndexController) Message() {
	this.TplName = "message.html"
}

func (this *IndexController) Details() {
	this.TplName = "details.html"
}

func (this *IndexController) Comment() {
	this.TplName = "comment.html"
}

func (this *IndexController) Error() {
	this.TplName = "404.html"
}
