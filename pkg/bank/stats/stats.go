package stats


import(
	"github.com/Aziz-eng/bank/pkg/bank/types"
)
// Avg расчитывает среднею сумму платежа
func Avg(payments []types.Payment) types.Money {
	amount := types.Money (0)
	for _, payment := range payments {
		if payment.Amount > 0 {
			amount += payment.Amount
		}
		
	}
	delta := amount / types.Money (len(payments))
	return delta
}