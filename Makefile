# Usage: TBD

.PHONY: image
image:
	@docker build -f cmd/kimbap-server/Dockerfile --target bin .

.PHONY: lint
lint:
	@docker build -f cmd/kimbap-server/Dockerfile --target lint .