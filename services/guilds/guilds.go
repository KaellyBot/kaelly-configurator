package guilds

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/guilds"
)

type GuildService interface {
	GetGuild(id string) (entities.Guild, error)
}

type GuildServiceImpl struct {
	guildRepo guilds.GuildRepository
}

func New(guildRepo guilds.GuildRepository) (*GuildServiceImpl, error) {
	return &GuildServiceImpl{guildRepo: guildRepo}, nil
}

func (service *GuildServiceImpl) GetGuild(id string) (entities.Guild, error) {
	return service.guildRepo.GetGuild(id)
}