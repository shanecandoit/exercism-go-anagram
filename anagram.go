package anagram

// package main

import (
	"sort"
	"strings"
)

// Detect possible anagrams from the subject to possible candidates
func Detect(subject string, candidates []string) (anagrams []string) {

	// Given "listen" and a list of candidates like "enlists" "google" "inlets" "banana"
	// the program should return a list containing "inlets".
	// tip: anagrams are identical when sorted

	// ignore case on subject and candidate
	// fmt.Println("Detect:", subject, candidates)
	subject = strings.ToLower(subject)

	sortedSubject := SortString(subject)
	// fmt.Println("sorted_subject", sorted_subject)

	anagramCands := []string{}
	for _, s := range candidates {

		if len(s) != len(subject) {
			// fmt.Println("bad_len", s)
			continue
		}

		cand := strings.ToLower(s)
		sortCand := SortString(cand)
		// fmt.Println("sort_cand", sort_cand)
		if sortCand == sortedSubject {
			unsortedCand := s
			if cand == subject {
				// dont match identical words
				continue
			}
			anagramCands = append(anagramCands, unsortedCand)
		}
	}

	return anagramCands
}

// SortString break, sort, join is wastefull
func SortString(w string) string {
	// https://stackoverflow.com/questions/22688651/golang-how-to-sort-string-or-byte
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")

	// this fails: panic: reflect: call of Swapper on zero Value
	// sort.Slice(w, func(i int, j int) bool { return w[i] < w[j] })
	// return string(w)
}
