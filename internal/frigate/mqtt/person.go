package frigate

import (
	"HESTIA/internal/utils"
	"bytes"
	"image"
)

func ProcessPersonDetected(payload []byte) error {
	//Convert MQTT MSG to Image
	snapshot, _, err := image.Decode(bytes.NewReader(payload))
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
