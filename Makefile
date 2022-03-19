build-image:
	docker build -t ugin -f containers/images/Dockerfile .

run-app-sqlite:
	docker-compose -f containers/composes/dc.sqlite.yml up

clean-app-sqlite:
	docker-compose -f containers/composes/dc.sqlite.yml down
