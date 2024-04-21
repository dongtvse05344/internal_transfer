package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/internal_transfer/app"
	"github.com/internal_transfer/pb"
	"github.com/internal_transfer/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load configuration: ", err)

	}
	myApp := new(app.App)

	if err := Init(myApp,
		app.WithStore,
	); err != nil {
		panic(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	defer cancel()

	pb.RegisterInternalTransferHandlerServer(ctx, grpcMux, myApp)

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	go func() {
		listener, err := net.Listen("tcp", config.HTTPServerAddress)
		if err != nil {
			log.Fatal("can not create listener", err)
		}
		log.Printf("start http server at %s", listener.Addr().String())
		err = http.Serve(listener, mux)
		if err != nil {
			log.Fatal("can not start http server", err)
		}
	}()

	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := myApp.Close(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}

func Init(app *app.App, initFunc ...app.InitFunc) error {
	for _, fn := range initFunc {
		if err := fn(app); err != nil {
			return err
		}
	}
	return nil
}
