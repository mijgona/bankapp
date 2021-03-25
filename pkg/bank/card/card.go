package card

import (
	"bank/pkg/bank/types"
)
//IssueCard выпуск карты
func IssueCard(currency types.Currency, color string, name string) types.Card{
	card:= types.Card{
		ID:			1000,
		PAN:		"5058 xxxx xxxx 0001",
		Balance:	0,
		Currency:	currency,
		Color:		color,
		Name:		name,
		Active:		true,
		MinBalance:	10_000_00,
	}
	return card
}
//Withdraw снятие денег со счета
func Withdraw(card *types.Card, amount types.Money)  {

	if !card.Active{
		return			
	}
	if(card.Balance<=amount){
		return
	}
	if(amount<0){
		return
	}
	if (amount>20_000_00){
		return	
	}
	card.Balance-=amount
}

//Deposit зачисление депозита на счет
func Deposit(card *types.Card, amount types.Money)  {
	if !card.Active{
		return			
	}
	if (amount>50_000_00){
		return	
	}
	if (amount<0){
		return	
	}
	card.Balance+=amount
	
}

//AddBonus расчитывает процент на остаток
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active{
		return
	}
	if (card.Balance<0){
		return
	}

	bonus:=(((int(card.MinBalance)/100)*percent)*daysInMonth)/daysInYear
	if (bonus<=0){
		return
	}

	if (bonus>5_000_00){
		return	   
	}
	card.Balance+=types.Money(bonus)
}

//Total вычисляет общую сумму средств на всех картах
//Отрицательные суммы и суммы на заблокированных картах игнорируются
func Total(cards []types.Card) types.Money  {
	sum:=types.Money(0)
	for _, card := range cards {
		if (card.Balance<0) {
			continue
		}
		if !card.Active {
			continue
		}
		sum+=card.Balance
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource{
	var paymentSource []types.PaymentSource

	for _, card := range cards {
		if card.Balance>0{
			if card.Active{
				paymentSource=append(paymentSource,types.PaymentSource{
					Type:"card", 
					Number:string(card.PAN), 
					Balance:card.Balance,
				})
			}
		}		
	}


	return paymentSource
}

   