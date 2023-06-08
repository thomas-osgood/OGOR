package lfichecker

// slice containing LFI filter evasion patterns.
var lfipatterns []string = []string{"../", "..//", "..%2f", "%2e%2e/", ".././"}
