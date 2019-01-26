GO:=$(shell command type go 2> /dev/null)
STRINGER:=$(shell command type stringer 2> /dev/null)

all: check sync-event-enum peer-event-enum peers-event-enum p2p-event-enum handler-func-enum

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

.PHONY: sync-event-enum
sync-event-enum:
	@(cd ./sync/types && stringer -type=eventEnum && echo "made p2p syncEvent enum")

.PHONY: peer-event-enum
peer-event-enum:
	@(cd ./peer && stringer -type=eventEnum && echo "made peer event enum")

.PHONY: peers-event-enum
peers-event-enum:
	@(cd ./peers && stringer -type=eventEnum && echo "made peers event enum")

.PHONY: handler-func-enum
handler-func-enum:
	@(cd ./handler && stringer -type=funcEnum && echo "made handler func enum")

.PHONY: p2p-event-enum
p2p-event-enum:
	@stringer -type=eventEnum && echo "made p2p event enum"
