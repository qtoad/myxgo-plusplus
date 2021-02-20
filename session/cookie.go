package session

import (
	_ "github.com/gin-gonic/gin"
	_ "net/http"
	_ "time"
)

//// WriteCookie は、ブラウザのcookieにセッションIDを書き込みます。
//func WriteCookie(c gin.Context, sessionID ID) error {
//	cookie := new(http.Cookie)
//	cookie.Name = setting.Session.CookieName
//	cookie.Value = string(sessionID)
//	cookie.Expires = time.Now().Add(setting.Session.CookieExpire)
//	c.SetCookie(cookie)
//	return nil
//}
//
//// ReadCookie は、ブラウザのcookieからセッションIDを読み込みます。
//func ReadCookie(c gin.Context) (ID, error) {
//	var sessionID ID
//	cookie, err := c.Cookie(setting.Session.CookieName)
//	if err != nil {
//		return sessionID, err
//	}
//	sessionID = ID(cookie.Value)
//	return sessionID, nil
//}
