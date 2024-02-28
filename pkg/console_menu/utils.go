package consolemenu

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	OS_PATH = "/etc/os-release"
)

func GetDistroName() string {
	data, err := os.ReadFile(OS_PATH)
	if err != nil {
		return "NONE"
	}

	content := string(data)
	regexID := regexp.MustCompile(`ID=(\w+)`)
	ID := regexID.FindStringSubmatch(content)[1]

	regexVersionID := regexp.MustCompile(`VERSION_ID="(.*)"`)
	versionID := regexVersionID.FindStringSubmatch(content)[1]
	caser := cases.Title(language.BrazilianPortuguese)

	return caser.String(fmt.Sprintf("%s %s", ID, versionID))
}

type MemStats struct {
	Total float64
	Used  float64
	Free  float64
}

func GetMemStats() *MemStats {
	memStats := &MemStats{}

	memInfo, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return memStats
	}

	memInfoStr := string(memInfo)
	lines := strings.Split(memInfoStr, "\n")

	cached := 0.0
	buffers := 0.0

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) >= 3 {
			switch parts[0] {
			case "MemTotal:":
				memStats.Total, _ = strconv.ParseFloat(parts[1], 64)
			case "MemFree:":
				memStats.Free, _ = strconv.ParseFloat(parts[1], 64)
			case "MemAvailable:":
				memStats.Free, _ = strconv.ParseFloat(parts[1], 64)
			case "Cached":
				cached, _ = strconv.ParseFloat(parts[1], 64)
			case "Buffers":
				buffers, _ = strconv.ParseFloat(parts[1], 64)
			}
		}
	}

	memStats.Total = memStats.Total / 1024.0 / 1024.0
	memStats.Free = memStats.Free / 1024.0 / 1024.0
	memStats.Used = (memStats.Total - memStats.Free - buffers - cached)

	return memStats
}

func GetIPV4() string {
	file, err := os.ReadFile("./ip")
	if err == nil {
		return strings.TrimSpace(string(file))
	}

	response, err := http.Get("https://ipv4.icanhazip.com/")
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return ""
	}

	ipv4 := strings.TrimSpace(string(bytes))
	os.WriteFile("./ip", []byte(ipv4), 0644)
	return ipv4
}
