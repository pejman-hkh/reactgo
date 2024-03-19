package view

import (
	"github.com/pejman-hkh/gdp/gox"
)

type Component struct {}
var react gox.Gox = gox.Gox{Component{}}

func (r Component) Link(props map[string]string, childrens string) string {
	return react.Run("a", map[string]string{`href` :props["to"]}, []string{``,childrens})
}

func (r Component) SideNav(props map[string]string, childrens string) string {
	return react.Run("nav", map[string]string{}, []string{``,childrens})
}

func (r Component) Content(props map[string]string, childrens string) string {
	return react.Run("main", map[string]string{}, []string{``,childrens})
}

func (r Component) Header(props map[string]string, childrens string) string {
	return react.Run("f", map[string]string{}, []string{`
	`, react.Run("header", map[string]string{}, []string{`
	`, react.Run("nav", map[string]string{}, []string{`
		`, react.Run("ul", map[string]string{}, []string{`
			`, react.Run("li", map[string]string{}, []string{`
			`, react.Run("a", map[string]string{`href` :`/`}, []string{`Home`}), `
			`}), `
		`}), `
	`}), `	`,childrens}), `
	`, react.Run("h1", map[string]string{}, []string{``,props["title"]}), `
	`})
}

func (r Component) Footer(props map[string]string, childrens string) string {
	return react.Run("footer", map[string]string{}, []string{`
	`, react.Run("Link", map[string]string{`to` :`https://www.github.com/pejman-hkh/gdp`}, []string{`https://www.github.com/pejman-hkh/gdp`}), `	`,childrens})
}


func (r Component) Layout(props map[string]string, childrens string) string {
	return react.Run("f", map[string]string{}, []string{`
	`, react.Run("html", map[string]string{}, []string{`
	`, react.Run("head", map[string]string{}, []string{`
	`}), `
	`, react.Run("Header", map[string]string{`title` :`test`}, []string{`test 
		`, react.Run("Link", map[string]string{`to` :`/about`}, []string{`About`}), `
		`, react.Run("Link", map[string]string{`to` :`/contact`}, []string{`About`}), `
	`}), `
	`, react.Run("SideNav", map[string]string{}, []string{`
		`, react.Run("li", map[string]string{}, []string{react.Run("a", map[string]string{`href` :`/contact`}, []string{`Contact`})}), `
	`}), `
	`, react.Run("Content", map[string]string{}, []string{`		`,childrens}), `
	`, react.Run("Footer", map[string]string{`title` :`test`}, []string{`test`}), `
	`}), `
	`})
}
