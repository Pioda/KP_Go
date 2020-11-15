package node

type Node interface {
	Eval(vars map[string]bool) bool
}

type Not struct{
	Expr Node
}

func (n Not) Eval(vars map[string]bool) bool{
	return !n.Expr.Eval(vars)
}

type And struct{
	LeftExpr Node
	RightExpr Node
}

func (n And) Eval(vars map[string]bool) bool{
	return n.LeftExpr && n.RightExpr
}

type Or struct{
	LeftExpr Node
	RightExpr Node
}

func (n Or) Eval(vars map[string]bool) bool{
	return n.LeftExpr || n.RightExpr
}

type Var struct{
	Name string
}

func (n Var) Eval(vars map[string]bool) bool{
	return vars[n.Name]
}