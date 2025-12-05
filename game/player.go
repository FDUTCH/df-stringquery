package game

import (
	"github.com/FDUTCH/go-stringquery/query"
	"github.com/df-mc/dragonfly/server/player"
)

var PlayerFormatter = query.NewFormatter[*player.Player]().
	WithOption("name", playerName).
	WithOption("nick", playerName).
	WithOption("nametag", nameTag).
	WithOption("nameTag", nameTag).
	WithOption("tag", nameTag).
	WithOption("food", food).
	WithOption("speed", speed).
	WithOption("health", health).
	WithOption("maxHealth", maxHealth).
	WithOption("pos", playerPosition).
	WithOption("position", playerPosition).
	WithOption("rot", playerRotation).
	WithOption("rotation", playerRotation).
	WithOption("yaw", playerYaw).
	WithOption("pitch", playerPitch).
	WithOption("type", playerType).
	WithOption("ip", ip).
	WithOption("address", ip).
	WithOption("experience", experience).
	WithOption("experienceLevel", experienceLevel).
	WithOption("experienceProggers", experienceProggers).
	WithOption("latency", latency).
	WithOption("ping", latency).
	WithOption("airSupply", airSupply).
	WithOption("maxAirSupply", maxAirSupply).
	WithOption("scale", scale).
	WithOption("onGround", onGround).
	WithOption("ground", onGround).
	WithOption("breathing", breathing).
	WithOption("eyeHeight", eyeHeight).
	WithOption("torsoHeight", torsoHeight).
	WithOption("enchantmentSeed", enchantmentSeed).
	WithOption("uuid", uuid).
	WithOption("id", uuid).
	WithOption("xuid", xuid).
	WithOption("deviceID", deviceID).
	WithOption("DID", deviceID).
	WithOption("deviceModel", deviceModel).
	WithOption("selfSignedID", selfSignedID).
	WithOption("locale", locale).
	WithOption("fallDistance", fallDistance).
	WithOption("flightSpeed", flightSpeed).
	WithOption("verticalFlightSpeed", verticalFlightSpeed).
	WithOption("absorption", absorption).
	WithOption("beaconAffected", beaconAffected).
	WithOption("dead", dead).
	WithOption("deathPos", deathPosition).
	WithOption("deathDim", deathDimension).
	WithOption("deathPosition", deathPosition).
	WithOption("deathDimension", deathDimension).
	WithOption("dim", playerDimension).
	WithOption("dimension", playerDimension).
	WithOption("world", playerDimension).
	WithOption("sprinting", sprinting).
	WithOption("sneaking", sneaking).
	WithOption("swimming", swimming).
	WithOption("crawling", crawling).
	WithOption("gliding", gliding).
	WithOption("flying", flying).
	WithOption("invisible", invisible).
	WithOption("immobile", immobile).
	WithOption("fireProof", fireProof).
	WithOption("mainHand", mainHand).
	WithOption("offHand", offHand).
	WithOption("helmet", helmet).
	WithOption("chestplate", chestplate).
	WithOption("leggings", leggings).
	WithOption("boots", boots).
	WithOption("gameMode", gameMode).
	WithOption("gamemode", gameMode).
	WithOption("gm", gameMode).
	WithOption("usingItem", usingItem)

func playerName(p *player.Player) any {
	return p.Name()
}

func nameTag(p *player.Player) any {
	return p.NameTag()
}

func food(p *player.Player) any {
	return p.Food()
}

func speed(p *player.Player) any {
	return p.Speed()
}

func health(p *player.Player) any {
	return p.Health()
}

func maxHealth(p *player.Player) any {
	return p.MaxHealth()
}

func playerPosition(p *player.Player) any {
	return entityPosition(p)
}

func playerRotation(p *player.Player) any {
	return entityRotation(p)
}

func playerYaw(p *player.Player) any {
	return entityYaw(p)
}

func playerPitch(p *player.Player) any {
	return entityPitch(p)
}

func playerType(p *player.Player) any {
	return entityType(p)
}

func experience(p *player.Player) any {
	return p.Experience()
}

func experienceLevel(p *player.Player) any {
	return p.ExperienceLevel()
}

func experienceProggers(p *player.Player) any {
	return p.ExperienceProgress()
}

func latency(p *player.Player) any {
	return p.Latency()
}

func airSupply(p *player.Player) any {
	return p.AirSupply()
}

func maxAirSupply(p *player.Player) any {
	return p.MaxAirSupply()
}

func scale(p *player.Player) any {
	return p.Scale()
}

func onGround(p *player.Player) any {
	return p.OnGround()
}

func breathing(p *player.Player) any {
	return p.Breathing()
}

func eyeHeight(p *player.Player) any {
	return p.EyeHeight()
}

func torsoHeight(p *player.Player) any {
	return p.TorsoHeight()
}

func enchantmentSeed(p *player.Player) any {
	return p.EnchantmentSeed()
}

func ip(p *player.Player) any {
	return p.Addr()
}

func uuid(p *player.Player) any {
	return p.UUID()
}

func xuid(p *player.Player) any {
	return p.XUID()
}

func deviceID(p *player.Player) any {
	return p.DeviceID()
}

func deviceModel(p *player.Player) any {
	return p.DeviceModel()
}

func selfSignedID(p *player.Player) any {
	return p.SelfSignedID()
}

func locale(p *player.Player) any {
	return p.Locale()
}

func fallDistance(p *player.Player) any {
	return p.FallDistance()
}

func flightSpeed(p *player.Player) any {
	return p.FlightSpeed()
}

func verticalFlightSpeed(p *player.Player) any {
	return p.VerticalFlightSpeed()
}

func absorption(p *player.Player) any {
	return p.Absorption()
}

func beaconAffected(p *player.Player) any {
	return p.BeaconAffected()
}

func dead(p *player.Player) any {
	return p.Dead()
}

func deathPosition(p *player.Player) any {
	pos, _, _ := p.DeathPosition()
	return formatVec3(pos)
}

func deathDimension(p *player.Player) any {
	_, dim, _ := p.DeathPosition()
	return dim
}

func playerDimension(p *player.Player) any {
	return p.Tx().World().Dimension()
}

func sprinting(p *player.Player) any {
	return p.Sprinting()
}

func sneaking(p *player.Player) any {
	return p.Sneaking()
}

func swimming(p *player.Player) any {
	return p.Swimming()
}

func crawling(p *player.Player) any {
	return p.Crawling()
}

func gliding(p *player.Player) any {
	return p.Gliding()
}

func flying(p *player.Player) any {
	return p.Flying()
}

func invisible(p *player.Player) any {
	return p.Invisible()
}

func immobile(p *player.Player) any {
	return p.Immobile()
}

func fireProof(p *player.Player) any {
	return p.FireProof()
}

func mainHand(p *player.Player) any {
	main, _ := p.HeldItems()
	return main
}

func offHand(p *player.Player) any {
	_, off := p.HeldItems()
	return off
}

func helmet(p *player.Player) any {
	return p.Armour().Helmet()
}

func chestplate(p *player.Player) any {
	return p.Armour().Chestplate()
}

func leggings(p *player.Player) any {
	return p.Armour().Leggings()
}

func boots(p *player.Player) any {
	return p.Armour().Boots()
}

func gameMode(p *player.Player) any {
	return formatGamemode(p.GameMode())
}

func usingItem(p *player.Player) any {
	return p.UsingItem()
}
