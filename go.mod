module go-my-manager-api

go 1.12

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190712062909-fae7ac547cb7

replace golang.org/x/net => github.com/golang/net v0.0.0-20190628185345-da137c7871d7

replace golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190717194535-128ec6dfca09

replace golang.org/x/text => github.com/golang/text v0.3.2

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4

replace golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45

replace golang.org/x/image => github.com/golang/image v0.0.0-20190703141733-d6a02ce849c9

replace golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190711165009-e47acb2ca7f9

replace golang.org/x/exp => github.com/golang/exp v0.0.0-20190627132806-fd42eb6b336f

replace cloud.google.com/go => github.com/googleapis/google-cloud-go v0.41.0

replace google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190716160619-c506a9f90610

replace google.golang.org/appengine => github.com/golang/appengine v1.6.1

replace google.golang.org/api => github.com/googleapis/google-api-go-client v0.7.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/jinzhu/gorm v1.9.10
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.2.9 // indirect
)
