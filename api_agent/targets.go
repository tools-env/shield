package api_agent

import (
	"bytes"
	"fmt"
	"github.com/starkandwayne/shield/db"
)

func FetchTargetsList(plugin, unused string) (*[]db.AnnotatedTarget, error) {

	// Data to be returned of proper type
	data := &[]db.AnnotatedTarget{}

	// Make uri based on options
	uri := fmt.Sprintf("v1/targets")
	joiner := "?"
	if plugin != "" {
		uri = fmt.Sprintf("%s%splugin=%s", uri, joiner, plugin)
		joiner = "&"
	}
	if unused != "" {
		uri = fmt.Sprintf("%s%sunused=%s", uri, joiner, unused)
		joiner = "&"
	}

	// Call generic API request
	err := makeApiCall(data, `GET`, uri, nil)
	return data, err
}

func GetTarget(uuid string) (*db.AnnotatedTarget, error) {

	// Data to be returned of proper type
	data := &db.AnnotatedTarget{}

	// Make uri based on options
	uri := fmt.Sprintf("v1/target/%s", uuid)

	// Call generic API request
	err := makeApiCall(data, `GET`, uri, nil)
	return data, err
}

func CreateTarget(content string) (*db.AnnotatedTarget, error) {

	data := struct {
		Status string `json:"ok"`
		UUID   string `json:"uuid"`
	}{}

	buf := bytes.NewBufferString(content)

	err := makeApiCall(&data, `POST`, `v1/targets`, buf)

	if err == nil {
		return GetTarget(data.UUID)
	}
	return nil, err
}

func UpdateTarget(uuid string, content string) (*db.AnnotatedTarget, error) {

	data := struct {
		Status string `json:"ok"`
	}{}

	buf := bytes.NewBufferString(content)

	uri := fmt.Sprintf("v1/target/%s", uuid)

	err := makeApiCall(&data, `PUT`, uri, buf)

	if err == nil {
		return GetTarget(uuid)
	}
	return nil, err
}

func DeleteTarget(uuid string) error {

	uri := fmt.Sprintf("v1/target/%s", uuid)

	data := struct {
		Status string `json:"ok"`
	}{}

	err := makeApiCall(&data, `DELETE`, uri, nil)

	return err
}
