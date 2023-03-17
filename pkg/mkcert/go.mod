module goshop.dev/headless/pkg/mkcert

go 1.19

replace goshop.dev/headless/pkg/config => ../config

replace (
	darvaza.org/acmefy => ../../../../darvaza.org/acmefy
	darvaza.org/acmefy/pkg/acme => ../../../../darvaza.org/acmefy/pkg/acme
	darvaza.org/acmefy/pkg/ca => ../../../../darvaza.org/acmefy/pkg/ca
	darvaza.org/acmefy/pkg/client => ../../../../darvaza.org/acmefy/pkg/client
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
	darvaza.org/acmefy/pkg/ca v0.1.0
	darvaza.org/core v0.9.2
	darvaza.org/darvaza/shared v0.5.1
	goshop.dev/headless/pkg/config v0.0.3
)

require (
	darvaza.org/acmefy v0.4.2 // indirect
	darvaza.org/acmefy/pkg/respond v0.1.0 // indirect
	darvaza.org/darvaza/shared/web v0.3.6 // indirect
	darvaza.org/slog v0.5.1 // indirect
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/amery/defaults v0.1.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.12.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.3 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	go.sancus.dev/config/expand v0.1.0 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	mvdan.cc/sh/v3 v3.6.0 // indirect
)
