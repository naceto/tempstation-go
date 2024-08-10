package web

import "embed"

// Content -
//
//go:embed index.html sensors/* generic/* common/*
var Content embed.FS
