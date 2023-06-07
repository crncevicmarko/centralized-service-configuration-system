package configstore

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	idempotencyKeys = "idempotencyKeys/%s"
	configs         = "configs/%s/%s"
	configsLabels   = "configs/%s/%s/%s"
	all             = "configs"
)

func createNewIdempotencyKey(id string) string {
	return fmt.Sprintf(idempotencyKeys, id)
}

func generateKey(version string, labels string) (string, string) {
	id := uuid.New().String()
	if labels != "" {
		return fmt.Sprintf(configsLabels, id, version, labels), id
	} else {
		return fmt.Sprintf(configs, id, version), id
	}

}

func constructKey(id string, version string, labels string) string {
	if labels != "" {
		return fmt.Sprintf(configsLabels, id, version, labels)
	} else {
		return fmt.Sprintf(configs, id, version)
	}

}

func constructKey1(group_id string, group_version string) string {
	return fmt.Sprintf(configs, group_id, group_version)
}
