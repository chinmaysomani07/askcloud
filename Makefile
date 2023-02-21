TAG_VERSION := $(shell git describe --tags --abbrev=0)

.PHONY: mytarget
mytarget:
	echo "$(TAG_VERSION)"
