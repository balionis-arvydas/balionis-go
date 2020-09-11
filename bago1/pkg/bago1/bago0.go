package bago1

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/pkg/errors"
	"log"
)

type MyEvent struct {
	Name string `json:"name"`
}

type MyResponse struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, event MyEvent) (MyResponse, error) {

	lc, _ := lambdacontext.FromContext(ctx)

	log.Printf("HandleRequest: event=%v, cognitoId=%v", event, lc.Identity.CognitoIdentityPoolID)

	var res MyResponse
	var err error

	if event.Name != "MyError" {
		res = MyResponse{Message: fmt.Sprintf("Hello %s!", event.Name )}
	} else {
		err = errors.New("")
	}

	log.Printf( "HandleRequest: res=%v, err=%v", res, err)

	return res, err
}