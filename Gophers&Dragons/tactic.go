package tactic

import "github.com/golangconf/gophers-and-dragons/game"

func ChooseCard(s game.State) game.CardType {
	return tactic2(s)
}

func tactic2(s game.State) game.CardType {

	if (s.Creep.Type == game.CreepClaws || s.Creep.Type == game.CreepDragon) && s.Can(game.CardParry) && s.Avatar.HP < 20 {
		return game.CardParry
	}

	if s.Creep.Type == game.CreepKubus && s.Can(game.CardPowerAttack) && s.Avatar.HP > 10 {
		return game.CardPowerAttack
	}

	if s.Avatar.HP > 10 {
		return game.CardAttack
	}

	if s.Avatar.HP < 4 && !(s.Can(game.CardHeal)) {
		return game.CardRetreat
	}

  if s.Avatar.HP < 4 && s.Can(game.CardHeal) {
		return game.CardHeal
	}

	if s.Avatar.HP < 5 && s.Can(game.CardStun) {
		return game.CardStun
	}

	if (s.Creep.Type == game.CreepFairy || s.Creep.Type == game.CreepClaws || s.Creep.Type == game.CreepDragon) && s.Can(game.CardStun) && s.Creep.IsStunned() == false {
		return game.CardStun
	}

	if s.Creep.IsStunned() == true && s.Avatar.HP < 10 && s.Can(game.CardHeal) {
		return game.CardHeal
	}

	if s.Creep.IsStunned() == true && s.Can(game.CardFirebolt) {
		return game.CardFirebolt
	}

	if (s.Creep.Type == game.CreepFairy || s.Creep.Type == game.CreepClaws) && s.Can(game.CardFirebolt) {
		return game.CardFirebolt
	}

	if s.Avatar.HP < 10 {
		if s.Can(game.CardHeal) {
			return game.CardHeal
		}
		return game.CardRetreat
	}

	if s.Creep.Type == game.CreepCheepy {
		return game.CardAttack
	}

	return game.CardRetreat
}
