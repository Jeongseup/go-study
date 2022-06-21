package interfaces

import (
	"study/models"
)

type IPlayerRepository interface {
	GetPlayerByName(name string) (models.PlayerModel, error)
}
