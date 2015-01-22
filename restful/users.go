package main

import (
    "fmt"
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "sync"
)

func main() {

    users := Users{
        Store: map[string]*User{},
    }

    handler := rest.ResourceHandler{
        EnableRelaxedContentType: true,
    }
handler.SetRoutes(
        rest.RouteObjectMethod("GET", "/users", &users, "GetAllUsers"),
        rest.RouteObjectMethod("POST", "/users", &users, "PostUser"),
        rest.RouteObjectMethod("GET", "/users/:id/:name", &users, "GetUser"),
        rest.RouteObjectMethod("PUT", "/users/:id", &users, "PutUser"),
        rest.RouteObjectMethod("DELETE", "/users/:id", &users, "DeleteUser"),
    )
    http.ListenAndServe(":8080", &handler)
}

type User struct {
    Id   string
    Name string
}

type Users struct {
    sync.RWMutex
    Store map[string]*User
}

func (u *Users) GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
    u.RLock()
    users := make([]*User, len(u.Store))
    i := 0
    for _, user := range u.Store {
        users[i] = user
        i++
    }
    u.RUnlock()
    w.WriteJson(&users)
}

func (u *Users) GetUser(w rest.ResponseWriter, r *rest.Request) {// changed by jch
    id := r.PathParam("id")
    name := r.PathParam("name")
    u.RLock()
    user := u.Store[id]
    Name := user.Name 
    Name = Name + name
    u.RUnlock()
    if user == nil {
        rest.NotFound(w, r)
        return
    }
  //  w.WriteJson(&user)

    w.WriteJson(&Name)
}

func (u *Users) PostUser(w rest.ResponseWriter, r *rest.Request) {
    user := User{}
    err := r.DecodeJsonPayload(&user)
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    u.Lock()
    id := fmt.Sprintf("%d", len(u.Store)) // stupid
    user.Id = id
    u.Store[id] = &user
    u.Unlock()
    w.WriteJson(&user)
}
func (u *Users) PutUser(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    u.Lock()
    if u.Store[id] == nil {
        rest.NotFound(w, r)
        return
    }
    user := User{}
    err := r.DecodeJsonPayload(&user)
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    user.Id = id
    u.Store[id] = &user
    u.Unlock()
    w.WriteJson(&user)
}

func (u *Users) DeleteUser(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    u.Lock()
    delete(u.Store, id)
    u.Unlock()
}
