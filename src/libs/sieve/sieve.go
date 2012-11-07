package sieve

const sieveSize = 10000000

type Sieve [sieveSize]bool

func markSieve(n int64, s *Sieve) {
	for m := 2*n; m < sieveSize; m += n {
		s[m] = true
	}
}

func MakeSieve(upTo int64) (*Sieve) {
	var s *Sieve = new(Sieve)
	s[0] = true
	s[1] = true
	for prime := int64(2); prime < upTo; {
		markSieve(prime,s)
		for prime++; s[prime]; prime++ { }
	}
	return s
}

func (s *Sieve) IsPrime(x int64) (bool) {
	if x < 0 { return false }
	return !s[x]
}
