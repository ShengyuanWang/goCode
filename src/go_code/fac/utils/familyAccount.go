package utils

import (
	"fmt"
)

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	details string
	flag    bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说明",
		flag:    false,
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("--------当前收支明细记录---------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("可以考虑开始花钱了")
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额的金额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("--------家庭收支记账软件---------")
		fmt.Println("         1 收支明细")
		fmt.Println("         2 登记收入")
		fmt.Println("         3 登记指出")
		fmt.Println("         4 退出软件")
		fmt.Println("请选择（1-4）:")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项:")

		}

		if !this.loop {
			break
		}

	}
}
