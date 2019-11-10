package china_list

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Sherlock-Holo/errors"
)

const (
	repoName    = "dnsmasq-china-list"
	repoAddress = "https://github.com/felixonmars/dnsmasq-china-list.git"
)

func lookupGitCmd() (path string, err error) {
	path, err = exec.LookPath("git")
	if err != nil {
		err = errors.WithStack(err)
	}
	return
}

func DownloadData() (dirPath string, err error) {
	gitPath, err := lookupGitCmd()
	if err != nil {
		return "", errors.WithMessage(err, "lookup git cmd failed")
	}

	dirPath = filepath.Join(os.TempDir(), repoName)

	cmd := exec.Command(gitPath, "clone", "--depth", "1", repoAddress, dirPath)
	errBuf := new(bytes.Buffer)
	cmd.Stderr = errBuf
	if err := cmd.Run(); err != nil {
		return "", errors.Wrapf(err, "download list failed, error output %s", errBuf.String())
	}

	return
}
