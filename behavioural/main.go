package main

import (
	cm "TMPS/behavioural/complain-manager"
	"TMPS/behavioural/lottery"
	"TMPS/behavioural/menu"
)

func main() {
	//vendingMachine := newVendingMachine(1, 10)
	//
	//err := vendingMachine.requestItem()
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//
	//err = vendingMachine.insertMoney(10)
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//
	//err = vendingMachine.dispenseItem()
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//
	//fmt.Println()
	//
	//err = vendingMachine.addItem(2)
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//
	//fmt.Println()
	//
	//err = vendingMachine.requestItem()
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//
	//err = vendingMachine.insertMoney(10)
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//
	//err = vendingMachine.dispenseItem()
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}

	ceo := &cm.CEO{}

	administrator := &cm.Administrator{}
	administrator.SetNext(ceo)

	chef := &cm.Chef{}
	chef.SetNext(administrator)

	cook := &cm.Cook{}
	cook.SetNext(chef)

	waiter := &cm.Waiter{}
	waiter.SetNext(cook)

	customer := &cm.Customer{Name: "Irina", Problem: cm.RentTable}
	//Patient visiting
	waiter.Execute(customer)

	smsOTP := &lottery.Sms{}
	o := lottery.Code{
		smsOTP,
	}
	o.GenAndSendCode(4)

	newMenu := menu.NewMenu([]string{"Tea", "Coffee"})

	observerFirst := &menu.Customer{Email: "abc@gmail.com"}
	observerSecond := &menu.Customer{Email: "xyz@gmail.com"}

	newMenu.Register(observerFirst)
	newMenu.Register(observerSecond)

	newMenu.UpdateMenu("pizza")

}
