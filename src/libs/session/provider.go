package session

type Provider interface {
	SessionInit(id string) (Session, error)
	SessionRead(id string) (Session, error)
	SessionDestroy(id string) error
	SessionGC(maxLifeTime int64)
}
