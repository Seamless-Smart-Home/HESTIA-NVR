package frigate

import (
	"HESTIA/internal/frigate"
	"HESTIA/internal/utils"
)

func ProcessEvent(payload []byte) error {
	// TODO: Check which camera event fired from

	// Get latest image from camera
	snapshot, err := frigate.Client.HTTP.Camera.GetLatestImage("main_bedroom")
	if err != nil {
		return err
	}

	// Process Image
	err = utils.FaceDetectionImgProcessing(snapshot)
	if err != nil {
		return err
	}

	return nil
}
