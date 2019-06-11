// Copyright 2012-2015 Samuel Stauffer. All rights reserved.
// Use of this source code is governed by a 3-clause BSD
// license that can be found in the LICENSE file.

package generator

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// CamelCase returns the string converted to camel case (e.g. some_name to SomeName)
func CamelCase(s string) string {
	prev := '_'
	return strings.Map(
		func(r rune) rune {
			if r == '_' {
				prev = r
				return -1
			}
			if prev == '_' {
				prev = r
				return unicode.ToUpper(r)
			}
			prev = r
			return r
		}, s)
}

func camelCase(st string) string {
	if strings.ToUpper(st) == st {
		st = strings.ToLower(st)
	}
	return CamelCase(st)
}

func lowerCamelCase(st string) string {
	if len(st) <= 1 {
		return strings.ToLower(st)
	}
	st = CamelCase(st)
	switch st {
	case "ID":
		return "id"
	case "URL":
		return "url"
	case "HTTP":
		return "http"
	case "API":
		return "api"
	case "TODO":
		return "todo"
	}
	return strings.ToLower(st[:1]) + st[1:]
}

// Converts a string to a valid Golang identifier, as defined in
// http://golang.org/ref/spec#identifier
// by converting invalid characters to the value of replace.
// If the first character is a Unicode digit, then replace is
// prepended to the string.
func validIdentifier(st string, replace string) string {
	var (
		invalidRune  = regexp.MustCompile("[^\\pL\\pN_]")
		invalidStart = regexp.MustCompile("^\\pN")
		out          string
	)
	out = invalidRune.ReplaceAllString(st, "_")
	if invalidStart.MatchString(out) {
		out = fmt.Sprintf("%v%v", replace, out)
	}
	return out
}

// Given a map with string keys, return a sorted list of keys.
// If m is not a map or doesn't have string keys then return nil.
func sortedKeys(m interface{}) []string {
	value := reflect.ValueOf(m)
	if value.Kind() != reflect.Map || value.Type().Key().Kind() != reflect.String {
		return nil
	}

	valueKeys := value.MapKeys()
	keys := make([]string, len(valueKeys))
	for i, k := range valueKeys {
		keys[i] = k.String()
	}
	sort.Strings(keys)
	return keys
}
