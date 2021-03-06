check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
        $(error Undefined $1$(if $2, ($2))$(if $(value @), \
                required by target `$@`)))

.PHONY: build-win64
build-win64:
	@cd go; \
	GOARCH="amd64" CGO_ENABLED=1 go build -o ../out/socap_x64.dll -buildmode=c-shared .

.PHONY: build-linux64
build-linux64:
	@cd go; \
	GOARCH="amd64" CGO_ENABLED=1 go build -o ../out/socap_x64.so -buildmode=c-shared .; \
	cp -f ../out/socap_x64.so /srv/games/servers/arma3_mods/SOCOMD_Core/@socap/socap_x64.so;

.PHONY: build-linux32
build-linux32:
	@cd go; \
	GOARCH=386	CGO_ENABLED=1 go build -o ../out/socap.so -buildmode=c-shared .;\
	cp -f ../out/socap.so /srv/games/servers/arma3_mods/SOCOMD_Core/@socap/socap.so;

.PHONY: build-addon
build-addon:
	@MakePbo.exe -N ./addon ./out/socap.pbo

.PHONY: build-addon-server
build-addon-server:
	@$$HOME/.local/bin/mikero/makepbo -N ./addon ./out/socap.pbo; \
	cp -f ./out/socap.pbo /srv/games/servers/arma3_mods/SOCOMD_Core/@socap/addons/socap.pbo;

# example make DATA_DIR=$(pwd)/tmp/data MAPS_DIR=$(pwd)/tmp/maps start-website 
.PHONY: start-website
start-website: 
	@:$(call check_defined, DATA_DIR)
	@:$(call check_defined, MAPS_DIR)
	@SOCAP_MAPS_DIR=${MAPS_DIR} SOCAP_DATA_DIR=${DATA_DIR} docker-compose up --build

.PHONY: start-website-prod
start-website-prod:
	@SOCAP_MAPS_DIR=/srv/socap/maps SOCAP_DATA_DIR=/srv/socap/data docker-compose up --build -d

addFile:
	@:$(call check_defined, FILE_NAME)
	@:$(call check_defined, FILE)
	@:$(call check_defined, INFO_NAME)
	@:$(call check_defined, INFO_WORLD)
	@:$(call check_defined, INFO_DURATION)
	curl -vX POST -H "Content-Type: application/json" http://localhost:9000/recieve.php?option=addFile\&fileName=${FILE_NAME} -d @${FILE}
	curl -vX POST http://localhost:9000/recieve.php?option=dbInsert\&worldName=${INFO_WORLD}\&missionName=${INFO_NAME}\&missionDuration=${INFO_DURATION}\&type=coop\&filename=${FILE_NAME}

