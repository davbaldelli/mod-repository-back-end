package controllers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/davide/ModRepository/models"
	"net/url"
	"strconv"
)

type DiscordBotControllerImpl struct {
	Session  *discordgo.Session
	Channels []string
}

func (d DiscordBotControllerImpl) NotifyCarAdded(car models.Car) error {
	for _, channel := range d.Channels {
		_, err := d.Session.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Type:        "image",
				Title:       fmt.Sprintf("%v %v has been added to the repository!", car.Brand.Name, car.ModelName),
				Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/cars/%v/%v/%v)", url.PathEscape(car.Brand.Name), url.PathEscape(car.ModelName), car.Year),
				Color:       12590120,
				Image:       &discordgo.MessageEmbedImage{URL: getFavImageUrl(car.Images)},
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
		if err != nil {
			return err
		}
	}
	return nil
}

func (d DiscordBotControllerImpl) NotifyCarUpdated(car models.Car) error {
	for _, channel := range d.Channels {
		_, err := d.Session.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Type:        "image",
				Title:       fmt.Sprintf("%v %v has been updated!", car.Brand.Name, car.ModelName),
				Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/cars/%v/%v/%v)", url.PathEscape(car.Brand.Name), url.PathEscape(car.ModelName), car.Year),
				Color:       12590120,
				Image:       &discordgo.MessageEmbedImage{URL: getFavImageUrl(car.Images)},
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
		if err != nil {
			return err
		}
	}
	return nil
}

func (d DiscordBotControllerImpl) NotifyTrackUpdated(track models.Track) error {
	for _, channel := range d.Channels {
		_, err := d.Session.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
			Type:        "image",
			Title:       fmt.Sprintf("%v has been updated!", track.Name),
			Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/tracks/%v/%v/%v)", url.PathEscape(track.Nation.Name), url.PathEscape(track.Name), track.Year),
			Color:       12590120,
			Image:       &discordgo.MessageEmbedImage{URL: getFavImageUrl(track.Images)},
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
		if err != nil {
			return err
		}
	}
	return nil
}

func (d DiscordBotControllerImpl) NotifyTrackAdded(track models.Track) error {
	for _, channel := range d.Channels {
		_, err := d.Session.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
			Type:        "image",
			Title:       fmt.Sprintf("%v has been added to the repository!", track.Name),
			Description: fmt.Sprintf("[Click here for more](https://www.acmodrepository.com/tracks/%v/%v/%v)", url.PathEscape(track.Nation.Name), url.PathEscape(track.Name), track.Year),
			Color:       12590120,
			Image:       &discordgo.MessageEmbedImage{URL: getFavImageUrl(track.Images)},
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
		if err != nil {
			return err
		}
	}
	return nil
}

func getFavImageUrl(images []models.Image) string {
	for _, image := range images {
		if image.Favorite {
			return image.Url[0:27] + ".jpg"
		}
	}
	return ""
}
