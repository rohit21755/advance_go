package main

import "fmt"

type Player struct {
	health int
}

func takeDamageFromExplosion(player Player) {
	fmt.Println("player is taking damage from explosion")
	explisionDamage := 20
	player.health -= explisionDamage
}

func main() {

}
