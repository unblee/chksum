RELEASE_BUILD_LDFLAGS = -s -w

.PHONY: test
test:
	go test -v ./...

.PHONY: devel-deps
devel-deps:
	go get github.com/motemen/gobump
	go get github.com/Songmu/goxz
	go get github.com/tcnksm/ghr

.PHONY: crossbuild
crossbuild: devel-deps
	$(eval ver = $(shell gobump show -r))
	goxz -pv=v$(ver) -build-ldflags="$(RELEASE_BUILD_LDFLAGS)" \
	  -d=./dist/v$(ver)

.PHONY: release
release: devel-deps
	_tools/release
	_tools/upload_artifacts
