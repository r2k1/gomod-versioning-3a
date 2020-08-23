package main

import (
	gomod_versioning_1a "github.com/r2k1/gomod-versioning-1a"
	gomod_versioning_2a "github.com/r2k1/gomod-versioning-2a"
)

func main() {
	err := gomod_versioning_1a.Func()
	if err != nil {
		panic(err)
	}
	err = gomod_versioning_2a.Func2()
	if err != nil {
		panic(err)
	}
}
