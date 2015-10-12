// skipper program main
//
// for a summary about skipper, please see the readme file.
package main

import (
	"flag"
	"github.com/zalando/skipper/run"
	"log"
	"strings"
	"time"
)

const (
	defaultAddress           = ":9090"
	defaultEtcdUrls          = "http://127.0.0.1:2379,http://127.0.0.1:4001"
	defaultStorageRoot       = "/skipper"
	defaultSourcePollTimeout = 180

	addressUsage                   = "address where skipper should listen on"
	etcdUrlsUsage                  = "urls where etcd can be found"
	insecureUsage                  = "set this flag to allow invalid certificates for tls connections"
	storageRootUsage               = "prefix for skipper related data in the provided etcd storage"
	innkeeperUrlUsage              = "url of the innkeeper API"
	sourcePollTimeoutUsage         = "polling timeout of the routing data sources, in milliseconds"
	oauthUrlUsage                  = "OAuth2 URL for Innkeeper authentication"
    oauthScopeUsage = "OAuth2 scope to access route definitions in Innkeeper"
    oauthCredentialsDirUsage = "directory where oauth credentials are stored: client.json and user.json"
	routesFileUsage                = "routes file to use instead of etcd"
	innkeeperAuthTokenUsage        = "fixed token for innkeeper authentication"
	innkeeperPreRouteFiltersUsage  = "global pre-route filters for routes from Innkeeper"
	innkeeperPostRouteFiltersUsage = "global post-route filters for routes from Innkeeper"
	devModeUsage                   = "enables developer time behavior, like ubuffered routing updates"
)

var (
	address                   string
	etcdUrls                  string
	insecure                  bool
	storageRoot               string
	innkeeperUrl              string
	sourcePollTimeout         int
	routesFile                string
	oauthUrl                  string
    oauthScope string
    oauthCredentialsDir string
	innkeeperAuthToken        string
	innkeeperPreRouteFilters  string
	innkeeperPostRouteFilters string
	devMode                   bool
)

func init() {
	flag.StringVar(&address, "address", defaultAddress, addressUsage)
	flag.StringVar(&etcdUrls, "etcd-urls", defaultEtcdUrls, etcdUrlsUsage)
	flag.BoolVar(&insecure, "insecure", false, insecureUsage)
	flag.StringVar(&storageRoot, "storage-root", defaultStorageRoot, storageRootUsage)
	flag.StringVar(&innkeeperUrl, "innkeeper-url", "", innkeeperUrlUsage)
	flag.IntVar(&sourcePollTimeout, "source-poll-timeout", defaultSourcePollTimeout, sourcePollTimeoutUsage)
	flag.StringVar(&routesFile, "routes-file", "", routesFileUsage)
	flag.StringVar(&oauthUrl, "oauth-url", "", oauthUrlUsage)
	flag.StringVar(&oauthScope, "oauth-scope", "", oauthScopeUsage)
	flag.StringVar(&oauthCredentialsDir, "oauth-credentials-dir", "", oauthCredentialsDirUsage)
	flag.StringVar(&innkeeperAuthToken, "innkeeper-auth-token", "", innkeeperAuthTokenUsage)
	flag.StringVar(&innkeeperPreRouteFilters, "innkeeper-pre-route-filters", "", innkeeperPreRouteFiltersUsage)
	flag.StringVar(&innkeeperPostRouteFilters, "innkeeper-post-route-filters", "", innkeeperPostRouteFiltersUsage)
	flag.BoolVar(&devMode, "dev-mode", false, devModeUsage)
	flag.Parse()
}

func main() {
	log.Fatal(run.Run(run.Options{
		Address:                   address,
		EtcdUrls:                  strings.Split(etcdUrls, ","),
		StorageRoot:               storageRoot,
		Insecure:                  insecure,
		InnkeeperUrl:              innkeeperUrl,
		SourcePollTimeout:         time.Duration(sourcePollTimeout) * time.Millisecond,
		RoutesFile:            routesFile,
		IgnoreTrailingSlash:       false,
		OAuthUrl:                  oauthUrl,
        OAuthScope: oauthScope,
        OAuthCredentialsDir: oauthCredentialsDir,
		InnkeeperAuthToken:        innkeeperAuthToken,
		InnkeeperPreRouteFilters:  innkeeperPreRouteFilters,
		InnkeeperPostRouteFilters: innkeeperPostRouteFilters,
		DevMode:                   devMode}))
}
