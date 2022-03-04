package utils

// 工具类不做存储
//type queue struct{}
//
//var Queue = new(queue)
//
//func (this *queue) Push(key string, value interface{}) error {
//	val, err := Json.Marshal(value)
//	if err != nil {
//		return err
//	}
//	return client.GetRedis().LPush(key, string(val)).Err()
//}
//
//func (this *queue) Pop(key string, val interface{}) error {
//	res, err := client.GetRedis().RPop(key).Result()
//	if err != nil {
//		return err
//	}
//	return Json.UnmarshalFromString(res, val)
//}
