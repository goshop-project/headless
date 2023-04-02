BINDIR := $(CURDIR)/bin

.PHONY: all build-bin

all: build-bin

build-bin:
	@mkdir -p $(BINDIR)
	@$(MAKE) -C $(CURDIR)/.. OUTDIR=$(BINDIR)
