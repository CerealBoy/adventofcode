C = $(shell printf "\033[35;1m-->\033[0m")
V := $(if $V,,@)
DAYFILE := ./year20${YEAR}/day${DAY}

build: ; $(info $(C) building...)
	$V go build .

new-day:
	$(V)if [ "${YEAR}" = "" ] || [ "${DAY}" = "" ]; then\
		echo "Please specify DAY (5) and YEAR (25)";\
		exit 2;\
	fi
	$(V)touch "${DAYFILE}-in" "${DAYFILE}-test"

day: new-day ; $(info $(C) adding new day...)
	$(V)if [ ! -f "${DAYFILE}.go" ]; then\
		sed -e 's/{{YEAR}}/$(YEAR)/g' -e 's/{{DAY}}/$(DAY)/g' day.go.tmpl > "${DAYFILE}.go";\
	fi

year: ; $(info $(C) adding new year...)
	$(V)if [ "${YEAR}" = "" ]; then\
		echo "Please specify YEAR (25)";\
		exit 2;\
	fi
	$(V)mkdir year20$(YEAR)
	$(V)sed 's/{{YEAR}}/$(YEAR)/g' year.go.tmpl > "./year20${YEAR}/cmd.go"

day.exs: new-day ; $(info $(C) adding a new elixir day...)
	$(V)if [ ! -f "${DAYFILE}.exs" ]; then\
		sed -e 's/{{YEAR}}/$(YEAR)/g' -e 's/{{DAY}}/$(DAY)/g' day.exs.tmpl > "${DAYFILE}.exs";\
	fi

.PHONY: build
