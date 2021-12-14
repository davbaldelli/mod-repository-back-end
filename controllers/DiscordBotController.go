package controllers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/davide/ModRepository/models/entities"
)

type DiscordBotControllerImpl struct {
	Session *discordgo.Session
	Channel string
}

func (d DiscordBotControllerImpl) NotifyCarAdded(car entities.Car) error {
	_, error := d.Session.ChannelMessageSendEmbed(d.Channel, &discordgo.MessageEmbed{
		URL:         fmt.Sprintf("https://www.acmodrepository.com/cars/%v",car.Id),
		Type:        "image",
		Title:       fmt.Sprintf("%v %v has been added to the repository!",car.Brand.Name, car.ModelName),
		Description: "Checkout here",
		Image:       &discordgo.MessageEmbedImage{URL: car.Image},
	})
	return error
}

func (d DiscordBotControllerImpl) NotifyCarUpdated(car entities.Car) error {
	_, error := d.Session.ChannelMessageSendEmbed(d.Channel, &discordgo.MessageEmbed{
		URL:         fmt.Sprintf("https://www.acmodrepository.com/cars/%v",car.Id),
		Type:        "image",
		Title:       fmt.Sprintf("%v %v has been updated!",car.Brand.Name, car.ModelName),
		Description: "Checkout here",
		Thumbnail:       &discordgo.MessageEmbedThumbnail{URL: car.Image},
	})
	return error
}

func (d DiscordBotControllerImpl) NotifyTrackUpdated(track entities.Track) error {
	_, error := d.Session.ChannelMessageSendEmbed(d.Channel, &discordgo.MessageEmbed{
		URL:         fmt.Sprintf("https://www.acmodrepository.com/tracks/%v",track.Id),
		Type:        "image",
		Title:       fmt.Sprintf("%v has been updated!",track.Name),
		Description: "Checkout here",
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: track.Image},
	})
	return error
}

func (d DiscordBotControllerImpl) NotifyTrackAdded(track entities.Track) error {
	_, error := d.Session.ChannelMessageSendEmbed(d.Channel, &discordgo.MessageEmbed{
		URL:         fmt.Sprintf("https://www.acmodrepository.com/tracks/%v",track.Id),
		Type:        "image",
		Title:       fmt.Sprintf("%v has been added to the repository!",track.Name),
		Description: "Checkout here",
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: track.Image},
	})
	return error
}

