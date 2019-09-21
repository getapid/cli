package transaction

import (
	"github.com/iv-p/apiping/svc/client/step"
)

// Transaction is the definition of a transaction
type Transaction struct {
	ID        string                 `yaml:"id"`
	Variables map[string]interface{} `yaml:"variables"`
	Steps     []step.Step            `yaml:"steps"`
}