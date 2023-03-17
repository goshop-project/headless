module goshop.dev/headless/pkg/auth

go 1.19

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
