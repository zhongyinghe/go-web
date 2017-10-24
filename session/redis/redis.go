package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"session"
	"sync"
	"time"
)

var pder = &Provider{listKey: "sessionListKey"}
var redisConn redis.Conn
var timeAccessedKey = "timeAccessed"

type SessionStore struct {
	sid string
}

func CreateSessionStore(sid string) *SessionStore {
	return &SessionStore{sid: sid}
}

func (st *SessionStore) Set(key, value interface{}) error {
	if key == timeAccessedKey {
		fmt.Println("the key is not " + timeAccessedKey)
		return nil
	}
	_, err := redisConn.Do("HSET", st.sid, key, value)
	if err != nil {
		fmt.Println(err)
	}
	st.UpdateAccessedTime(st.sid)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
	value, err := redis.String(redisConn.Do("HGET", st.sid, key))
	st.UpdateAccessedTime(st.sid)
	pder.SessionUpdate(st.sid)
	if err == nil {
		return value
	}
	return nil
}

func (st *SessionStore) Delete(key interface{}) error {
	redisConn.Do("HDEL", st.sid, key)
	st.UpdateAccessedTime(st.sid)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

func (st *SessionStore) UpdateAccessedTime(sid string) {
	redisConn.Do("HSET", st.sid, timeAccessedKey, time.Now().Unix())
}

type Provider struct {
	lock    sync.Mutex //用来锁
	listKey string     //存储sessionId的队列键
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	newsess := CreateSessionStore(sid)
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	newsess := CreateSessionStore(sid)
	return newsess, nil
}

func (pder *Provider) SessionDestroy(sid string) error {
	redisConn.Do("DEL", sid)
	redisConn.Do("LREM", pder.listKey, 1, sid)
	return nil
}

func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	rdConn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic("Second Goroutine Connect to redis error")
	}
	rdConn.Do("AUTH", "123456") //授权

	for {
		sid, _ := redis.String(rdConn.Do("RPOP", pder.listKey))

		if sid == "" {
			break
		}

		timeAccessed, _ := redis.Int64(rdConn.Do("HGET", sid, timeAccessedKey))
		if timeAccessed+maxlifetime < time.Now().Unix() {
			//pder.SessionDestroy(sid)
			rdConn.Do("DEL", sid)
			rdConn.Do("LREM", pder.listKey, 1, sid)
		} else {
			rdConn.Do("RPUSH", pder.listKey, sid)
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	redisConn.Do("LREM", pder.listKey, 1, sid)
	redisConn.Do("LPUSH", pder.listKey, sid)
	return nil
}

func init() {
	var err error
	redisConn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic("Connect to redis error")
	}

	redisConn.Do("AUTH", "123456") //授权

	session.Register("redis", pder)
}
