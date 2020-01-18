package main

import "myProject/gtool/smap"

func main() {
	smap := smap.NewSmap()
	smap.Add("allen", 1)
	smap.Remove("boy")
}