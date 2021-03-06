/**
 * Copyright 2015 @ S1N1 Team.
 * name : order_item
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package shopping

// 订单商品项
type OrderItem struct {
	Id         int     `db:"id" pk:"yes" auto:"yes" json:"id"`
	SnapshotId int     `db:"snapshot_id"`
	Quantity   int     `db:"quantity"`
	Sku        string  `db:"sku"`
	Fee        float32 `db:"fee"`
}
