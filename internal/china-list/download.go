package china_list

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/Sherlock-Holo/errors"
)

const repoAddress = "https://github.com/felixonmars/dnsmasq-china-list.git"

func lookupGitCmd() (path string, err error) {
	path, err = exec.LookPath("git")
	if err != nil {
		err = errors.WithStack(err)
	}
	return
}

func DownloadData() (dirPath string, err error) {
	log.Println("use git to download list")

	gitPath, err := lookupGitCmd()
	if err != nil {
		return "", errors.WithMessage(err, "lookup git cmd failed")
	}

	dirPath, err = ioutil.TempDir(os.TempDir(), "")
	if err != nil {
		return "", errors.Wrap(err, "create temp git directory failed")
	}

	cmd := &exec.Cmd{
		Path:   gitPath,
		Args:   append([]string{gitPath}, "clone", "--depth", "1", repoAddress, dirPath),
		Stderr: os.Stderr,
	}

	if err := cmd.Run(); err != nil {
		return "", errors.Wrap(err, "download list failed")
	}

	return
}
