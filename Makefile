build:
	make -C ./services/service-auth build-local-image
	make -C ./services/service-workspace build-local-image
	make -C ./services/service-notification build-local-image
	make -C ./services/service-audit build-local-image
	make -C ./services/service-meta build-local-image
	make -C ./services/service-data build-local-image