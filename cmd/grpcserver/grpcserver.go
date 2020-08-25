package main

import (
	"fmt"
	"log"
	"net"
	"time"

	_ "time/tzdata"

	"github.com/namsral/flag"
	"github.com/rkojedzinszky/thermo-center/aggregator"
	"github.com/rkojedzinszky/thermo-center/configurator"
	"google.golang.org/grpc"

	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	grpcPort := flag.Int("grpc-port", 8079, "Grpc Port")
	dbName := flag.String("dbname", "thermo-center", "Database name")
	dbHost := flag.String("dbhost", "postgres", "Database host to connect to")
	dbUser := flag.String("dbuser", "thermo-center", "Database user")
	dbPassword := flag.String("dbpassword", "thermo-center", "Database password")
	timeZone := flag.String("time-zone", "Europe/Budapest", "Time-zone where application runs")

	flag.Parse()

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"database=%s host=%s user=%s password=%s sslmode=disable",
			*dbName, *dbHost, *dbUser, *dbPassword,
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(2)

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
