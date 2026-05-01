package backtracking

import "strconv"

func restoreIpAddresses(s string) []string {
	return getIps([]byte(s), 3, "")
}

func getIps(s []byte, dots int, prefix string) []string {
	ips := make([]string, 0)

	if dots == 0 {
		if ipIsValid(s) {
			ips = append(ips, prefix+string(s))
		}

		return ips
	}

	for i := 1; i <= len(s); i++ {
		if !ipIsValid(s[:i]) {
			break
		}

		ips = append(ips, getIps(s[i:], dots-1, prefix+string(s[:i])+".")...)
	}

	return ips
}

func ipIsValid(s []byte) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) != 1 && s[0] == '0' {
		return false
	}

	num, _ := strconv.Atoi(string(s))

	if num > 255 {
		return false
	}

	return true
}
