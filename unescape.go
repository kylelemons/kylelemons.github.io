package main

import (
	"bufio"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func unescape(s string) string {
	s = html.UnescapeString(s)
	s = strings.Replace(s, "â€™", "'", -1)
	s = strings.Replace(s, "\n\n\n", "\n\n", -1)

	s = regexp.MustCompile(`_[\S]+_`).ReplaceAllStringFunc(s, func(matched string) string {
		matched = strings.Replace(matched, "_", " ", -1)
		matched = strings.TrimSpace(matched)
		return "_" + matched + "_"
	})
	s = regexp.MustCompile("`[\\S]+`").ReplaceAllStringFunc(s, func(matched string) string {
		matched = strings.Replace(matched, "`", " ", -1)
		matched = strings.TrimSpace(matched)
		return "`" + matched + "`"
	})
	s = regexp.MustCompile(`\[\[.*?\]\[.*?\]\]`).ReplaceAllStringFunc(s, func(matched string) string {
		groups := regexp.MustCompile(`\[\[(.*?)\]\[(.*?)\]\]`).FindStringSubmatch(matched)
		return fmt.Sprintf("[%s](%s)", groups[1], groups[2])
	})
	s = regexp.MustCompile(`(?m)^\*+\s`).ReplaceAllStringFunc(s, func(matched string) string {
		return strings.Replace(matched, "*", "#", -1)
	})
	s = regexp.MustCompile(`(?m)^  \S`).ReplaceAllStringFunc(s, func(matched string) string {
		return "  " + matched
	})

	return s
}

func main() {
	if len(os.Args) > 1 {
		unescapeFiles(os.Args[1:])
	} else {
		unescapeStdin()
	}
}

func unescapeFiles(files []string) {
	file := func(filename string) {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Printf("Read: %s", err)
			return
		}
		data = []byte(unescape(string(data)))
		if err := ioutil.WriteFile(filename, data, 0644); err != nil {
			log.Printf("Write: %s", err)
		}
	}

	for _, f := range files {
		file(f)
	}
}

func unescapeStdin() {
	ready := make(chan string)
	go func() {
		defer close(ready)
		lines := bufio.NewScanner(os.Stdin)
		for lines.Scan() {
			ready <- unescape(lines.Text())
		}
	}()

	batch := func() (lines []string, err error) {
		fmt.Println("# Paste some lines:")
		for {
			var timeout <-chan time.Time
			if len(lines) > 0 {
				timeout = time.After(50 * time.Millisecond)
			}
			select {
			case line, ok := <-ready:
				if !ok {
					return lines, io.EOF
				}
				lines = append(lines, line)
			case <-timeout:
				return
			}
		}
	}

	for {
		lines, err := batch()
		if len(lines) > 0 {
			fmt.Println()
			fmt.Println("# Unescaped:")
		}
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println()
		if err != nil {
			return
		}
	}
}
