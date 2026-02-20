package input

import (
    "bufio"
    "fmt"
    "net/url"
    "os"
    "strings"
)

func ReadURLs(path string) ([]string, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("open input file: %w", err)
    }
    defer f.Close()

    var out []string
    sc := bufio.NewScanner(f)
    for sc.Scan() {
        line := strings.TrimSpace(sc.Text())
        if line == "" {
            continue
        }
        out = append(out, NormalizeURL(line))
    }
    if err := sc.Err(); err != nil {
        return nil, fmt.Errorf("scan input file: %w", err)
    }
    return out, nil
}

func NormalizeURL(s string) string {
    s = strings.TrimSpace(s)
    if s == "" {
        return s
    }

    u, err := url.Parse(s)
    if err == nil && u.Scheme != "" {
        return s
    }
    // If parsing failed or scheme is missing, treat it as host/path and add scheme.
    return "https://" + s
}
