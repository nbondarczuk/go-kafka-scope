#
# This file provides all targets needed in work with swagger.
#

swagger/pull:
	@docker pull quay.io/goswagger/swagger

swagger/version:
	@$(SWAGGER_TOOL_PATH)/swagger.sh version

swagger/generate:
	@$(SWAGGER_TOOL_PATH)/swagger.sh generate

swagger/serve:
	@$(SWAGGER_TOOL_PATH)/swagger.sh serve

swagger/help:
	@echo
	@echo 'Swagger utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make swagger/pull'
	@echo '    make swagger/version'
	@echo '    make swagger/generate'
	@echo '    make swagger/serve'
