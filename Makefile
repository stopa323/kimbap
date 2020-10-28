# Usage: TBD

opts = DOCKER_BUILDKIT=1

.PHONY: image
image:
	@$(opts) docker build -f cmd/kimbap-server/Dockerfile --target bin .

.PHONY: lint
lint:
	@$(opts) docker build -f cmd/kimbap-server/Dockerfile --target lint .

.PHONY: unit-tests
unit-tests:
	@$(opts) docker build -f cmd/kimbap-server/Dockerfile --target unit-tests .