#!make

## run: Запускает приложение в дебаг режиме (с флагом -race)
run:
	@bash -c "go run -race *.go $(ARGS)"

## build: Сбилдит бинарник
build:
	@bash -c "go build -o bin *.go"

help: Makefile
	@echo " > Список команд по "$(PROJECT_NAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
