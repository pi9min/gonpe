package firebase

import (
	"context"

	"firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App

func NewApp(ctx context.Context, serviceAccountKey []byte) (*firebase.App, error) {
	if app != nil {
		return app, nil
	}

	var err error
	o := option.WithCredentialsJSON(serviceAccountKey)
	app, err = firebase.NewApp(ctx, nil, o)
	if err != nil {
		return nil, err
	}

	return app, nil
}
