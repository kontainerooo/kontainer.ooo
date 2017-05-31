package websocket

import (
	"context"
	"net/http"
	"reflect"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/sessions"
)

// The Authenticator may be implemented for use with the websocket Server
type Authenticator interface {
	// Mux is basicaly a http.Handler but with an interface and a boolean return parameter
	// the interface may contain session data
	// the boolean should be set to true if a response was send and the conection should not be upgraded
	Mux(http.ResponseWriter, *http.Request) (interface{}, bool)

	// GetID is used to set the Service ProtoID for the Authentication
	GetID() ProtoID

	// GetEndpoint is used to add the Logical Endpoint to the Authentication Service
	GetEndpoint() *ServiceEndpoint
}

// TODO: Add Token Invalidator
type tokenAuth struct {
	SessionStore *sessions.CookieStore
	ServiceID    ProtoID
	EndpointID   ProtoID
	SigningKey   []byte
	SessionName  string
	ValidTokens  []string
	DecodeFunc   DecodeRequestFunc
	EncodeFunc   EncodeResponseFunc
	Endpoint     endpoint.Endpoint
}

func (t *tokenAuth) Mux(w http.ResponseWriter, r *http.Request) (interface{}, bool) {
	if r.URL.Path == "/auth" {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil, true
		}
		tokenString := r.PostForm.Get("token")
		valid := false

		for i, r := range t.ValidTokens {
			if r == tokenString {
				valid = true
				t.ValidTokens = append(t.ValidTokens[:i], t.ValidTokens[i+1:]...)
				break
			}
		}

		if !valid {
			http.Error(w, "token invalid", http.StatusForbidden)
			return nil, true
		}

		token, err := jwt.ParseWithClaims(tokenString, &TokenAuthClaims{}, func(token *jwt.Token) (interface{}, error) {
			return t.SigningKey, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil, true
		}

		claims, ok := token.Claims.(*TokenAuthClaims)
		if !ok || !token.Valid {
			http.Error(w, "token invalid", http.StatusForbidden)
			return nil, true
		}

		session, err := t.SessionStore.Get(r, t.SessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil, true
		}
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		val := reflect.ValueOf(claims.Data)

		for _, key := range val.MapKeys() {
			session.Values[key.String()] = val.MapIndex(key).Interface()
		}

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil, true
		}

		w.Write([]byte("Authenticated!"))

		return nil, true
	}

	session, err := t.SessionStore.Get(r, t.SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, true
	}

	if session.IsNew {
		return nil, false
	}

	return session.Values, false
}

func (t *tokenAuth) GetID() ProtoID {
	return t.ServiceID
}

// TokenAuthClaims is the jwt Claims Type used by the TokenAuth
type TokenAuthClaims struct {
	Data interface{}
	jwt.StandardClaims
}

func (t *tokenAuth) tokenEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		values, err := t.Endpoint(ctx, request)
		if err != nil {
			return nil, err
		}

		c := TokenAuthClaims{
			Data: values,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

		s, err := token.SignedString(t.SigningKey)
		if err != nil {
			return nil, err
		}

		t.ValidTokens = append(t.ValidTokens, s)
		return s, nil
	}
}

func (t *tokenAuth) GetEndpoint() *ServiceEndpoint {
	e, _ := NewServiceEndpoint("Authentication", t.EndpointID, t.tokenEndpoint(), t.DecodeFunc, t.EncodeFunc)
	return e
}

// NewTokenAuth returns a new Authenticator which returns a JWT as a websocket response to a valid auth message
// This token can then be used to get a session cookie
func NewTokenAuth(
	serviceID, endpointID ProtoID,
	authenticationKey, encryptionKey string,
	signingKey string,
	sessionName string,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	end endpoint.Endpoint,
) Authenticator {
	// TODO: add validation of auth, encryptionKey
	return &tokenAuth{
		SessionStore: sessions.NewCookieStore([]byte(authenticationKey), []byte(encryptionKey)),
		ServiceID:    serviceID,
		EndpointID:   endpointID,
		SigningKey:   []byte(signingKey),
		SessionName:  sessionName,
		DecodeFunc:   dec,
		EncodeFunc:   enc,
		Endpoint:     end,
	}
}
