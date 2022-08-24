package server

import "light-controller/internal/types"

type Config interface {
	GetGinMode() types.GinMode
	GetGinDomain() string
	GetGinJwtCookie() string
}
