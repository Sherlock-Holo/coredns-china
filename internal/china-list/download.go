package china_list

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	errors "golang.org/x/xerrors"
)

const repoAddress = "https://github.com/felixonmars/dnsmasq-china-list.git"

func lookupGitCmd() (path string, err error) {
	path, err = exec.LookPath("git")
	if err != nil {
		err = errors.Errorf("lookup git path failed: %w", err)
	}
	return
}

func DownloadData() (dirPath string, err error) {
	log.Println("use git to download list")

	gitPath, err := lookupGitCmd()
	if err != nil {
		return "", errors.Errorf("lookup git cmd failed: %w", err)
	}

	dirPath, err = ioutil.TempDir(os.TempDir(), "")
	if err != nil {
		return "", errors.Errorf("create temp git directory failed: %w", err)
	}

	cmd := &exec.Cmd{
		Path:   gitPath,
		Args:   append([]string{gitPath}, "clone", "--depth", "1", repoAddress, dirPath),
		Stderr: os.Stderr,
	}

	if err := cmd.Run(); err != nil {
		return "", errors.Errorf("download list failed: %w", err)
	}

	return
}
