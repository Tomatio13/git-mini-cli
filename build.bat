@echo off
echo Windows向け軽量Gitコマンドをビルドしています...

REM Windowsバイナリをビルド
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o git-cli.exe main.go

if %ERRORLEVEL% EQU 0 (
    echo ビルド完了: git-cli.exe
    echo ファイルサイズ:
    dir git-cli.exe | findstr git-cli.exe
) else (
    echo ビルドエラーが発生しました
)

pause