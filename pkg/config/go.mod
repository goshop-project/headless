module goshop.dev/headless/pkg/config

go 1.19

replace (
	darvaza.org/acmefy => ../../../../darvaza.org/acmefy
	darvaza.org/acmefy/pkg/ca => ../../../../darvaza.org/acmefy/pkg/ca
	darvaza.org/cache => ../../../../darvaza.org/cache
	darvaza.org/cache/x/groupcache => ../../../../darvaza.org/cache/x/groupcache
	darvaza.org/cache/x/memcache => ../../../../darvaza.org/cache/x/memcache
	darvaza.org/cache/x/simplelru => ../../../../darvaza.org/cache/x/simplelru
	darvaza.org/core => ../../../../darvaza.org/core
	darvaza.org/darvaza/acme => ../../../../darvaza.org/darvaza/acme
	darvaza.org/darvaza/agent => ../../../../darvaza.org/darvaza/agent
	darvaza.org/darvaza/shared => ../../../../darvaza.org/darvaza/shared
	darvaza.org/darvaza/shared/web => ../../../../darvaza.org/darvaza/shared/web
	darvaza.org/gossipcache => ../../../../darvaza.org/gossipcache
	darvaza.org/middleware => ../../../../darvaza.org/middleware
	goshop.dev/proto/gen => ../../../proto/gen/go
)

require (
	github.com/BurntSushi/toml v1.2.1
	github.com/amery/defaults v0.1.0
	github.com/go-playground/validator/v10 v10.12.0
	go.sancus.dev/config/expand v0.1.0
)

require (
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/leodido/go-urn v1.2.3 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	mvdan.cc/sh/v3 v3.6.0 // indirect
)
