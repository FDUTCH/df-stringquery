package game

import (
	"fmt"

	"github.com/FDUTCH/go-stringquery/query"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

var WorldFormatter = query.NewFormatter[*world.World]().
	WithOption("name", worldName).
	WithOption("dimension", worldDimension).
	WithOption("time", time).
	WithOption("worldSpawn", worldSpawn).
	WithOption("difficulty", difficulty)

func worldName(w *world.World) any {
	return w.Name()
}

func worldDimension(w *world.World) any {
	return w.Dimension()
}

func time(w *world.World) any {
	return w.Time()
}

func worldSpawn(w *world.World) any {
	return formatPos(w.Spawn())
}

func difficulty(w *world.World) any {
	return formatDifficulty(w.Difficulty())
}

func formatPos(pos cube.Pos) string {
	return fmt.Sprintf("%d %d %d", pos.X(), pos.Y(), pos.Z())
}

func formatDifficulty(difficulty world.Difficulty) string {
	switch difficulty {
	case world.DifficultyPeaceful:
		return "Peaceful"
	case world.DifficultyEasy:
		return "Easy"
	case world.DifficultyNormal:
		return "Normal"
	case world.DifficultyHard:
		return "Hard"
	}
	return fmt.Sprintf("%T", difficulty)
}
