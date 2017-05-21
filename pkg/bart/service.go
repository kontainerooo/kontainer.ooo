// Package bart provides the permission management system (bus) for the kontainerooo plattform
package bart

import (
	"context"
	"errors"
	"reflect"
	"regexp"

	"github.com/kontainerooo/kontainer.ooo/pkg/user"
	ws "github.com/kontainerooo/kontainer.ooo/pkg/websocket"
)

var refRegexp = regexp.MustCompile("ref")

// Bus is a permission management system
type Bus interface {
	// GetOff should be used as a websocket before middleware
	GetOff(ws.ProtoID, ws.ProtoID, interface{}, interface{}) error

	// GetOn should be used as a websocket after middleware
	GetOn(ws.ProtoID, ws.ProtoID, interface{}, interface{}) error
}

type bus struct {
	ue      user.Endpoints
	admins  map[uint]bool
	tiers   map[uint]string
	fieldID map[string]int
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
	val := reflect.ValueOf(data).Elem().Elem()
	typ := val.Type()

	if val.Kind() != reflect.Struct {
		return errors.New("data malformatted")
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

func (b *bus) GetOff(srv, me ws.ProtoID, data interface{}, session interface{}) error {
	service := srv.String()
	method := me.String()

	if service == "KTG" && method == "AUT" {
		return nil
	}

	sessionMap, ok := session.(map[string]interface{})
	if !ok {
		return errors.New("session malformatted")
	}

	idInterface, ok := sessionMap["ID"]
	if !ok {
		return errors.New("no id present in session data")
	}

	id64, ok := idInterface.(float64)
	if !ok {
		return errors.New("id malformatted")
	}
	id := uint(id64)

	if b.IsAdmin(id) {
		return nil
	}

	err := b.CheckServiceAccess(service, method, id)
	if err != nil {
		return err
	}

	err = b.CheckTierConformity(service, method, data, id)
	if err != nil {
		return err
	}

	err = b.CheckID(service, method, data, id)
	if err != nil {
		return err
	}

	return nil
}

func (b *bus) GetOn(srv, me ws.ProtoID, data interface{}, session interface{}) error {

	return nil
}

// NewBus returns a new bus
func NewBus(ue user.Endpoints) Bus {
	return &bus{
		ue:      ue,
		admins:  make(map[uint]bool),
		tiers:   make(map[uint]string),
		fieldID: make(map[string]int),
	}
}
