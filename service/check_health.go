package service

import (
	"context"
	"log"
	"portfolio-go/infra"
)

func CheckHealth(c context.Context) error {
	qry := "SELECT *FROM Users"
	_, err := infra.DB.Query(qry)
	if err != nil {
		log.Println("Error connecting to infra")
	}
	return err
}
