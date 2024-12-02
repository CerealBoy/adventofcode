C = $(shell printf "\033[35;1m-->\033[0m")
V := $(if $V,,@)
DAYFILE := ./year20${YEAR}/day${DAY}

build: ; $(info $(C) building...)
	$V go build .

day: ; $(info $(C) adding new day...)
	$(V)if [ "${YEAR}" = "" ] || [ "${DAY}" = "" ]; then\
		echo "Please specify DAY (5) and YEAR (25)";\
		exit 2;\
	fi
	$(V)if [ -f "${DAYFILE}-in" ] || [ -f "${DAYFILE}-test" ] || [ -f "${DAYFILE}.go" ]; then\
		echo "Files already exist!";\
		exit 3;\
	fi
	$(V)touch "${DAYFILE}-in" "${DAYFILE}-test"
	$(V)sed -e 's/{{YEAR}}/$(YEAR)/g' -e 's/{{DAY}}/$(DAY)/g' day.go.tmpl > "${DAYFILE}.go"

year: ; $(info $(C) adding new year...)
	$(V)if [ "${YEAR}" = "" ]; then\
		echo "Please specify YEAR (25)";\
		exit 2;\
	fi
	$(V)mkdir year20$(YEAR)
	$(V)sed 's/{{YEAR}}/$(YEAR)/g' year.go.tmpl > "./year20${YEAR}/cmd.go"

