package game

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

// Query ...
func Query(obj any, str string) string {
	switch val := obj.(type) {
	case *world.World:
		return WorldFormatter.Query(str, val)
	case *player.Player:
		return PlayerFormatter.Query(str, val)
	case world.Entity:
		return EntityFormatter.Query(str, val)
	}
	return str
}
