package main

// orderData contains item_id: order_id
var orderData = make(map[int][]int)

func main() {
	loadOrder()
	loadRules()
}
