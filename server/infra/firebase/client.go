package firebase

import (
	"context"

	"firebase.google.com/go"
)

var app *firebase.App

func NewApp(ctx context.Context) (*firebase.App, error) {
	if app != nil {
		return app, nil
	}

	var err error
	app, err = firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}

	return app, nil
}
