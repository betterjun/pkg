/*
	包mongo，Mongodb访问模块，封装对mongodb的数据访问接口。
*/
package mongodb

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

type DBSession struct {
	*mgo.Session
	Name string
}

/// 连接数据库.
func (session *DBSession) Connect(servers, dbname, user, password string, timeout, maxlink int) (err error) {
	di := &mgo.DialInfo{
		Addrs:    strings.Split(strings.TrimSpace(servers), ";"),
		Username: user,
		Password: password,
		Database: dbname,
		//Source:    dbname,
		Timeout:   time.Second * time.Duration(timeout),
		PoolLimit: maxlink,
		FailFast:  true,
	}

	session.Session, err = mgo.DialWithInfo(di)
	if err != nil {
		return err
	}
	session.Session.SetMode(mgo.Monotonic, true)
	session.Name = dbname

	return err
}

/// 建立一个副本session，副本用完就需要关闭.
func (session *DBSession) Clone() *DBSession {
	newSession := &DBSession{}
	newSession.Session = session.Session.Clone()
	newSession.Name = session.Name
	return newSession
}

/// 从session获取collection，已绑定数据库.
func (session *DBSession) Collection(name string) *mgo.Collection {
	return session.DB(session.Name).C(name)
}

/// 全局默认数据库session.
var defaultDBSession *DBSession

/// 初始化默认Mongodb数据库.
func InitDefaultDBSession(servers, dbname, user, password string, timeout, maxlink int) (err error) {
	defaultDBSession = &DBSession{}
	return defaultDBSession.Connect(servers, dbname, user, password, timeout, maxlink)
}

/// 获取一个副本session，副本用完就需要关闭.
func GetClonedSession() *DBSession {
	if defaultDBSession == nil {
		panic("empty mongodb session")
	}

	return defaultDBSession.Clone()
}
