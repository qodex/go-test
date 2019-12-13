package main

//Contains returns true if provided slice of strings contains provided string
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//MergeUnique adds unique elements from s2 to s1
func MergeUnique(s1 []string, s2 []string, exclude string) []string {
	for _, s := range s2 {
		if !Contains(s1, s) && s != exclude {
			s1 = append(s1, s)
		}
	}
	return s1
}
