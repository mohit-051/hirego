package handlers

import "github.com/mohit-051/hirego/config"

var Base *base

type base struct {
	conf *config.Config
}

func (b *base) Initialize(conf *config.Config) {
	Base = &base{
		conf: conf,
	}
}
