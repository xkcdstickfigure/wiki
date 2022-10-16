package sessionAuth

import (
	"net/http"

	"alles/wiki/env"
	"alles/wiki/store"
)

const cookieName = "wiki_token"

func GetSession(db store.Store, r *http.Request) (store.Session, error) {
	token, err := r.Cookie(cookieName)
	if err != nil {
		return store.Session{}, err
	}

	session, err := db.SessionGetByToken(r.Context(), token.Value)
	return session, err
}

func UseSession(db store.Store, w http.ResponseWriter, r *http.Request) (store.Session, error) {
	// read existing session
	session, err := GetSession(db, r)
	if err == nil {
		return session, nil
	}

	// create new session
	session, err = db.SessionCreate(r.Context(), store.Session{
		Address:   r.RemoteAddr,
		UserAgent: r.Header.Get("user-agent"),
	})
	if err != nil {
		return session, err
	}

	// write cookie
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    session.Token,
		Domain:   env.Domain,
		Path:     "/",
		MaxAge:   365 * 24 * 60 * 60,
		SameSite: http.SameSiteLaxMode,
	})

	return session, err
}
