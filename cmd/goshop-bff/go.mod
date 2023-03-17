module goshop.dev/headless/cmd/goshop-bff

go 1.19

replace (
	goshop.dev/headless/pkg/config => ../../pkg/config/
	goshop.dev/headless/pkg/server => ../../pkg/server/
)

replace (
	darvaza.org/acmefy => ../../../../darvaza.org/acmefy
	darvaza.org/acmefy/pkg/ca => ../../../../darvaza.org/acmefy/pkg/ca
	darvaza.org/acmefy/pkg/magic => ../../../../darvaza.org/acmefy/pkg/magic
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
	darvaza.org/slog v0.5.1
	github.com/spf13/cobra v1.7.0
	go.sancus.dev/config v0.11.0
	go.sancus.dev/config/flags/cobra v0.1.0
	goshop.dev/headless/pkg/config v0.0.3
	goshop.dev/headless/pkg/server v0.1.1
)

require github.com/BurntSushi/toml v1.2.1 // indirect

require (
	darvaza.org/core v0.9.2 // indirect
	darvaza.org/darvaza/acme v0.1.1 // indirect
	darvaza.org/darvaza/agent v0.2.2 // indirect
	darvaza.org/darvaza/shared v0.5.1 // indirect
	darvaza.org/darvaza/shared/web v0.3.6 // indirect
	darvaza.org/middleware v0.2.2 // indirect
	darvaza.org/slog/handlers/discard v0.4.0 // indirect
	darvaza.org/slog/handlers/filter v0.4.0 // indirect
	darvaza.org/slog/handlers/zerolog v0.4.0 // indirect
	github.com/amery/defaults v0.1.0 // indirect
	github.com/cloudflare/tableflip v1.2.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.12.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/pprof v0.0.0-20230406165453-00490a63f317 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/onsi/ginkgo/v2 v2.9.2 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/qtls-go1-19 v0.3.2 // indirect
	github.com/quic-go/qtls-go1-20 v0.2.2 // indirect
	github.com/quic-go/quic-go v0.33.0 // indirect
	github.com/rs/zerolog v1.29.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	go.sancus.dev/config/expand v0.1.0 // indirect
	go.sancus.dev/core v0.18.2 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.8.0 // indirect
	mvdan.cc/sh/v3 v3.6.0 // indirect
)
