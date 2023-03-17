package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	_ "time/tzdata"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/namsral/flag"
	"github.com/rkojedzinszky/thermo-center/aggregator"
	"github.com/rkojedzinszky/thermo-center/configurator"
	"google.golang.org/grpc"
)

func main() {
	grpcPort := flag.Int("grpc-port", 8079, "Grpc Port")
	dbName := flag.String("dbname", "thermo-center", "Database name")
	dbHost := flag.String("dbhost", "postgres", "Database host to connect to")
	dbUser := flag.String("dbuser", "thermo-center", "Database user")
	dbPassword := flag.String("dbpassword", "thermo-center", "Database password")
	timeZone := flag.String("time-zone", "Europe/Budapest", "Time-zone where application runs")

	flag.Parse()

	db, err := pgxpool.New(context.TODO(),
		fmt.Sprintf(
			"database=%s host=%s user=%s password=%s sslmode=disable pool_max_conns=2",
			*dbName, *dbHost, *dbUser, *dbPassword,
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	loc, err := time.LoadLocation(*timeZone)
	if err != nil {
		log.Println("Loading time-zone:", err)
		loc, _ = time.LoadLocation("")
	}

	grpcServer := grpc.NewServer()

	agg, err := aggregator.NewAggregator(db, loc)
	if err != nil {
		log.Fatal(err)
	}

	aggregator.RegisterAggregatorServer(grpcServer, agg)
	configurator.RegisterConfiguratorServer(grpcServer, configurator.NewConfigurator(db, loc))

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(listen)
}
