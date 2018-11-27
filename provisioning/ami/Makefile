export AWS_PROFILE        := saxsir
export AWS_DEFAULT_REGION := ap-northeast-1

PACKER_VERSION := 1.3.2

build: packer.json
	docker run --rm -it \
		-v ~/.aws:/root/.aws \
		-e TZ=Asia/Tokyo \
		-e AWS_PROFILE=$(AWS_PROFILE) \
		-e AWS_DEFAULT_REGION=$(AWS_DEFAULT_REGION) \
		-v $(CURDIR):/work \
		-w /work \
		hashicorp/packer:$(PACKER_VERSION) build --only=amazon-ebs $<
