package china_list

import (
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	china_list "github.com/Sherlock-Holo/coredns-china/internal/china-list"
	"github.com/Sherlock-Holo/errors"
)

func Run(output string) error {
	dataPath, err := china_list.DownloadData()
	if err != nil {
		return errors.WithMessage(err, "download data failed")
	}

	defer func() {
		_ = os.RemoveAll(dataPath)
	}()

	var domains []china_list.Domain

	for _, fileName := range china_list.DataFileNames {
		subDomains, err := china_list.Parse(filepath.Join(dataPath, fileName))
		if err != nil {
			return errors.WithMessagef(err, "parse file %s failed", fileName)
		}

		domains = append(domains, subDomains...)
	}

	// dedup
	m := make(map[china_list.Domain]struct{}, len(domains))
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
		return errors.Wrapf(err, "create file %s failed", output)
	}
	defer func() {
		_ = file.Close()
	}()

	if _, err := io.WriteString(file, sb.String()); err != nil {
		return errors.Wrap(err, "write domains failed")
	}

	return nil
}
