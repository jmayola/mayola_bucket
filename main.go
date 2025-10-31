package main

import (
	"github.com/jmayola/mayola_bucket/api"
	"github.com/jmayola/mayola_bucket/utils"
)

func main() {
	utils.GetEnv()

	api.StartApi()
}
