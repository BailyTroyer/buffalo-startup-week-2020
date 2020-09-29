# ============================================================================
# Globals
# ============================================================================
S3_URL:=s3://presentation.bflobox.com

.PHONY: build-static-assets
build-static-assets:
	yarn install
	yarn run build

.PHONY: push-static-assets
push-static-assets:
	aws s3 sync public $(S3_URL) --cache-control max-age=30