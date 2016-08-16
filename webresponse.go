package main

import (
	"strconv"
	"time"

	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
	"math"
	"net"
	"net/http"
	"os"
)

var (
	// Version is the version of the software
	Version string
	// BuildStmp is the build date
	BuildStmp string
	// GitHash is the git build hash
	GitHash string
)

var (
	// command line parameters
	port  = 8020
	path  = "/"
	count = 1
)

func main() {
	// setup time ref to UTC
	time.Local = time.UTC

	// create cli app
	cliApp := cli.NewApp()
	cliApp.EnableBashCompletion = true
	cliApp.Usage = ""
	timeStmp, err := strconv.Atoi(BuildStmp)
	if err != nil {
		timeStmp = 0
	}
	cliApp.Version = Version + ", build on " + time.Unix(int64(timeStmp), 0).String() + ", git hash " + GitHash
	cliApp.Name = "webresponse"
	cliApp.Authors = []cli.Author{cli.Author{Name: "sfr"}}
	cliApp.Copyright = "Sebastienfr " + strconv.Itoa(time.Now().Year())

	// define flags
	cliApp.Flags = []cli.Flag{
		cli.IntFlag{
			Value: port,
			Name:  "port",
			Usage: "Set the listening port of the webserver",
		},
		cli.StringFlag{
			Value: path,
			Name:  "path",
			Usage: "Set the the path the server listens to",
		},
	}

	// define main loop
	cliApp.Action = func(c *cli.Context) error {
		// parse parameters
		port = c.Int("port")
		path = c.String("path")

		// logger
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(os.Stdout)
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, ForceColors: true})

		fmt.Print("* --------------------------------------------------- *\n")
		fmt.Printf("|   port : %d\n", port)
		fmt.Printf("|   path : %s\n", path)
		fmt.Print("* --------------------------------------------------- *\n")

		// create a negroni web server
		n := negroni.New()

		// add middleware for logging
		n.Use(negronilogrus.NewMiddlewareFromLogger(logrus.StandardLogger(), "webresponse"))

		// add recovery middleware in case of panic in handler func
		recovery := negroni.NewRecovery()
		recovery.PrintStack = false
		n.Use(recovery)

		router := mux.NewRouter()

		router.PathPrefix(path).Subrouter().Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete).
			HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				// count requests
				if count++; count >= math.MaxInt32 {
					count = 1
				}

				// get hostname
				host, err := os.Hostname()
				if err != nil {
					host = "unknown"
				}

				// get ips
				var ips []string
				ifaces, err := net.Interfaces()
				for _, iface := range ifaces {
					addrs, err := iface.Addrs()
					if err != nil {
						continue
					}
					for _, addr := range addrs {
						ips = append(ips, addr.String())
					}
				}

				// build response
				response := map[string]interface{}{
					"host":   host,
					"ips":    ips,
					"count":  count,
					"header": r.Header,
					"url":    r.URL,
				}

				logrus.WithField("request", r).WithField("response", response).Debug("handling incoming request")

				// send response
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(response)
			})

		// use the router
		n.UseHandler(router)

		// run the server
		n.Run(fmt.Sprintf(":%d", port))

		return nil
	}

	// start the cliApp
	err = cliApp.Run(os.Args)

	if err != nil {
		logrus.WithField("error", err).Fatal("error running webresponse")
	}
}
