package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q", provideName)
	}

	manager := &Manager{
		cookieName:  cookieName,
		provider:    provider,
		maxLifeTime: maxLifeTime,
	}

	return manager, nil
}

var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic(fmt.Sprintf("session: Register called twice for provider %s", name))
	}

	provides[name] = provider
}

func (m *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	m.lock.Lock()
	defer m.lock.Unlock()

	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		id := m.sessionID()
		session, _ = m.provider.SessionInit(id)
		cookie := &http.Cookie{
			Name:     m.cookieName,
			Value:    url.QueryEscape(id),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(m.maxLifeTime),
		}
		http.SetCookie(w, cookie)
	} else {
		id, _ := url.QueryUnescape(cookie.Value)
		session, _ = m.provider.SessionRead(id)
	}

	return
}

func (m *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		m.lock.Lock()
		defer m.lock.Unlock()

		_ = m.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := &http.Cookie{
			Name:     m.cookieName,
			Path:     "/",
			HttpOnly: true,
			Expires:  expiration,
			MaxAge:   -1,
		}
		http.SetCookie(w, cookie)
	}
}

func (m *Manager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.provider.SessionGC(m.maxLifeTime)
	time.AfterFunc(time.Duration(m.maxLifeTime), func() {
		m.GC()
	})
}
