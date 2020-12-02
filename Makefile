SHELL := /bin/bash

run: up seed browse

up:
	docker-compose -f zarf/compose/compose.yaml -f zarf/compose/compose-config.yaml up --detach --remove-orphans

down:
	docker-compose -f zarf/compose/compose.yaml down --remove-orphans

browse:
	python -m webbrowser "http://localhost"

logs:
	docker-compose -f zarf/compose/compose.yaml logs -f
