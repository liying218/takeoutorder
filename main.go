package main

import "takeoutorder/router"

func main() {
	r := router.Router()

	r.Run(":8081")
}
