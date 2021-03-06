package china_list

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"

	errors "golang.org/x/xerrors"
)

const (
	ignorePrefix = "server=/"
	ignoreSuffix = "/114.114.114.114"
)

var (
	DataFileNames = []string{
		"accelerated-domains.china.conf",
	}
)

type Domain string

func Parse(path string) ([]Domain, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Errorf("parse file %s failed: %w", path, err)
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	var domains []Domain
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// ignore empty line
			continue
		}

		if !strings.HasPrefix(line, ignorePrefix) {
			log.Printf("ignore %s", line)
			continue
		}

		if !strings.HasSuffix(line, ignoreSuffix) {
			log.Printf("ignore %s", line)
			continue
		}

		line = string(bytes.ReplaceAll(bytes.ReplaceAll([]byte(line), []byte(ignorePrefix), nil), []byte(ignoreSuffix), nil))
		domains = append(domains, Domain(line))
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.Errorf("scan file %s failed: %w", path, err)
	}

	return domains, nil
}
