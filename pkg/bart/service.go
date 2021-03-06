// Package bart provides the permission management system (bus) for the kontainerooo plattform
package bart

import (
	"context"
	"errors"
	"reflect"
	"regexp"

	jwt "github.com/dgrijalva/jwt-go"
	pb "github.com/kontainerooo/kontainer.ooo/pkg/kentheguru/pb"
	"github.com/kontainerooo/kontainer.ooo/pkg/user"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

var refRegexp = regexp.MustCompile("ref")

// Bus is a permission management system
type Bus interface {
	// GetOff should be used as a websocket before middleware
	GetOff(ws.ProtoID, ws.ProtoID, *ws.MiddlewareData, interface{}) error

	// GetOn should be used as a websocket after middleware
	GetOn(ws.ProtoID, ws.ProtoID, *ws.MiddlewareData, interface{}) error

	// LostAndFound should be used as a websocket before middleware
	LostAndFound(ws.ProtoID, ws.ProtoID, *ws.MiddlewareData, interface{}) error
}

type bus struct {
	ue         user.Endpoints
	signingKey []byte
	admins     map[uint]bool
	tiers      map[uint]string
	fieldID    map[string]int
}

func (b *bus) IsAdmin(id uint) bool {
	admin, ok := b.admins[id]
	if !ok {
		res, err := b.ue.GetUserEndpoint(context.Background(), user.GetUserRequest{
			ID: id,
		})
		if err != nil {
			return false
		}
		b.admins[id] = res.(user.GetUserResponse).User.Admin
		admin = b.admins[id]
	}

	return admin
}

func (b *bus) CheckServiceAccess(srv, me string, id uint) error {
	switch srv {
	case "KMI":
		switch me {
		case "ADD":
		case "REM":
			return errors.New("not allowed")
		}
	}

	return nil
}

func (b *bus) CheckTierConformity(srv, me string, data interface{}, id uint) error {
	// TODO: Create Tiers
	return nil
}

func (b *bus) CheckID(srv, me string, data interface{}, id uint) error {
	val := reflect.ValueOf(data)
	typ := val.Type()

	if val.Kind() != reflect.Struct {
		return errors.New("data malformed")
	}

	fieldID, ok := b.fieldID[srv+me]

	if !ok {
		var i int
		for i = 0; i < val.NumField(); i++ {
			tag := typ.Field(i).Tag.Get("bart")
			if refRegexp.MatchString(tag) {
				b.fieldID[srv+me] = i
				fieldID = i
				break
			}
		}
		if fieldID == 0 && i > 0 {
			b.fieldID[srv+me] = -1
			fieldID = -1
		}
	}

	if fieldID != -1 {
		idInterface := val.Field(fieldID)
		if idInterface.Uint() != uint64(id) {
			return errors.New("wrong id")
		}
	}

	return nil
}

func (b *bus) GetOff(srv, me ws.ProtoID, data *ws.MiddlewareData, session interface{}) error {
	service := srv.String()
	method := me.String()

	if service == "KTG" && method == "AUT" {
		return nil
	}

	sessionMap, ok := session.(map[interface{}]interface{})
	if !ok {
		return errors.New("session malformed")
	}

	idInterface, ok := sessionMap["ID"]
	if !ok {
		return errors.New("no id present in session data")
	}

	id64, ok := idInterface.(float64)
	if !ok {
		return errors.New("id malformed")
	}
	id := uint(id64)

	if b.IsAdmin(id) {
		return nil
	}

	err := b.CheckServiceAccess(service, method, id)
	if err != nil {
		return err
	}

	err = b.CheckTierConformity(service, method, data.Value, id)
	if err != nil {
		return err
	}

	err = b.CheckID(service, method, data.Value, id)
	if err != nil {
		return err
	}

	return nil
}

func (b *bus) GetOn(srv, me ws.ProtoID, data *ws.MiddlewareData, session interface{}) error {
	if srv == ws.ProtoIDFromString("KTG") && me == ws.ProtoIDFromString("AUT") {
		res, ok := data.Value.(*pb.AuthenticationResponse)
		if !ok {
			return errors.New("malformed response")
		}

		if res.Error != "" {
			return nil
		}

		token, err := jwt.ParseWithClaims(res.Token, &ws.TokenAuthClaims{}, func(token *jwt.Token) (interface{}, error) {
			return b.signingKey, nil
		})
		if err != nil {
			return err
		}

		claimsWrapper, ok := token.Claims.(*ws.TokenAuthClaims)
		if !ok || !token.Valid {
			return errors.New("token invalid")
		}

		claimsData, ok := claimsWrapper.Data.(map[string]interface{})
		if !ok {
			return errors.New("malformed claims")
		}

		val := reflect.ValueOf(session).Elem()
		valMap := make(map[interface{}]interface{})

		for k, v := range claimsData {
			valMap[k] = v
		}

		val.Set(reflect.ValueOf(valMap))
	}

	return nil
}

func (b *bus) LostAndFound(srv, me ws.ProtoID, data *ws.MiddlewareData, session interface{}) error {
	if srv != ws.ProtoIDFromString("USR") || me != ws.ProtoIDFromString("GET") {
		return nil
	}

	req, ok := data.Value.(user.GetUserRequest)
	if !ok {
		return errors.New("request malformed")
	}

	if req.ID != 0 {
		return nil
	}

	sessionMap, ok := session.(map[interface{}]interface{})
	if !ok {
		return errors.New("session malformed")
	}

	idInterface, ok := sessionMap["ID"]
	if !ok {
		return errors.New("no id present in session data")
	}

	id64, ok := idInterface.(float64)
	if !ok {
		return errors.New("id malformed")
	}
	req.ID = uint(id64)

	reflect.ValueOf(data).Elem().Set(reflect.ValueOf(ws.MiddlewareData{Value: req}))

	return nil
}

// NewBus returns a new bus
func NewBus(signingKey string, ue user.Endpoints) Bus {
	return &bus{
		ue:         ue,
		signingKey: []byte(signingKey),
		admins:     make(map[uint]bool),
		tiers:      make(map[uint]string),
		fieldID:    make(map[string]int),
	}
}
