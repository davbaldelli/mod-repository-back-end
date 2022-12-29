package controllers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/davide/ModRepository/models/entities"
	"net/url"
	"strconv"
)

type DiscordBotControllerImpl struct {
	Session  *discordgo.Session
	Channels []string
}

func (d DiscordBotControllerImpl) NotifyCarAdded(car entities.Car) error {
	for _, channel := range d.Channels {
		_, error := d.Session.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Type:        "image",
				Title:       fmt.Sprintf("%v %v has been added to the repository!", car.Brand.Name, car.ModelName),
				Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/cars/%v/%v/%v)", url.PathEscape(car.Brand.Name), url.PathEscape(car.ModelName), car.Year),
				Color:       12590120,
				Image:       &discordgo.MessageEmbedImage{URL: car.Images[0].Url},
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
		if error != nil {
			return error
		}
	}
	return nil
}

func (d DiscordBotControllerImpl) NotifyCarUpdated(car entities.Car) error {
	for _, channel := range d.Channels {
		_, error := d.Session.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Type:        "image",
				Title:       fmt.Sprintf("%v %v has been updated!", car.Brand.Name, car.ModelName),
				Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/cars/%v/%v/%v)", url.PathEscape(car.Brand.Name), url.PathEscape(car.ModelName), car.Year),
				Color:       12590120,
				Image:       &discordgo.MessageEmbedImage{URL: car.Images[0].Url},
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
		if error != nil {
			return error
		}
	}
	return nil
}

func (d DiscordBotControllerImpl) NotifyTrackUpdated(track entities.Track) error {
	for _, channel := range d.Channels {
		_, error := d.Session.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
			Type:        "image",
			Title:       fmt.Sprintf("%v has been updated!", track.Name),
			Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/tracks/detail/%v)", track.Id),
			Color:       12590120,
			Image:       &discordgo.MessageEmbedImage{URL: track.Images[0].Url},
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
		if error != nil {
			return error
		}
	}
	return nil
}

func (d DiscordBotControllerImpl) NotifyTrackAdded(track entities.Track) error {
	for _, channel := range d.Channels {
		_, error := d.Session.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
			Type:        "image",
			Title:       fmt.Sprintf("%v has been added to the repository!", track.Name),
			Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/tracks/detail/%v)", track.Id),
			Color:       12590120,
			Image:       &discordgo.MessageEmbedImage{URL: track.Images[0].Url},
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
		if error != nil {
			return error
		}
	}
	return nil
}
