package db

import (
	"fmt"

	"github.com/pczerox/gambituser/models"
	"github.com/pczerox/gambituser/tools"
)

func SignUp(signUp models.SignUp) error {
	fmt.Println("Comienzar registro")

	err := DBConnect()
	if err != nil {
		return err
	}

	//Confirma el cierre de la base de datos
	defer Db.Close()

	statement := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + signUp.UserEmail + "','" + signUp.UserUUID + "','" + tools.FechaMySQL() + "')"

	fmt.Println(statement)
	_, err = Db.Exec(statement)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecución Exitosa")
	return nil
}
