package main

import "app/life/life"

func main() {
	life.NewLife(1000, 1000, 2).Start()
}
