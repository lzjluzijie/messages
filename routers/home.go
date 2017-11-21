package routers

import macaron "gopkg.in/macaron.v1"

//Home function shows the main page of website.
func Home(c *macaron.Context) {
	c.HTML(200, "home")
}
