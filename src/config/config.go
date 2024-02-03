package config

import "sync"

var Lock = sync.Mutex{}

var LANGUAGE = "en"
var DEFAULT_CONFIRMED = true
