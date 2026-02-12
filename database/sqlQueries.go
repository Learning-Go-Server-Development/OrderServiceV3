package database

const (
	insertOrder = "insert into orders (cid, order_number, date_entered) values (?, ?, ?) "

	updateOrder = " UPDATE orders SET order_number = ?, date_updated = ? " +
		" where id = ? "

	selectOrderByID = "SELECT id, cid, order_number, date_entered, date_updated " +
		" from  orders " +
		" where id = ? "

	selectOrderList = "SELECT id, cid, order_number, date_entered, date_updated " +
		" from  orders " +
		" where cid = ? "

	deleteOrder = " DELETE from orders " +
		" where id = ? "

	insertItem = "insert into order_items (order_id, product_id) values (?, ?) "

	updateItem = " UPDATE order_items SET product_id = ? " +
		" where id = ? "

	selectItemList = "SELECT id, order_id, product_id " +
		" from  order_items " +
		" where order_id = ? "

	deleteItem = " DELETE from order_items " +
		" where id = ? "
)
