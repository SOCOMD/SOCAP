check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
        $(error Undefined $1$(if $2, ($2))$(if $(value @), \
                required by target `$@`)))

.PHONY: build-win64
build-win64:
	@GOARCH="amd64"; \
	CGO_ENABLED=1; \
	go build -o ./bin/socap_x64.dll -buildmode=c-shared .

.PHONY: build-linux64
build-linux64:
	@GOARCH="amd64"; \
	CGO_ENABLED=1; \
	go build -o ./bin/socap_x64.so -buildmode=c-shared .

.PHONY: build-addon
build-addon:
	@MakePbo.exe -N ./addon ./bin/socap.pbo


# example make DATA_DIR=$(pwd)/tmp/data MAPS_DIR=$(pwd)/tmp/maps start-website 
.PHONY: start-website
start-website: 
	@:$(call check_defined, DATA_DIR)
	@:$(call check_defined, MAPS_DIR)
	@SOCAP_MAPS_DIR=${MAPS_DIR} SOCAP_DATA_DIR=${DATA_DIR} docker-compose up --build

addFile:
	@:$(call check_defined, FILE_NAME)
	@:$(call check_defined, FILE)
	@:$(call check_defined, INFO_NAME)
	@:$(call check_defined, INFO_WORLD)
	@:$(call check_defined, INFO_DURATION)
	curl -vX POST -H "Content-Type: application/json" http://localhost:8080/recieve.php?option=addFile\&fileName=${FILE_NAME} -d @${FILE}
	curl -vX POST http://localhost:8080/recieve.php?option=dbInsert\&worldName=${INFO_WORLD}\&missionName=${INFO_NAME}\&missionDuration=${INFO_DURATION}\&type=coop\&filename=${FILE_NAME}
