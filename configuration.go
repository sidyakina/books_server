package main

import (
	"errors"
	"os"
	"strconv"
)


type config struct {
	reconnect int
	serverPort, pgHost, pgPort, pgPass, pgUser, pgNameDB string
}

func setConfig() (config, error) {
	c := config{}
	r, err := strconv.Atoi(os.Getenv("RECONNECT"))
	if r == 0 || err != nil {
		r = 5
	}
	c.reconnect = r
	var temp string
	temp = os.Getenv("PG_HOST")
	if temp == "" {
		return c, errors.New("empty PG_HOST")
	}
	c.pgHost = temp

	temp = os.Getenv("PG_PORT")
	if temp == "" {
		return c, errors.New("empty PG_PORT")
	}
	c.pgPort = temp
	temp = os.Getenv("PG_PASS")
	if temp == "" {
		return c, errors.New("empty PG_PASS")
	}
	c.pgPass = temp
	temp = os.Getenv("PG_USER")
	if temp == "" {
		return c, errors.New("empty PG_USER")
	}
	c.pgUser = temp
	temp = os.Getenv("PG_DB_NAME")
	if temp == "" {
		return c, errors.New("empty PG_DB_NAME")
	}
	c.pgNameDB = temp

	temp = os.Getenv("SERVER_PORT")
	if temp == "" {
		temp = "3333"
	}
	c.serverPort = temp
	return c, nil
}