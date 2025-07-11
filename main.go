package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "status":
		gitStatus()
	case "add":
		gitAdd(args)
	case "commit":
		gitCommit(args)
	case "push":
		gitPush(args)
	case "pull":
		gitPull(args)
	case "log":
		gitLog(args)
	case "clone":
		gitClone(args)
	case "help", "-h", "--help":
		showHelp()
	default:
		fmt.Printf("不明なコマンド: %s\n", command)
		showHelp()
	}
}

func showHelp() {
	fmt.Println("Windows向け軽量Gitコマンド（Git非依存版）")
	fmt.Println("使用方法:")
	fmt.Println("  git-cli status          - 作業ディレクトリの状態を表示")
	fmt.Println("  git-cli add <ファイル>   - ファイルをステージングエリアに追加")
	fmt.Println("  git-cli commit -m <メッセージ> - 変更をコミット")
	fmt.Println("  git-cli push            - リモートリポジトリにプッシュ")
	fmt.Println("  git-cli pull            - リモートリポジトリから取得")
	fmt.Println("  git-cli log             - コミット履歴を表示")
	fmt.Println("  git-cli clone <URL>     - リポジトリをクローン")
	fmt.Println("  git-cli help            - このヘルプを表示")
}

func getRepository() (*git.Repository, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainOpen(wd)
	if err != nil {
		return nil, fmt.Errorf("gitリポジトリが見つかりません: %v", err)
	}
	return repo, nil
}

func gitStatus() {
	repo, err := getRepository()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	status, err := worktree.Status()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	if status.IsClean() {
		fmt.Println("作業ディレクトリはクリーンです")
		return
	}

	fmt.Println("変更されたファイル:")
	for file, fileStatus := range status {
		var statusText string
		switch {
		case fileStatus.Staging == git.Untracked:
			statusText = "?? " // 未追跡
		case fileStatus.Staging == git.Added:
			statusText = "A  " // 追加（ステージング済み）
		case fileStatus.Staging == git.Modified:
			statusText = "M  " // 変更（ステージング済み）
		case fileStatus.Worktree == git.Modified:
			statusText = " M " // 変更（未ステージング）
		case fileStatus.Worktree == git.Deleted:
			statusText = " D " // 削除（未ステージング）
		default:
			statusText = "   "
		}
		fmt.Printf("%s %s\n", statusText, file)
	}
}

func gitAdd(args []string) {
	if len(args) == 0 {
		fmt.Println("エラー: 追加するファイルを指定してください")
		fmt.Println("使用例: git-cli add ファイル名")
		return
	}

	repo, err := getRepository()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	for _, file := range args {
		if file == "." {
			err = worktree.AddGlob("*")
		} else {
			_, err = worktree.Add(file)
		}
		if err != nil {
			fmt.Printf("エラー: ファイル '%s' を追加できませんでした: %v\n", file, err)
			continue
		}
		fmt.Printf("ファイルを追加しました: %s\n", file)
	}
}

func gitCommit(args []string) {
	if len(args) < 2 || args[0] != "-m" {
		fmt.Println("エラー: コミットメッセージを指定してください")
		fmt.Println("使用例: git-cli commit -m \"コミットメッセージ\"")
		return
	}

	repo, err := getRepository()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	commit, err := worktree.Commit(args[1], &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Git CLI User",
			Email: "user@git-cli.local",
			When:  time.Now(),
		},
	})

	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	fmt.Printf("コミットが作成されました: %s\n", commit.String()[:8])
}

func gitPush(args []string) {
	repo, err := getRepository()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	err = repo.Push(&git.PushOptions{})
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Println("既に最新の状態です")
			return
		}
		fmt.Printf("エラー: プッシュに失敗しました: %v\n", err)
		return
	}

	fmt.Println("プッシュが完了しました")
}

func gitPull(args []string) {
	repo, err := getRepository()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
	})

	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Println("既に最新の状態です")
			return
		}
		fmt.Printf("エラー: プルに失敗しました: %v\n", err)
		return
	}

	fmt.Println("プルが完了しました")
}

func gitLog(args []string) {
	repo, err := getRepository()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	ref, err := repo.Head()
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	commitIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		fmt.Printf("エラー: %v\n", err)
		return
	}

	count := 0
	maxCount := 10
	if len(args) > 0 && strings.Contains(args[0], "-") {
		// 簡易的なオプション解析（-n 数値形式のみ対応）
		for i, arg := range args {
			if arg == "-n" && i+1 < len(args) {
				fmt.Sscanf(args[i+1], "%d", &maxCount)
				break
			}
		}
	}

	err = commitIter.ForEach(func(c *object.Commit) error {
		if count >= maxCount {
			return fmt.Errorf("達成")
		}
		fmt.Printf("%s %s\n", c.Hash.String()[:8], strings.Split(c.Message, "\n")[0])
		count++
		return nil
	})

	if err != nil && err.Error() != "達成" {
		fmt.Printf("エラー: %v\n", err)
	}
}

func gitClone(args []string) {
	if len(args) == 0 {
		fmt.Println("エラー: クローンするリポジトリのURLを指定してください")
		fmt.Println("使用例: git-cli clone https://github.com/user/repo.git")
		return
	}

	url := args[0]
	var directory string

	if len(args) > 1 {
		directory = args[1]
	} else {
		// URLからリポジトリ名を抽出
		parts := strings.Split(url, "/")
		repoName := parts[len(parts)-1]
		if strings.HasSuffix(repoName, ".git") {
			repoName = repoName[:len(repoName)-4]
		}
		directory = repoName
	}

	// ディレクトリが既に存在するかチェック
	if _, err := os.Stat(directory); !os.IsNotExist(err) {
		fmt.Printf("エラー: ディレクトリ '%s' は既に存在します\n", directory)
		return
	}

	fmt.Printf("リポジトリをクローンしています: %s -> %s\n", url, directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Printf("エラー: クローンに失敗しました: %v\n", err)
		return
	}

	fmt.Printf("クローンが完了しました: %s\n", directory)
}