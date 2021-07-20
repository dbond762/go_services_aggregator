package memory

import (
	"container/list"
	"github.com/dbond762/go_services_aggregator/src/libs/session"
	"sync"
	"time"
)

var provider = &Provider{list: list.New()}

type SessionStore struct {
	sid          string                      // unique session id
	timeAccessed time.Time                   // last access time
	value        map[interface{}]interface{} // session value stored inside
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	provider.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
	provider.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	provider.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

type Provider struct {
	lock     sync.Mutex               // lock
	sessions map[string]*list.Element // save in memory
	list     *list.List               // gc
}

func (p *Provider) SessionInit(sid string) (session.Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newSess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := p.list.PushBack(newSess)
	p.sessions[sid] = element
	return newSess, nil
}

func (p *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := p.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := p.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (p *Provider) SessionDestroy(sid string) error {
	if element, ok := p.sessions[sid]; ok {
		delete(p.sessions, sid)
		p.list.Remove(element)
		return nil
	}
	return nil
}

func (p *Provider) SessionGC(maxLifeTime int64) {
	p.lock.Lock()
	defer p.lock.Unlock()

	for {
		element := p.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			p.list.Remove(element)
			delete(p.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (p *Provider) SessionUpdate(sid string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if element, ok := p.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		p.list.MoveToFront(element)
		return nil
	}
	return nil
}

func init() {
	provider.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", provider)
}
