/*
Property of Refresh Wardrobe
*/
package main

import (
	"fmt"
	"net"
	"os"

	"github.com/RefreshWardrobe/refresh-mobile-app-service/internal"
	"github.com/RefreshWardrobe/refresh-mobile-app-service/pkg"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	state, err := pkg.NewState(log)
	if err != nil {
		log.Error("unable to generate new state for application")
		os.Exit(1)
	}

	// creating a new grpc servr, insecure for now to allow for http connctions but need to setup keyvault later
	server := grpc.NewServer()
	internal.Register(server)

	l, err := net.Listen("tcp", fmt.Sprint(":%d", 9092))
	if err != nil {
		log.Error("unable to recieve tcp framees and create listener", "error", err)
		os.Exit(1)
	}

	//host server to listen to tcp port packets
	server.Server(1)
}
