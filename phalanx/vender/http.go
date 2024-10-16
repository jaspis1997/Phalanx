package vender

import (
	"crypto/rand"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type (
	HttpContext = *gin.Context
)

type HttpEngine = *gin.Engine

const (
	HttpContentType = "Content-Type"
)

const (
	ContentType_ImageJpeg   = "image/jpg"
	ContentType_OctetStream = "application/octet-stream"
)

type (
	CookieStore = cookie.Store
)

func NewEngine() HttpEngine {
	e := gin.New()
	e.Use(gin.Recovery())
	return e
}

func NewCookieStore() CookieStore {
	buf := make([]byte, 64)
	rand.Read(buf)
	return cookie.NewStore(buf)
}

func Sessions(store CookieStore) gin.HandlerFunc {
	return sessions.Sessions("session", store)
}

type SessionValues map[string]func() any

func (sv SessionValues) CheckSession(ctx HttpContext) {
	session := sessions.Default(ctx)
	for k, init := range sv {
		v := session.Get(k)
		if v == nil {
			v = init()
			session.Set(k, v)
		}
		ctx.Set(k, v)
	}
}
