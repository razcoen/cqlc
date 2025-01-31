package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	logger := slog.Default()

	fs := flag.NewFlagSet("cassandra-healthcheck", flag.ExitOnError)
	timeout := fs.Duration("timeout", time.Minute, "timeout for establishing cassandra session")
	sleep := fs.Duration("sleep", 3*time.Second, "sleep between attempts to establish cassandra session")
	keyspace := fs.String("keyspace", "system", "cassandra keyspace")
	host := fs.String("host", "localhost", "cassandra host")
	if err := fs.Parse(os.Args[1:]); err != nil {
		logger.With("error", err).Error("failed to parse flags")
		os.Exit(1)
	}

	logger = logger.
		With("keyspace", *keyspace).
		With("timeout", *timeout).
		With("sleep", *sleep).
		With("host", *host)
	logger.Info("starting cassandra healthcheck")
	deadline := time.Now().Add(*timeout)
	var err error
	attempt := 0
	for time.Now().Before(deadline) {
		cluster := gocql.NewCluster(*host)
		cluster.Keyspace = *keyspace
		_, err = cluster.CreateSession()
		if err == nil {
			logger.Info("successfuly established cassandra session")
			return
		}
		attempt++
		logger.
			With("error", err).
			With("attempt", attempt).
			With("deadline", fmt.Sprintf("%ds", int(time.Until(deadline).Seconds()))).
			Warn("create casssandra session atttempt failed")
		time.Sleep(*sleep)
	}
	if err == nil {
		logger.Info("successfuly established cassandra session")
	}
	logger.With("error", err).Error("failed to establish cassandra session")
}
