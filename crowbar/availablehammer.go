package main

import crowbar "github.com/VictorLowther/crowbar-api"

func init() {

	maker := func() crowbar.Crudder { return &crowbar.AvailableHammer{} }
	singularName := "availablehammer"
	app.AddCommand(makeCommandTree(singularName, maker))
}
