GO:=$(shell command type go 2> /dev/null)
STRINGER:=$(shell command type stringer 2> /dev/null)

all: check sync-status peer-interface-event peers-interface-event p2p-interface-event

.PHONY: check
check: check/go check/stringer
	@-exit 0

.PHONY: check/go
check/go:
ifndef GO
	$(error "go is not available in PATH")
endif
	@-exit 0

.PHONY: check/stringer
check/stringer:
ifndef STRINGER
	$(error "stringer is not available in PATH. Please run 'make install/stringer'")
endif
	@-exit 0

.PHONY: install/stringer
install/stringer: check/go
	@go get -u golang.org/x/tools/cmd/stringer

.PHONY: sync-status
sync-status:
	@stringer -type=syncStatus && echo "made syncStatus enum"

.PHONY: peer-interface-event
peer-interface-event:
	@stringer -type=peerInterfaceEvent && echo "made peerInterfaceEvent enum"

.PHONY: peers-interface-event
peers-interface-event:
	@stringer -type=peersInterfaceEvent && echo "made peersInterfaceEvent enum"

.PHONY: p2p-interface-event
p2p-interface-event:
	@stringer -type=p2pInterfaceEvent && echo "made p2pInterfaceEvent enum"
