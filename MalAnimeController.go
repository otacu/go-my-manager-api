package main

import (
	"context"
	"github.com/labstack/echo"
	animeProto "go-my-manager-api/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

func Search(c echo.Context) error {
	req := new(MalAnimeSearchRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if req.Name == "" || req.Text == "" {
		return c.JSON(http.StatusBadRequest, "param error")
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(malAnimeServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return c.JSON(http.StatusBadRequest, "connect service error")
	}
	defer conn.Close()
	client := animeProto.NewAnimeProtoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Search(ctx, &animeProto.AnimeSearchRequest{Name: req.Name, Text: req.Text})
	if err != nil {
		log.Fatalf("error when search mal anime: %v", err)
		return c.JSON(http.StatusBadRequest, "error when search mal anime")
	}
	log.Printf("search result: %s", resp)

	return c.JSON(http.StatusOK, resp)
}
