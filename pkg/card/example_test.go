package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampleWithdraw_positive()  {
	card:=types.Card{Balance:20_000_00, Active:true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	//Output: 1000000
}
 func ExampleDeposit_positive(){
	card:=types.Card{Balance:20_000_00, Active:true}
	Deposit(&card,10_000_00)
	fmt.Println(card.Balance)
	//Output:
	//3000000
	
 }
 func ExampleDeposit_inactive(){
	card:=types.Card{Balance:20_000_00, Active:false}
	Deposit(&card,10_000_00)
	fmt.Println(card.Balance)
	//Output:
	//2000000
	
 }
 func ExampleDeposit_limit(){
	card:=types.Card{Balance:20_000_00, Active:false}
	Deposit(&card,50_001_00)
	fmt.Println(card.Balance)
	//Output:
	//2000000
	
 }

 func ExampleAddBonus_positive()  {
	card:=types.Card{Balance:20_000_00, Active:true, MinBalance: 10_000_00}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)
	
	//Output:
	//2002465
 }

 func ExampleAddBonus_negative()  {
	card:=types.Card{Balance:20_000_00, Active:false, MinBalance: 10_000_00}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)
	
	//Output:
	//2000000
 }

 
 func ExampleAddBonus_negativeBalance()  {
	card:=types.Card{Balance:-20_000_00, Active:true, MinBalance: 10_000_00}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)
	
	//Output:
	//-2000000
 }

 
 func ExampleAddBonus_limit()  {
	card:=types.Card{Balance:20_000_00, Active:true, MinBalance: 20_277_7800_00}
	AddBonus(&card,3,30,365)
	fmt.Println(card.Balance)
	
	//Output:
	//2000000
 }

 func ExampleTotal()  {
	cards:= []types.Card{
	   {
		   Balance:1_000_00,
		   Active: true,
	   },
	}
	fmt.Println(Total(cards))
	cards= []types.Card{
	   {
		   Balance:1_000_00,
		   Active: true,
	   },
	}
	fmt.Println(Total(cards))
	
	cards= []types.Card{
		{
			Balance:1_000_00,
			Active: false,
		},
	 }
	 fmt.Println(Total(cards))
	 cards= []types.Card{
		{
			Balance:-1_000_00,
			Active: true,
		},
	 }
	 fmt.Println(Total(cards))

	 
	 //Output: 
	 //100000
	 //100000
	 //0
	 //0
 }
 func ExamplePaymentSource()  {
	 cards:=[]types.Card{
		{
			ID:  0,
			PAN: "5058 xxxx xxxx 8888",
			Balance: 1_000_000,
			MinBalance: 0,
			Currency:   "",
			Color:      "",
			Name:       "",
			Active:     true,
		},
		{
			ID:  0,
			PAN: "5058 xxxx xxxx 9999",
			Balance: 1_000_000,
			MinBalance: 0,
			Currency:   "",
			Color:      "",
			Name:       "",
			Active:     true,
		},
		{
			ID:  0,
			PAN: "5058 xxxx xxxx 7777",
			Balance: -1_000_000,
			MinBalance: 0,
			Currency:   "",
			Color:      "",
			Name:       "",
			Active:     true,
		},
		{
			ID:  0,
			PAN: "5058 xxxx xxxx 7777",
			Balance: 1_000_000,
			MinBalance: 0,
			Currency:   "",
			Color:      "",
			Name:       "",
			Active:     false,
		},
	 }

	 pSource:=PaymentSources(cards)
	 for _, p := range pSource {
	 	fmt.Println(p.Number)		 
	 }
	 //Output:
	 //5058 xxxx xxxx 8888
	 //5058 xxxx xxxx 9999

 }