package lobby

import (
  "errors"
  "fmt"
)

type Lobby struct {
  Tables []Table
}

func (l *Lobby) joinTable(p *Player, tableId int) (*Table, error) {
  table := l.getTableById(tableId)

  if table == nil {
    return nil, errors.New("Table not found.")
  }

  if len(table.Players) >= 5 {
    return nil, errors.New("This table is full.")
  }

  table.Players = append(table.Players, *p)
  fmt.Println("A player has joined the table.")

  return table, nil
}

//Create table
func (l *Lobby) createTable() *Table {
  id := len(l.Tables) + 1
  players := make([]Player, 0, 5)
  table := Table{id, players}
  l.Tables = append(l.Tables, table)

  fmt.Println("Creating new table...")

  return &table
}

//Get table by ID
func (l *Lobby) getTableById(tableId int) *Table {
  for _, table := range l.Tables {
    if table.ID == tableId {
      return &table
    }
  }

  fmt.Println("Table not found.")
  return nil
}

//Get all tables
func (l *Lobby) getTables() []Table {
  return l.Tables
}
