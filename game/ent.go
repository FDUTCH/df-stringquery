package game

import (
	"fmt"
	"strconv"

	"github.com/FDUTCH/go-stringquery/query"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

var EntityFormatter = query.NewFormatter[world.Entity]().
	WithOption("pos", entityPosition).
	WithOption("position", entityPosition).
	WithOption("rot", entityRotation).
	WithOption("rotation", entityRotation).
	WithOption("yaw", entityYaw).
	WithOption("pitch", entityPitch).
	WithOption("type", entityType)

func entityPosition(entity world.Entity) any {
	pos := entity.Position()
	return formatVec3(pos)
}

func entityRotation(entity world.Entity) any {
	rot := entity.Rotation()
	return fmt.Sprintf("y%.2f p%.2f", rot.Yaw(), rot.Pitch())
}

func entityYaw(entity world.Entity) any {
	return strconv.FormatFloat(entity.Rotation().Yaw(), 'f', 2, 64)
}

func entityPitch(entity world.Entity) any {
	return strconv.FormatFloat(entity.Rotation().Pitch(), 'f', 2, 64)
}

func entityType(entity world.Entity) any {
	return entity.H().Type().EncodeEntity()
}

func formatVec3(pos mgl64.Vec3) string {
	return fmt.Sprintf("%.2f %.2f %.2f", pos.X(), pos.Y(), pos.Z())
}

func formatGamemode(gamemode world.GameMode) string {
	switch gamemode {
	case world.GameModeAdventure:
		return "Adventure"
	case world.GameModeSurvival:
		return "Survival"
	case world.GameModeCreative:
		return "Creative"
	case world.GameModeSpectator:
		return "Spectator"
	}
	return fmt.Sprintf("%T", gamemode)
}
