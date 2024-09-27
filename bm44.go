package main

func IsValid(s string) bool {
	dict := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	st := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if val, ok := dict[s[i]]; !ok {
			st = append(st, s[i])
		} else {
			if len(st) == 0 || st[len(st)-1] != val {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}
