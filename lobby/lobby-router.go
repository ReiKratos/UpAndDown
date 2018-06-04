package lobby

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

var controller = &Controller{Lobby: Lobby{}}

// Route defines a route
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
  Route {
    "Index",
    "GET",
    "/",
    controller.index,
  },
  Route {
    "CreateTable",
    "POST",
    "/create-table",
    controller.createNewTable,
  },
  Route {
    "JoinTable",
    "POST",
    "/join-table/{id}",
    controller.joinTable,
  },
  Route {
    "GetTableById",
    "GET",
    "/table/{id}",
    controller.getTableById,
  },
  Route {
    "GetTables",
    "GET",
    "/tables",
    controller.getTables,
  },
}

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
      var handler http.Handler
      log.Println("Registering route ", route.Name)
      handler = route.HandlerFunc

      router.
        Methods(route.Method).
        Path(route.Pattern).
        Name(route.Name).
        Handler(handler)
  }
  return router

}
