package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/pczerox/gambituser/awsgo"
	"github.com/pczerox/gambituser/db"
	"github.com/pczerox/gambituser/models"
)

func main() {
	fmt.Println("Iniciar proyecto gambitUser")

	lambda.Start(startLambda)
}

func startLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.StartAWS()

	if !ValidParameters() {
		fmt.Println("Error en los parámetros. Debe enviar 'SecretName'")
		err := errors.New("error en los parámetros debe enviar SecretName")

		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el Secret" + err.Error())
		return event, err
	}

	err = db.SignUp(data)
	return event, err
}

func ValidParameters() bool {
	var getParameter bool
	_, getParameter = os.LookupEnv("SecretName")

	return getParameter
}
