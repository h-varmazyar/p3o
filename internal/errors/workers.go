package errors

import(
	
	"github.com/h-varmazyar/p3o/pkg/errors"
)
var (
	ErrVisitWorkerStartedBefore = errors.NewWithCode("visit_worker_start_before", 40001)
	ErrNilVisitChannel          = errors.NewWithCode("nil_visit_chan", 40002)
)