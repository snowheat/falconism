package SDServer

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	SDConfig "github.com/snowheat/falconism/config"
)

//Set ...
func Set(s *fasthttp.Server, router *fasthttprouter.Router) {
	s.Handler = router.Handler
	s.Name = SDConfig.SERVER_NAME
	//s.Concurrency = 500
	s.MaxConnsPerIP = SDConfig.MAX_CONNECTION_PER_IP
	s.DisableHeaderNamesNormalizing = true
	s.MaxRequestsPerConn = SDConfig.MAX_REQUEST_PER_CONNECTION
}

//Run ...
func Run(s *fasthttp.Server) {
	fmt.Println(`
     ______      __                 _
    / ____/___ _/ /________  ____  (_)________ ___
   / /_  / __ '/ / ___/ __ \/ __ \/ / ___/ __ '__ \
  / __/ / /_/ / / /__/ /_/ / / / / (__  ) / / / / /
 /_/    \__,_/_/\___/\____/_/ /_/_/____/_/ /_/ /_/

 *** Ultra-minimalist blog engine optimized for cheap VPS ***

 Powered by Go & Valyala/Fasthttp

 Now listening on: http://localhost:` + SDConfig.PORT + `
 Blog admin: http://localhost:` + SDConfig.PORT + `/admin

 Application started. Press CTRL+C to shut down.
	`)
	log.Fatal(s.ListenAndServe(":" + SDConfig.PORT))
}
