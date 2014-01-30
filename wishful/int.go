package wishful

type Int int

func (o Int) Concat(x Semigroup) Semigroup {
	return Int(int(o) + int(x.(Int)))
}
