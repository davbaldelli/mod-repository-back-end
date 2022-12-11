package main

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/v4"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/davide/ModRepository/controllers"
	repo "github.com/davide/ModRepository/repositories/mysql"
	"github.com/davide/ModRepository/routes"
	"github.com/davide/ModRepository/routes/handlers"
	"google.golang.org/api/option"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

type Credentials struct {
	Username string
	Password string
	Host     string
}

type Secret struct {
	Secret       string
	DiscordToken string
	Channels     []string
}

func main() {

	var cred Credentials

	if jsonFile, err := os.Open("credentials.json"); err != nil {
		log.Fatal("no credentials file")
	} else {
		if byteValue, err := ioutil.ReadAll(jsonFile); err != nil {
			log.Fatal("err pasrsing json")
		} else {
			if err := json.Unmarshal(byteValue, &cred); err != nil {
				log.Fatal("err pasrsing json")
			}
		}
	}

	var secret Secret

	if secretFile, err := os.Open("secret.json"); err != nil {
		log.Fatal("no credentials file")
	} else {
		if secretByte, err := ioutil.ReadAll(secretFile); err != nil {
			log.Fatal("err pasrsing json")
		} else {
			if err := json.Unmarshal(secretByte, &secret); err != nil {
				log.Fatal("err pasrsing json")
			}
		}
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(fmt.Errorf("error initializing app: %v", err))
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	dg, err := discordgo.New("Bot " + secret.DiscordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", cred.Username, cred.Password, cred.Host, "mod_repo")
	dbase, err3 := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err3 != nil {
		log.Fatal("error connecting to database")
	}

	carRepo := repo.CarRepositoryImpl{Db: dbase}
	trackRepo := repo.TrackRepositoryImpl{Db: dbase}
	nationRepo := repo.NationsRepositoryImpl{Db: dbase}
	brandRepo := repo.BrandRepositoryImpl{Db: dbase}
	userRepo := repo.UserRepositoryImpl{Db: dbase}
	authorRepo := repo.AuthorsRepositoryImpl{Db: dbase}
	logsRepo := repo.LogRepositoryImpl{Db: dbase}
	serversRepo := repo.ServersRepositoryImpl{Db: dbase}

	web := routes.Web{
		CarHandler: handlers.CarsHandlerImpl{
			CarCtrl:        controllers.CarControllerImpl{Repo: carRepo},
			FirebaseCtrl:   controllers.FirebaseControllerImpl{Client: client, Context: ctx},
			DiscordBotCtrl: controllers.DiscordBotControllerImpl{Session: dg, Channels: secret.Channels},
		},
		TracksHandler: handlers.TrackHandlerImpl{
			TrackCtrl:      controllers.TrackControllerImpl{Repo: trackRepo},
			FirebaseCtrl:   controllers.FirebaseControllerImpl{Client: client, Context: ctx},
			DiscordBotCtrl: controllers.DiscordBotControllerImpl{Session: dg, Channels: secret.Channels},
		},
		NationHandler:   handlers.NationsHandlerImpl{CtrlNations: controllers.NationControllerImpl{Repo: nationRepo}},
		BrandsHandler:   handlers.BrandsHandlerImpl{BrandCtrl: controllers.BrandControllerImpl{Repo: brandRepo}},
		UsersHandler:    handlers.UserHandlerImpl{UserCtrl: controllers.UserControllerImpl{Repo: userRepo}, Secret: secret.Secret},
		AuthorsHandler:  handlers.AuthorHandlerImpl{AuthorsCtrl: controllers.AuthorsControllerImpl{Repo: authorRepo}},
		LogsHandler:     handlers.LogsHandlerImpl{Ctrl: controllers.LogControllerImpl{Repo: logsRepo}},
		ServersHandler:  handlers.ServersHandlerImpl{Ctrl: controllers.ServersControllerImpl{Repo: serversRepo}},
		Middleware:      handlers.MiddlewareImpl{Secret: secret.Secret},
		FirebaseHandler: handlers.FirebaseHandlerImpl{Ctrl: controllers.FirebaseControllerImpl{Client: client, Context: context.Background()}},
	}
	web.Listen()
}
