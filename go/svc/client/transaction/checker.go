package transaction

import (
	"github.com/iv-p/apiping/svc/client/step"
	"github.com/iv-p/apiping/svc/client/variables"
)

type Checker interface {
	check(Transaction, variables.Variables) SingleTransactionResult
}

type TransactionChecker struct {
	stepChecker      step.Checker
	stepInterpolator Interpolator

	Checker
}

type StepResult struct {
	Step     step.Step
	Response step.HTTPResponse
	Result   step.ValidationResult
}

type SingleTransactionResult struct {
	SequenceIds []string
	Steps       map[string]StepResult
}

func NewStepChecker(stepChecker step.Checker, interpolator Interpolator) Checker {
	return &TransactionChecker{
		stepChecker:      stepChecker,
		stepInterpolator: interpolator,
	}
}

func (c *TransactionChecker) check(transaction Transaction, vars variables.Variables) SingleTransactionResult {
	res := SingleTransactionResult{
		Steps: make(map[string]StepResult),
	}
	for _, step := range transaction.Steps {
		vars = vars.Merge("variables", step.Variables)
		prepared := c.stepInterpolator.interpolate(step, vars.Get())
		response, result := c.stepChecker.Check(prepared)
		res.SequenceIds = append(res.SequenceIds, prepared.ID)
		res.Steps[prepared.ID] = StepResult{prepared, response, result}
	}
	return res
}