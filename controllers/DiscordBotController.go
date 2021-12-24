package controllers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/davide/ModRepository/models/entities"
	"strconv"
)

type DiscordBotControllerImpl struct {
	Session *discordgo.Session
	Channel string
}

func (d DiscordBotControllerImpl) NotifyCarAdded(car entities.Car) error {
	_, error := d.Session.ChannelMessageSendComplex(d.Channel, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Type:        "image",
			Title:       fmt.Sprintf("%v %v has been added to the repository!", car.Brand.Name, car.ModelName),
			Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/cars/%v/%v/%v/)", car.Brand.Name, car.ModelName, car.Year),
			Color:       12590120,
			Image:       &discordgo.MessageEmbedImage{URL: car.Image},
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Davide",
				IconURL: "https://i.imgur.com/M4Am9z1.jpg",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Year",
					Value: strconv.Itoa(int(car.Year)),
				},
				{
					Name:  "Author",
					Value: fmt.Sprintf("[%v](%v)", car.Author.Name, car.Author.Link),
				},
			},
		},
	})
	return error
}

func (d DiscordBotControllerImpl) NotifyCarUpdated(car entities.Car) error {
	_, error := d.Session.ChannelMessageSendComplex(d.Channel, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Type:        "image",
			Title:       fmt.Sprintf("%v %v has been updated!", car.Brand.Name, car.ModelName),
			Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/cars/%v/%v/%v/)", car.Brand.Name, car.ModelName, car.Year),
			Color:       12590120,
			Image:       &discordgo.MessageEmbedImage{URL: car.Image},
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Davide",
				IconURL: "https://i.imgur.com/M4Am9z1.jpg",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Year",
					Value: strconv.Itoa(int(car.Year)),
				},
				{
					Name:   "Author",
					Value:  fmt.Sprintf("[%v](%v)", car.Author.Name, car.Author.Link),
					Inline: true,
				},
				{
					Name:   "Version",
					Value:  car.Version,
					Inline: true,
				},
			},
		},
	})
	return error
}

func (d DiscordBotControllerImpl) NotifyTrackUpdated(track entities.Track) error {
	_, error := d.Session.ChannelMessageSendEmbed(d.Channel, &discordgo.MessageEmbed{
		Type:        "image",
		Title:       fmt.Sprintf("%v has been updated!", track.Name),
		Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/tracks/%v)", track.Id),
		Color:       12590120,
		Image:       &discordgo.MessageEmbedImage{URL: track.Image},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Davide",
			IconURL: "https://i.imgur.com/M4Am9z1.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Location",
				Value: fmt.Sprintf("%v, %v", track.Location, track.Nation.Name),
			},
			{
				Name:  "Year",
				Value: strconv.Itoa(int(track.Year)),
			},
			{
				Name:   "Author",
				Value:  fmt.Sprintf("[%v](%v)", track.Author.Name, track.Author.Link),
				Inline: true,
			},
			{
				Name:   "Version",
				Value:  track.Version,
				Inline: true,
			},
		},
	})
	return error
}

func (d DiscordBotControllerImpl) NotifyTrackAdded(track entities.Track) error {
	_, error := d.Session.ChannelMessageSendEmbed(d.Channel, &discordgo.MessageEmbed{
		Type:        "image",
		Title:       fmt.Sprintf("%v has been added to the repository!", track.Name),
		Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/tracks/%v)", track.Id),
		Color:       12590120,
		Image:       &discordgo.MessageEmbedImage{URL: track.Image},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Davide",
			IconURL: "https://i.imgur.com/M4Am9z1.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Location",
				Value:  fmt.Sprintf("%v, %v", track.Location, track.Nation.Name),
				Inline: true,
			},
			{
				Name:   "Year",
				Value:  strconv.Itoa(int(track.Year)),
				Inline: true,
			},
			{
				Name:  "Author",
				Value: fmt.Sprintf("[%v](%v)", track.Author.Name, track.Author.Link),
			},
		},
	})
	return error
}
