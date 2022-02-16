package data

import "embed"

// IgnitionData contains the source data for building the ignition file
//go:embed ignition/*
var IgnitionData embed.FS
