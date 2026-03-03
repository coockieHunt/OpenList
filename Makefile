APP_NAME := Openlist

LDFLAGS := -s -w

TAGS := netgo osusergo static_build

.PHONY: all build compress

all:  build  compress

build:
	@echo "Compil start $(APP_NAME)..."
	CGO_ENABLED=0 go build -trimpath -tags "$(TAGS)" -ldflags="$(LDFLAGS)" -o $(APP_NAME) .
	@echo "[info] Init size $$(ls -lh $(APP_NAME) | awk '{print $$5}')"

compress:
	@if command -v upx >/dev/null 2>&1; then \
		echo "Run UPX (Ultra Brute)..."; \
		upx --ultra-brute --lzma $(APP_NAME); \
		echo "[info] final size : $$(ls -lh $(APP_NAME) | awk '{print $$5}')"; \
	else \
		echo "[error] UPX not found."; \
	fi

