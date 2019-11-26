package config

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"github.com/uestc-acm/acm-training/util"
	"os"
	"regexp"
	"testing"
)

func TestVersionConsistency(t *testing.T) {
	readmeFile, err := os.Open("../README.md")
	assert.Nil(t, err, "Error occurs when opening README.md: %s", err)

	reader := bufio.NewReader(readmeFile)
	for {
		buf, _, err := reader.ReadLine()
		util.CheckIOError(err)
		re := regexp.MustCompile(`!\[\]\(https://img\.shields\.io/badge/version-(\d+.\d+\.\d+)-blue\.svg\)`)
		groups := re.FindStringSubmatch(string(buf))
		if len(groups) != 0 {
			assert.Equal(t, Version, groups[1])
			break
		}
	}
}
