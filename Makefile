build:
	make -C ./services/service-auth build-local-image
	make -C ./services/service-workspace build-local-image
	make -C ./services/service-notifications build-local-image
	make -C ./services/service-audit build-local-image