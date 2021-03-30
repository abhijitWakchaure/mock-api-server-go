package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// Controller ...
type Controller struct {
	Users []User
}

// ReadUser ...
func (c *Controller) ReadUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for _, v := range c.Users {
		if v.ID == id && v.isDeleted == false {
			EncodeResponse(w, 200, v)
			return
		}
	}
	EncodeResponse(w, 404, "User not found")
}

// CreateUser ...
func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	for _, v := range c.Users {
		if v.Username == user.Username {
			EncodeResponse(w, 400, "User already exists")
			return
		}
	}
	user.ID = bson.NewObjectId().Hex()
	user.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	user.ModifiedAt = 0
	if user.Roles == nil {
		user.Roles = []Role{ROLEUSER}
	}
	c.Users = append(c.Users, user)
	EncodeResponse(w, 201, user)
}

// UpdateUser ...
func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var user User
	var eUser *User
	_ = json.NewDecoder(r.Body).Decode(&user)
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].ID == id {
			eUser = &c.Users[i]
			break
		}
	}
	if eUser == nil || eUser.isDeleted {
		EncodeResponse(w, 404, "User not found")
		return
	}
	if user.Username != "" {
		eUser.Username = user.Username
	}
	if user.Password != "" {
		eUser.Password = user.Password
	}
	if user.Fullname != "" {
		eUser.Fullname = user.Fullname
	}
	if user.Mobile != "" {
		eUser.Mobile = user.Mobile
	}
	if user.Roles != nil {
		eUser.Roles = user.Roles
	}
	eUser.Blocked = user.Blocked
	eUser.ModifiedAt = time.Now().UnixNano() / int64(time.Millisecond)
	EncodeResponse(w, 200, eUser)
}

// DeleteUser ...
func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].ID == id {
			c.Users[i].isDeleted = true
			EncodeResponse(w, 200, "User deleted successfully")
			return
		}
	}
	EncodeResponse(w, 400, "Invalid input")
}

// ListUsers ...
func (c *Controller) ListUsers(w http.ResponseWriter, r *http.Request) {
	var ret []User
	for _, v := range c.Users {
		if !v.isDeleted {
			ret = append(ret, v)
		}
	}
	EncodeResponse(w, 200, ret)
}

// EncodeResponse ...
func EncodeResponse(w http.ResponseWriter, code int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
