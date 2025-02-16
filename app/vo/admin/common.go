package admin

func GetList(total int64, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"total": total,
		"data":  data,
	}
}
