// Copyright 2019 Adel Abdelhak.
// Use of this source code is governed by the Apache
// license that can be found in the LICENSE.txt file.

package main

import (
	"flag"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/blakelead/couchbase_exporter/collector"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	cl "github.com/blakelead/confloader"
	log "github.com/sirupsen/logrus"
)

var (
	version = "devel"

	serverListenAddress string
	serverMetricsPath   string
	serverTimeout       time.Duration
	dbUsername          string
	dbPassword          string
	dbURI               string
	dbTimeout           time.Duration
	tlsEnabled          bool
	tlsSkipInsecure     bool
	tlsCACert           string
	tlsClientCert       string
	tlsClientKey        string
	logLevel            string
	logFormat           string
	scrapeCluster       bool
	scrapeNode          bool
	scrapeBucket        bool
	scrapeXDCR          bool
	configFile          string
)

func main() {

	// Initialize global variables, initialize logger and display user defined values.
	initEnv()
	initLogger()
	displayInfo()

	// Context encapsulates connection and scraping details for the exporters.
	context := collector.Context{
		URI:             dbURI,
		Username:        dbUsername,
		Password:        dbPassword,
		Timeout:         dbTimeout,
		TLSEnabled:      tlsEnabled,
		TLSSkipInsecure: tlsSkipInsecure,
		TLSCACert:       tlsCACert,
		TLSClientCert:   tlsClientCert,
		TLSClientKey:    tlsClientKey,
		ScrapeCluster:   scrapeCluster,
		ScrapeNode:      scrapeNode,
		ScrapeBucket:    scrapeBucket,
		ScrapeXDCR:      scrapeXDCR,
	}

	// Exporters are initialized, meaning that metrics files are loaded and
	// Exporter objects are created and filled with metrics metadata.
	collector.InitExporters(context)

	// Handle metrics path with the prometheus handler.
	http.Handle(serverMetricsPath, promhttp.Handler())

	// Handle paths other than given metrics path.
	if serverMetricsPath != "/" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`<body><p>See <a href="` + serverMetricsPath + `">Metrics</a></p></body>`))
		})
	}

	// Create the server with provided listen address and server timeout.
	s := &http.Server{
		Addr:        serverListenAddress,
		ReadTimeout: serverTimeout,
	}

	// Start the server.
	log.Info("Started listening at ", serverListenAddress)
	log.Fatal(s.ListenAndServe())
}

