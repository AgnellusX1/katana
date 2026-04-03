package common

import "errors"

var ErrOutOfScope = errors.New("out of scope")
var ErrMaxDepthReached = errors.New("max depth reached")
var ErrDomainPageBudgetExhausted = errors.New("domain page budget exhausted")
