package main

import ap "adapterdemo/adapterpattern"

func main() {
	var t ap.Target
	t = ap.NewAdapter()
	t.Request()
}
