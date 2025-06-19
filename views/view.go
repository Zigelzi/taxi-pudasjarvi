package views

import (
	"github.com/Zigelzi/taxi-pudasjarvi/components"
	"github.com/a-h/templ"
)

func Get() map[string]templ.Component {
	views := map[string]templ.Component{
		"/": components.Index(),
	}
	return views
}
