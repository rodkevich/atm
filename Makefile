SONAR_PROJECT_UUID=79eb599c42c7860d70bed9912d59fd222b3dcaf9
SHELL := /bin/bash

psql:
	psql -h localhost -p 5432 -W todos -U admin -E

run:
	docker-compose start

sonar_test_frontend:
	#export PATH='$${PATH}:/home/ioan/Desktop/sonar/bin'; \
#	echo ${PATH}; \
	echo "${PWD}"
	docker run \
			-it --rm --network=host -v "${PWD}"/frontend:/usr/src \
			 sonarsource/sonar-scanner-cli \
			-D sonar.host.url=http://localhost:9000 \
			-D sonar.projectBaseDir=./ \
			-D sonar.sources=. \
			-D sonar.projectName='atm'\
			-D sonar.projectKey='atm'\
			-D sonar.exclusions=**/node_modules/**/* \
			-D sonar.login=$(SONAR_PROJECT_UUID)


sonar_change_admin_pass:
	curl -u admin:admin -X POST "http://localhost:9000/api/users/change_password?login=admin&previousPassword=admin1&password=admin" | bash

sonar_change_admin_001:
	curl 'http://localhost:9000/api/user_tokens/generate' -H --compressed -H 'X-XSRF-TOKEN: gg401offf8d48qgrnnulia3m7d' -H 'Connection: keep-alive' --data-raw 'name=atm' | bash


sonar_test_backend:
	#export PATH='$${PATH}:/home/ioan/Desktop/sonar/bin'; \
#	echo ${PATH}; \
	echo "${PWD}"
	docker run \
			-it --rm --network=host -v "${PWD}"/backend:/usr/src \
			 sonarsource/sonar-scanner-cli \
			-D sonar.host.url=http://localhost:9000 \
			-D sonar.projectBaseDir=./ \
			-D sonar.sources=. \
			-D sonar.projectName='atm'\
			-D sonar.projectKey='atm'\
			-D sonar.login=$(SONAR_PROJECT_UUID)

#proto:
#	@echo "--> Generating gRPC clients"
#	@docker run -v `pwd`/api:/api -v `pwd`/time/goclient:/goclient -v `pwd`/frontend/src/jsclient:/jsclient jfbrandhorst/grpc-web-generators protoc -I /api \
#	  --go_out=plugins=grpc,paths=source_relative:/goclient \
#	  --js_out=import_style=commonjs:/jsclient \
#	  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:/jsclient \
#	 /api/time/v1/time_service.proto

clean_project:
	rm -rf `find . -name __pycache__`
	rm -f `find . -type f -name '*.py[co]' `
	rm -rf `find . -type d -name '*.pytest_cache'`
	rm -rf `find . -type d -name '*.scannerwork'`
	rm -rf `find . -type d -name '*.idea'`
	rm -rf `find . -type d -name '*.mypy_cache'`
	rm -rf `find . -type d -name '*.vscode'`
	rm -rf `find . -type d -name '*.hypothesis'`
	rm -rf `find . -type d -name '.ipynb_checkpoints'`
	rm -f `find . -type f -name '*~' `
	rm -f `find . -type f -name '.*~' `
	rm -f `find . -type f -name '@*' `
	rm -f `find . -type f -name '#*#' `
	rm -f `find . -type f -name '*.orig' `
	rm -f `find . -type f -name '*.rej' `
	rm -f .coverage
	rm -rf coverage
	rm -rf build
	rm -rf htmlcov
	rm -rf dist

#go mod init
#go get -u google.golang.org/grpc
# https://github.com/ktr0731/evans

#go get github.com/ktr0731/evans

# download binary from https://github.com/grpc/grpc-web/releases
#sudo mv ~/Downloads/protoc-gen-grpc-web-1.2.1-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web
#chmod +x /usr/local/bin/protoc-gen-grpc-web

gen:
	@echo "--> Generating gRPC clients"
	protoc -I api/persons/v1 \
		--go_out=,paths=source_relative:persons/goclient \
		--go-grpc_out=,paths=source_relative:persons/ \
		--js_out=import_style=commonjs:frontend/src/jsclient \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:frontend/src/jsclient \
		api/persons/v1/persons.proto

evans:
	#evans --path /path/to --path . --proto proto/city.proto
	evans --path ${PWD}/Desktop/code/src/github.com/rodkevich/atm/api/persons/v1/persons.proto -p 9090
	#evans /home/ioan/Desktop/code/src/github.com/rodkevich/atm/api/persons/v1/persons.proto -p 9090

run-frontend:
	@echo "--> Starting frontend"
	cd ./frontend; \
	npm run serve
