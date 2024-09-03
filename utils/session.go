package utils

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Session struct {
	LastActive string `json:"lastActive" validate:"required,datetime"`
}

func (SessData Session) Map() map[interface{}]interface{} {
	in := &SessData

	var inInterface map[interface{}]interface{}
	inconv, _ := json.Marshal(in)

	json.Unmarshal(inconv, &inInterface)

	return inInterface
}

func GetSession(eContext echo.Context, name string) (*sessions.Session, error) {
	sess, err := session.Get(name, eContext)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func CreateSession(eContext echo.Context, name string, sessionData Session) (*sessions.Session, error) {
	sess, err := GetSession(eContext, name)
	if err != nil {
		return nil, err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values = sessionData.Map()
	if err := sess.Save(eContext.Request(), eContext.Response()); err != nil {
		return nil, err
	}

	return sess, nil
}
