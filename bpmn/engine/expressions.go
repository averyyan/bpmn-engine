package engine

import (
	"strings"

	"github.com/antonmedv/expr"
)

// TODO 解释器
func evaluateExpression(expression string, variableContext map[string]interface{}) (interface{}, error) {
	expression = strings.TrimSpace(expression)
	expression = strings.TrimPrefix(expression, "=")
	return expr.Eval(expression, variableContext)
}
