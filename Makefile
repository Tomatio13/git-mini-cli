# Windows向け軽量Gitコマンド ビルドファイル

.PHONY: build build-windows build-linux clean test help

# デフォルトターゲット
build: build-windows

# Windowsバイナリビルド（軽量化オプション付き）
build-windows:
	@echo "Windows向けバイナリをビルドしています..."
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o git-cli.exe main.go
	@echo "ビルド完了: git-cli.exe"
	@ls -lh git-cli.exe

# Windows x86 (32bit)バイナリビルド
build-windows-x86:
	@echo "Windows x86向けバイナリをビルドしています..."
	GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o git-cli-x86.exe main.go
	@echo "ビルド完了: git-cli-x86.exe"
	@ls -lh git-cli-x86.exe

# Linuxバイナリビルド（テスト用）
build-linux:
	@echo "Linux向けバイナリをビルドしています..."
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o git-cli main.go
	@echo "ビルド完了: git-cli"
	@ls -lh git-cli

# すべてのプラットフォーム向けビルド
build-all: build-windows build-windows-x86 build-linux

# ビルドファイルのクリーンアップ
clean:
	@echo "ビルドファイルをクリーンアップしています..."
	rm -f git-cli.exe git-cli-x86.exe git-cli

# 簡単なテスト実行
test:
	@echo "基本テストを実行しています..."
	go run main.go help

# ヘルプ表示
help:
	@echo "使用可能なコマンド:"
	@echo "  make build            - Windowsバイナリをビルド"
	@echo "  make build-windows-x86 - Windows x86バイナリをビルド"
	@echo "  make build-linux      - Linuxバイナリをビルド"
	@echo "  make build-all        - すべてのプラットフォーム向けビルド"
	@echo "  make clean        - ビルドファイルをクリーンアップ"
	@echo "  make test         - 基本テストを実行"
	@echo "  make help         - このヘルプを表示"