func initEnv() {
	// Default parameters.
	serverListenAddress = "127.0.0.1:9191"
	serverMetricsPath = "/metrics"
	serverTimeout = 10 * time.Second
	dbURI = "http://127.0.0.1:8091"
	dbTimeout = 10 * time.Second
	tlsEnabled = false
	tlsSkipInsecure = false
	tlsCACert = ""
	tlsClientCert = ""
	tlsClientKey = ""
	logLevel = "info"
	logFormat = "text"
	scrapeCluster = true
	scrapeNode = true
	scrapeBucket = true
	scrapeXDCR = true

	// Get configuration file values if a configuration file is found (either json or yml).
	configFile = "config.json"
	config, err := cl.Load(configFile)
	if err != nil {
		configFile = "config.yml"
		config, err = cl.Load(configFile)
		if err != nil {
			configFile = ""
		}
	}
	if err == nil {
		serverListenAddress = config.GetString("web.listenAddress")
		serverMetricsPath = config.GetString("web.telemetryPath")
		serverTimeout = config.GetDuration("web.timeout")
		dbUsername = config.GetString("db.user")
		dbPassword = config.GetString("db.password")
		dbURI = config.GetString("db.uri")
		dbTimeout = config.GetDuration("db.timeout")
		tlsEnabled = config.GetBool("db.tls.enabled")
		tlsSkipInsecure = config.GetBool("db.tls.skipInsecure")
		tlsCACert = config.GetString("db.tls.caCert")
		tlsClientCert = config.GetString("db.tls.clientCert")
		tlsClientKey = config.GetString("db.tls.clientKey")
		logLevel = config.GetString("log.level")
		logFormat = config.GetString("log.format")
		scrapeCluster = config.GetBool("scrape.cluster")
		scrapeNode = config.GetBool("scrape.node")
		scrapeBucket = config.GetBool("scrape.bucket")
		scrapeXDCR = config.GetBool("scrape.xdcr")
	}

	// Get environment variables values.
	if val, ok := os.LookupEnv("CB_EXPORTER_LISTEN_ADDR"); ok {
		serverListenAddress = val
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_TELEMETRY_PATH"); ok {
		serverMetricsPath = val
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_SERVER_TIMEOUT"); ok {
		serverTimeout, _ = time.ParseDuration(val)
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_DB_USER"); ok {
		dbUsername = val
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_DB_PASSWORD"); ok {
		dbPassword = val
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_DB_URI"); ok {
		dbURI = val
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_DB_TIMEOUT"); ok {
		dbTimeout, _ = time.ParseDuration(val)
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_TLS_ENABLED"); ok {
		tlsEnabled, _ = strconv.ParseBool(val)
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_TLS_SKIP_INSECURE"); ok {
		tlsSkipInsecure, _ = strconv.ParseBool(val)
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_TLS_CA_CERT"); ok {
		tlsCACert = val
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_TLS_CLIENT_CERT"); ok {
		tlsClientCert = val
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_TLS_CLIENT_KEY"); ok {
		tlsClientKey = val
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_LOG_LEVEL"); ok {
		logLevel = val
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_LOG_FORMAT"); ok {
		logFormat = val
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_SCRAPE_CLUSTER"); ok {
		scrapeCluster, _ = strconv.ParseBool(val)
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_SCRAPE_NODE"); ok {
		scrapeNode, _ = strconv.ParseBool(val)
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_SCRAPE_BUCKET"); ok {
		scrapeBucket, _ = strconv.ParseBool(val)
	}
	if val, ok := os.LookupEnv("CB_EXPORTER_SCRAPE_XDCR"); ok {
		scrapeXDCR, _ = strconv.ParseBool(val)
	}

	// Get command-line values.
	flag.StringVar(&serverListenAddress, "web.listen-address", serverListenAddress, "Address to listen on for HTTP requests.")
	flag.StringVar(&serverMetricsPath, "web.telemetry-path", serverMetricsPath, "Path under which to expose metrics.")
	flag.DurationVar(&serverTimeout, "web.timeout", serverTimeout, "Server read timeout in seconds.")
	flag.StringVar(&dbURI, "db.uri", dbURI, "Couchbase node URI with port.")
	flag.DurationVar(&dbTimeout, "db.timeout", dbTimeout, "Couchbase client timeout in seconds.")
	flag.BoolVar(&tlsEnabled, "tls.enabled", tlsEnabled, "If true, all TLS settings must be set.")
	flag.BoolVar(&tlsSkipInsecure, "tls.skip-insecure", tlsSkipInsecure, "If true, self-signed certificate won't be verified.")
	flag.StringVar(&tlsCACert, "tls.ca-cert", tlsCACert, "Path to the root certificate.")
	flag.StringVar(&tlsClientCert, "tls.client-cert", tlsClientCert, "Path to the client certificate.")
	flag.StringVar(&tlsClientKey, "tls.client-key", tlsClientKey, "Path to the client key.")
	flag.StringVar(&logLevel, "log.level", logLevel, "Log level: info, debug, warn, error, fatal.")
	flag.StringVar(&logFormat, "log.format", logFormat, "Log format: text or json.")
	flag.BoolVar(&scrapeCluster, "scrape.cluster", scrapeCluster, "If false, cluster metrics won't be scraped.")
	flag.BoolVar(&scrapeNode, "scrape.node", scrapeNode, "If false, node metrics won't be scraped.")
	flag.BoolVar(&scrapeBucket, "scrape.bucket", scrapeBucket, "If false, bucket metrics won't be scraped.")
	flag.BoolVar(&scrapeXDCR, "scrape.xdcr", scrapeXDCR, "If false, XDCR metrics won't be scraped.")
	flag.Parse()
}

func initLogger() {
	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	switch logFormat {
	case "json":
		log.SetFormatter(&log.JSONFormatter{
			PrettyPrint: true,
		})
	default:
		log.SetFormatter(&log.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	}
	log.SetOutput(os.Stdout)
}

func displayInfo() {
	log.Info("Couchbase Exporter Version: ", version)
	log.Info("config.file=", configFile)
	log.Info("web.listen-address=", serverListenAddress)
	log.Info("web.telemetry-path=", serverMetricsPath)
	log.Info("web.timeout=", serverTimeout)
	log.Info("db.uri=", dbURI)
	log.Info("db.timeout=", dbTimeout)
	log.Info("tls.enabled=", tlsEnabled)
	log.Info("tls.skip-insecure=", tlsSkipInsecure)
	log.Info("tls.ca-cert=", tlsCACert)
	log.Info("tls.client-cert=", tlsClientCert)
	log.Info("tls.client-key=", tlsClientKey)
	log.Info("log.level=", logLevel)
	log.Info("log.format=", logFormat)
	log.Info("scrape.cluster=", scrapeCluster)
	log.Info("scrape.node=", scrapeNode)
	log.Info("scrape.bucket=", scrapeBucket)
	log.Info("scrape.xdcr=", scrapeXDCR)
}
