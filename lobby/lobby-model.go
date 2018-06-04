package lobby

type Table struct {
  ID int
  Players []Player
}

type Player struct {
  PlayerID int
  Points float64
}

type Tables []Table
