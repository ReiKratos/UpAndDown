package lobby

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
    "net/http"
    "strconv"
)

//Controller ...
type Controller struct {
  Lobby Lobby
}

// Create table.
func (c *Controller) createNewTable(w http.ResponseWriter, r *http.Request) {
  newTable := c.Lobby.createTable()

  if newTable.ID == 0 {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  data, _ := json.Marshal(newTable)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  w.Write(data)
  return
}

// Create table.
func (c *Controller) index(w http.ResponseWriter, r *http.Request) {
  data, _ := json.Marshal("Welcome to Sobidesce")

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.Write(data)
  return
}

// Get table by ID.
func (c *Controller) getTableById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  tableId, err := strconv.Atoi(id);

  if err != nil {
    fmt.Println("Error on table ID", err)
  }

  table := c.Lobby.getTableById(tableId)

  if table == nil {
    w.WriteHeader(404)
    w.Write([]byte("404 - Table not found."))
    return
  }

  data, _ := json.Marshal(table)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  w.Write(data)
  return
}

// Get all tables.
func (c *Controller) getTables(w http.ResponseWriter, r *http.Request) {
  tables := c.Lobby.getTables()

  if tables == nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  data, _ := json.Marshal(tables)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  w.Write(data)
  return
}

// Join player.
func (c *Controller) joinTable(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  tableId, err := strconv.Atoi(id);

  if err != nil {
    http.Error(w, "Not valid table", 400)
    return
  }

  if r.Body == nil {
    http.Error(w, "Please send a request body", 400)
    return
  }

  body, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()

  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  var player Player
  err = json.Unmarshal(body, &player)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

  table, err := c.Lobby.joinTable(&player, tableId)

  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  data, _ := json.Marshal(table)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  w.Write(data)
  return
}
