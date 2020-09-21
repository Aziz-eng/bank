package card

import(
	"github.com/Aziz-eng/bank/pkg/bank/types"
)
// IssueCard is Issue a new card
func IssueCard(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: types.TJS,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	
}