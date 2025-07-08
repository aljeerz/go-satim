package satim

func OrderStatusFromCode(code int64) SatimOrderStatus {
	switch code {
	case 0:
		return "Pending"
	case 1:
		return "PartiallyPaid"
	case 2:
		return "Paid"
	case 4:
		return "Refunded"
	case 5:
		return "Pending"
	case 6:
		return "Declined"
	case 7:
		return "Cancelled"
	default:
		return "Failed"
	}
}
