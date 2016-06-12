package rev

import "crypto/sha1"

func eq(a []byte, b []byte) bool {
	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for idx := 0; idx < len(a); idx++ {
		if a[idx] != b[idx] {
			return false
		}
	}

	return true
}

func hash(parentHash []byte, key []byte, value []byte) []byte {
	h := sha1.New()
	h.Write(parentHash)
	h.Write(key)
	h.Write(value)
	return h.Sum(nil)
}
