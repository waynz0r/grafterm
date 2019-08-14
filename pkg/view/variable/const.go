package variable

import (
	"github.com/waynz0r/grafterm/pkg/model"
)

// ConstVariabler is used to manager contant variables in the application.
type ConstVariabler struct {
	cfg model.Variable
}

// Scope Satisfies Variabler interface.
func (c ConstVariabler) Scope() Scope {
	return ScopeDashboard
}

// IsRepeatable Satisfies Variabler interface.
func (c ConstVariabler) IsRepeatable() bool {
	return false
}

// GetValue Satisfies Variabler interface.
func (c ConstVariabler) GetValue() string {
	return c.cfg.Constant.Value
}
