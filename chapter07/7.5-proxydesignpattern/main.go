package main

import p "proxydesignpattern/proxydp"

func main() {
	proxy := p.NewProxy()
	proxy.Request()

}
