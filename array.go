package main

//func main() {
//	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	fmt.Printf("%v\n", paginate(items, 4, 2))
//}
//
//func paginate(slice []int, pageNumber, itemsPer int) []int {
//	start := clamp(pageNumber*itemsPer, 0, len(slice))
//	end := clamp(start+itemsPer, 0, len(slice))
//	return slice[start:end]
//}
//
//func clamp(val, min, max int) int {
//	fmt.Printf("val: %v\n", val)
//	fmt.Printf("min: %v\n", min)
//	fmt.Printf("max: %v\n", max)
//	if val < min {
//		return min
//	}
//	if val > max {
//		return max
//	}
//
//	return val
//}
//
//// var arr = []int{1, 2, 3, 4, 5, 6, 7}
//// var pageNumber = 1
//// var itemsPer = 3
////
//// var leftOver int
//// if pageNumber == 0 && len(arr) > itemsPer {
//// 	leftOver = itemsPer
//// }
//// else if pageNumber == 0 && len(arr) < itemsPer {
//// 	leftOver = len(arr)
//// }
//// else if len(arr) > ((pageNumber * itemsPer) + itemsPer) {
////
//// }
//// // if pageNumber == 0 {
//// // 	leftOver = len(arr) - itemsPer
//// // } else if  {
//// // 	leftOver = len(arr) % (pageNumber * itemsPer)
//// // } else {
//// // 	leftOver = len(arr[leftOver:])
//// // }
////
//// fmt.Printf("len(arr): %v\n", len(arr))
//// fmt.Printf("pageNumber: %v\n", pageNumber)
//// fmt.Printf("itemsPer: %v\n", itemsPer)
//// fmt.Printf("leftOver: %v\n", leftOver)
