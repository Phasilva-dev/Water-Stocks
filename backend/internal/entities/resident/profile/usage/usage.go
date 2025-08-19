package usage

import (
	"simulation/internal/misc"
)

type UsageProfile struct {
	selector *misc.PercentSelector[TimeRangeCalculator] 

}