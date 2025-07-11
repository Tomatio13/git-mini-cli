# Windows向け軽量Gitコマンド（Git非依存版）

Windows環境で動作する超軽量なGitコマンドラインツールです。**Gitがインストールされていない環境でも**基本的なGit操作を簡単に実行できます。

## 特徴

- **Git非依存**: Gitがインストールされていなくても動作
- **軽量**: 単一バイナリファイル（x64版: 約8MB、x86版: 約8MB）
- **シンプル**: 複雑な機能を省いた基本的なGit操作のみ
- **Windows最適化**: Windows環境（x64/x86）での動作を前提に設計
- **Go-git使用**: Pure Goで実装されたGitライブラリを使用

## サポートしているコマンド

| コマンド | 説明 |
|---------|------|
| `git-cli status` | 作業ディレクトリの状態を表示 |
| `git-cli add <ファイル>` | ファイルをステージングエリアに追加 |
| `git-cli commit -m "<メッセージ>"` | 変更をコミット |
| `git-cli push` | リモートリポジトリにプッシュ |
| `git-cli pull` | リモートリポジトリから取得 |
| `git-cli log` | コミット履歴を表示（最新10件） |
| `git-cli clone <URL>` | リポジトリをクローン |
| `git-cli help` | ヘルプを表示 |

## ビルド方法

### Windows環境でのビルド

```batch
# build.batを実行
build.bat
```

### Makeを使用したビルド

```bash
# Windowsバイナリをビルド
make build

# Windows x86バイナリをビルド
make build-windows-x86

# Linuxバイナリをビルド（テスト用）
make build-linux

# すべてのプラットフォーム向けビルド
make build-all

# ビルドファイルをクリーンアップ
make clean
```

## 使用例

```bash
# リポジトリをクローン
git-cli clone https://github.com/user/repo.git

# 現在の状態を確認
git-cli status

# ファイルを追加
git-cli add README.md
git-cli add .

# コミット
git-cli commit -m "初回コミット"

# リモートにプッシュ
git-cli push

# リモートから取得
git-cli pull

# ログを確認
git-cli log
```

## 必要な環境

- Windows 10/11（x64またはx86）
- **.NET Frameworkや追加ライブラリは不要**

## 注意事項

- このツールはGo-gitライブラリを使用してGit機能を直接実装しています
- 複雑なGit操作には対応していません
- エラーメッセージは日本語で表示されます
- 認証が必要なリポジトリの場合、SSH設定が必要な場合があります

## ライセンス

MIT License