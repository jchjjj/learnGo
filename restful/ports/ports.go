package main

import (
    "fmt"
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "sync"
)

func main() {

    ports := ports{
        Store: map[string]*Port{},
    }

    handler := rest.ResourceHandler{
        EnableRelaxedContentType: true,
    }
handler.SetRoutes(
        rest.RouteObjectMethod("GET", "/ports", &ports, "GetAllPorts"),
        rest.RouteObjectMethod("POST", "/ports", &ports, "PostPort"),
        rest.RouteObjectMethod("GET", "/ports/:id/:name", &ports, "GetPort"),
        rest.RouteObjectMethod("PUT", "/ports/:id", &ports, "PutPort"),
        rest.RouteObjectMethod("DELETE", "/ports/:id", &ports, "DeletePort"),
    )
    http.ListenAndServe(":8080", &handler)
}

type Port struct {
    Id   string
    Name string
}

type ports struct {
    sync.RWMutex
    Store map[string]*Port
}

func (u *ports) GetAllPorts(w rest.ResponseWriter, r *rest.Request) {
    u.RLock()
    ports := make([]*Port, len(u.Store))
    i := 0
    for _, Port := range u.Store {
        ports[i] = Port
        i++
    }
    u.RUnlock()
    w.WriteJson(&ports)
}

func (u *ports) GetPort(w rest.ResponseWriter, r *rest.Request) {// changed by jch
    id := r.PathParam("id")
    name := r.PathParam("name")
    u.RLock()
    Port := u.Store[id]
    Name := Port.Name 
    Name = Name + name
    u.RUnlock()
    if Port == nil {
        rest.NotFound(w, r)
        return
    }
  //  w.WriteJson(&Port)

    w.WriteJson(&Name)
}

func (u *ports) PostPort(w rest.ResponseWriter, r *rest.Request) {
    Port := Port{}
    err := r.DecodeJsonPayload(&Port)
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    u.Lock()
    id := fmt.Sprintf("%d", len(u.Store)) // stupid
    Port.Id = id
    u.Store[id] = &Port
    u.Unlock()
    w.WriteJson(&Port)
}
func (u *ports) PutPort(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    u.Lock()
    if u.Store[id] == nil {
        rest.NotFound(w, r)
        return
    }
    Port := Port{}
    err := r.DecodeJsonPayload(&Port)
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    Port.Id = id
    u.Store[id] = &Port
    u.Unlock()
    w.WriteJson(&Port)
}

func (u *ports) DeletePort(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    u.Lock()
    delete(u.Store, id)
    u.Unlock()
}
