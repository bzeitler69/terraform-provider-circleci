package circlecigo

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type Client struct {
	sling                *sling.Sling
	Projects             *ProjectService
	SshKeys              *SshKeyService
	EnvironmentVariables *EnvironmentVariableService
}

func NewClient(authToken string) *Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	base := sling.New().Client(client).Base("https://circleci.com/api/v1.1/").Set("Content-Type", "application/json")

	return &Client{
		sling:                base,
		Projects:             newProjectService(base, authToken),
		SshKeys:              newSshKeyService(base, authToken),
		EnvironmentVariables: newEnvironmentVariableService(base, authToken),
	}
}
