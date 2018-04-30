package libs

import (
	"time"
	"sync"
	"net/http"
	"net/url"
	"io"
	"crypto/rand"
	"strconv"
	"encoding/base64"
)

type User struct {
Account  string
Nickname string
Mail     string
Password string
Time     int64
}

/*会话*/
type Session struct {
SessionID        string            //唯一id
LastTimeAccessed time.Time         //最后访问时间
Values           map[string]*User  //其它对应值(保存用户所对应的一些值，比如用户权限之类)
}

/*Session会话管理*/
type SessionMgr struct {
CookieName  string            	 //客户端cookie名称
Lock        sync.RWMutex      	 //互斥(保证线程安全)
MaxLifeTime int64        	     //垃圾回收时间
Sessions    map[string]*Session  //保存session的指针[sessionID] = session
}

//创建会话管理器(cookieName:在浏览器中cookie的名字;maxLifeTime:最长生命周期)
func NewSessionMgr(cookieName string, maxLifeTime int64) *SessionMgr {
mgr := &SessionMgr{CookieName: cookieName, MaxLifeTime: maxLifeTime, Sessions: make(map[string]*Session)}
//启动定时回收
go mgr.GC()
return mgr
}

//在开始页面登陆页面，开始Session
func (mgr *SessionMgr) StartSession(w http.ResponseWriter, r *http.Request) string {
mgr.Lock.Lock()
defer mgr.Lock.Unlock()

//无论原来有没有，都重新创建一个新的session
newSessionID := url.QueryEscape(mgr.NewSessionID())

//存指针
var session *Session = &Session{SessionID: newSessionID, LastTimeAccessed: time.Now(), Values: make(map[string]*User)}
mgr.Sessions[newSessionID] = session
//让浏览器cookie设置过期时间

cookie := http.Cookie{Name: mgr.CookieName, Value: newSessionID, Path: "/", HttpOnly: true, MaxAge: int(mgr.MaxLifeTime)}
http.SetCookie(w, &cookie)

return newSessionID
}

//结束Session
func (mgr *SessionMgr) EndSession(w http.ResponseWriter, r *http.Request) {
cookie, err := r.Cookie(mgr.CookieName)
if err != nil || cookie.Value == "" {
return
} else {
mgr.Lock.Lock()
defer mgr.Lock.Unlock()

delete(mgr.Sessions, cookie.Value)

//让浏览器cookie立刻过期
expiration := time.Now()
cookie := http.Cookie{Name: mgr.CookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
http.SetCookie(w, &cookie)
}
}

//结束session
func (mgr *SessionMgr) EndSessionBy(sessionID string) {
mgr.Lock.Lock()
defer mgr.Lock.Unlock()

delete(mgr.Sessions, sessionID)
}

//设置session里面的值
func (mgr *SessionMgr) SetSessionVal(sessionID string, account string, user *User) {
mgr.Lock.Lock()
defer mgr.Lock.Unlock()

if session, ok := mgr.Sessions[sessionID]; ok {
session.Values[account] = user
}
}

//得到session里面的值
func (mgr *SessionMgr) GetSessionVal(sessionID string, account string) (interface{}, bool) {
mgr.Lock.RLock()
defer mgr.Lock.RUnlock()

if session, ok := mgr.Sessions[sessionID]; ok {
if val, ok := session.Values[account]; ok {
return val, ok
}
}
return nil, false
}

//得到sessionID列表
func (mgr *SessionMgr) GetSessionIDList() []string {
mgr.Lock.RLock()
defer mgr.Lock.RUnlock()

sessionIDList := make([]string, 0)

for k, _ := range mgr.Sessions {
sessionIDList = append(sessionIDList, k)
}

return sessionIDList[0:len(sessionIDList)]
}

//判断Cookie的合法性（每进入一个页面都需要判断合法性）
func (mgr *SessionMgr) CheckCookieValid(w http.ResponseWriter, r *http.Request) string {
var cookie, err = r.Cookie(mgr.CookieName)

if cookie == nil || err != nil {
	return ""
}

mgr.Lock.Lock()
defer mgr.Lock.Unlock()

sessionID := cookie.Value

if session, ok := mgr.Sessions[sessionID]; ok {
session.LastTimeAccessed = time.Now() //判断合法性的同时，更新最后的访问时间
return sessionID
}

return ""
}

//更新最后访问时间
func (mgr *SessionMgr) GetLastAccessTime(sessionID string) time.Time {
mgr.Lock.RLock()
defer mgr.Lock.RUnlock()

if session, ok := mgr.Sessions[sessionID]; ok {
return session.LastTimeAccessed
}
return time.Now()
}

//GC回收
func (mgr *SessionMgr) GC() {
mgr.Lock.Lock()
defer mgr.Lock.Unlock()

for sessionID, session := range mgr.Sessions {
//删除超过时限的session
if session.LastTimeAccessed.Unix()+mgr.MaxLifeTime < time.Now().Unix() {
delete(mgr.Sessions, sessionID)
}
}

//定时回收
time.AfterFunc(time.Duration(mgr.MaxLifeTime)*time.Second, func() { mgr.GC() })
}

//创建唯一ID
func (mgr *SessionMgr) NewSessionID() string {
b := make([]byte, 32)
if _, err := io.ReadFull(rand.Reader, b); err != nil {
nano := time.Now().UnixNano() //微秒
return strconv.FormatInt(nano, 10)
}
return base64.URLEncoding.EncodeToString(b)
}
