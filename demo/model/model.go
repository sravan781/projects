package model

type Person struct {
	Name string `json:"name"`
	Pin  string `json:"Pin" binding:"required"`
}
type Deposite struct {
	Name        string `json:"name"`
	Account_num int    `json:"account_num" unique:"true"`
	Pin         string `json:"Pin" binding:"required"`
	Balance     int    `json:"balance"`
}

type Transfer struct {
	From   int    `json:"from"`
	To     int    `json:"to"`
	Pin    string `json:"pin"`
	Amount int    `json:"amount"`
}

type State struct {
	Name         string   `json:"name"`
	Account_num  int      `json:"account_num"`
	Pin          string   `json:"Pin"`
	Balance      int      `json:"balance"`
	Transactions []string `json:"transactions"`
}

type Reset struct {
	Account_num int    `json:"account_num"`
	Pin         string `json:"Pin"`
	New_pin     string `json:"new_pin"`
}
