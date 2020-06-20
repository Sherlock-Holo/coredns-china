package china_list

import (
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	chinalist "github.com/Sherlock-Holo/coredns-china/internal/china-list"
	errors "golang.org/x/xerrors"
)

func Run(output string) error {
	dataPath, err := chinalist.DownloadData()
	if err != nil {
		return errors.Errorf("download data failed: %w", err)
	}

	defer func() {
		_ = os.RemoveAll(dataPath)
	}()

	var domains []chinalist.Domain

	for _, fileName := range chinalist.DataFileNames {
		subDomains, err := chinalist.Parse(filepath.Join(dataPath, fileName))
		if err != nil {
			return errors.Errorf("parse file %s failed: %w", fileName, err)
		}

		domains = append(domains, subDomains...)
	}

	// dedup
	m := make(map[chinalist.Domain]struct{}, len(domains))
	for _, domain := range domains {
		m[domain] = struct{}{}
	}
	domains = domains[:0]
	for domain := range m {
		domains = append(domains, domain)
	}

	sort.Slice(domains, func(i, j int) bool {
		return domains[i] < domains[j]
	})

	sb := new(strings.Builder)
	for _, domain := range domains {
		sb.WriteString(string(domain))
		sb.WriteRune('\n')
	}

	file, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Errorf("create file %s failed: %w", output, err)
	}
	defer func() {
		_ = file.Close()
	}()

	if _, err := io.WriteString(file, sb.String()); err != nil {
		return errors.Errorf("write domains failed: %w", err)
	}

	return nil
}
