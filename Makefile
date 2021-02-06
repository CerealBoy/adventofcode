C = $(shell printf "\033[35;1m-->\033[0m")
V := $(if $V,,@)

build: ; $(info $(C) building...)
	$V go build .

