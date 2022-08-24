package types

type GinMode int

const (
	GinDebug GinMode = iota
	GinTest
	GinRelease
)
