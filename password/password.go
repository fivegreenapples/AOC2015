package password

type Password string

func (p *Password) IncrementUntilValid() {
	p.increment()
	for !p.isValid() {
		p.increment()
	}
}

func (p *Password) increment() {

	length := len(*p)
	if length == 0 {
		*p = "a"
		return
	}

	chars := []byte(*p)
	idx := length - 1

	for {
		char := chars[idx]
		if char == 'z' {
			char = 'a'
		} else {
			char = char + 1
		}
		chars[idx] = char

		if char == 'a' {
			idx--
			continue
		}

		break
	}

	*p = Password(chars)
}

func (p *Password) isValid() bool {
	chars := []byte(*p)

	var prev, prev2 byte
	var hasStraight, hasIOL bool
	mapPairs := map[byte]bool{}
	for _, cur := range chars {

		if cur == prev+1 && prev == prev2+1 {
			hasStraight = true
		}

		if cur == 'i' || cur == 'o' || cur == 'l' {
			hasIOL = true
		}

		if cur == prev {
			mapPairs[cur] = true
		}

		prev2 = prev
		prev = cur
	}

	return hasStraight && !hasIOL && len(mapPairs) >= 2
}
