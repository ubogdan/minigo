package main


func f0() *int {
	p := &Point{
		x: 1,
		y: 2,
	}
	x := &p.x
	return x
}

func f1() bytes {
	f := &StmtFor{
		rng: &ForRangeClause{
			invisibleMapCounter: &ExprVariable{
				id: 1,
			},
		},
	}
	mapCounter := &Relation{
		name: S("2"),
		expr: f.rng.invisibleMapCounter,
	}
	return mapCounter.name
}

func main() {
	p := f0()
	fmtPrintf("%d\n", *p)
	s := f1()
	fmtPrintf("%s\n", s)
}

type Point struct {
	x int
	y int
}

type StmtFor struct {
	rng *ForRangeClause
}

type ForRangeClause struct {
	invisibleMapCounter *ExprVariable
}

type ExprVariable struct {
	id int
}

func (e *ExprVariable) f() {
}

type Relation struct {
	name bytes
	expr Expr
}

type Expr interface {
	f()
}
