package frigate

import (
	"HESTIA/internal/database"
	frigateHTTP "HESTIA/internal/frigate/http"
	"HESTIA/internal/models"
	mqttClient "HESTIA/internal/mqtt"

	"gorm.io/gorm/clause"
)

type FrigateClient struct {
	HTTP *frigateHTTP.Client
}

var Client *FrigateClient

func Connect(host string) error {
	httpClient, err := frigateHTTP.NewClient(host, nil)
	if err != nil {
		return err
	}

	Client = &FrigateClient{
		HTTP: httpClient,
	}

	config, err := Client.HTTP.Config.Get()
	if err != nil {
		return err
	}

	loadCameras(config)
	err = initMqttSubs()
	if err != nil {
		return err
	}

	return nil
}

func loadCameras(config *frigateHTTP.ConfigResponse) {
	// Load Cameras into DB & Create Folder Structure
	for _, cam := range config.Cameras {
		camera := models.Cameras{
			FrigateName: cam.Name,
		}

		database.Instance.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "frigate_name"}},
			DoNothing: true,
		}).Create(&camera)
	}
}

func initMqttSubs() error {
	// Init Camera Topics To Follow
	var cameras []models.Cameras
	result := database.Instance.Find(&cameras)
	if result.Error != nil {
		return result.Error
	}

	for _, y := range cameras {
		topic := "frigate/" + y.FrigateName + "/person/snapshot"
		mqttClient.Client.Subscribe(topic, nil)
	}

	mqttClient.Client.Subscribe("frigate/events", nil)

	return nil
}
