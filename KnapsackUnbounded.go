package main
 
import "fmt"
 
type item struct {
    name string
    value int
    weight, volume float64
} 

type Result struct {
	counts []int
	sum    int
}
 
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
 
//O(n!)
func knap(wants []item, w, v float64) (res Result) {
	if len(wants) == 0 {
		return
	}
	n := len(wants) - 1
	maxCount := min(int(w/wants[n].weight), int(v/wants[n].volume))
	for count:=0; count<=maxCount; count++ {
		sol := knap(wants[:n],
			w-float64(count)*wants[n].weight,
			v-float64(count)*wants[n].volume)
		sol.sum += wants[n].value * count
		if sol.sum > res.sum {
			sol.counts = append(sol.counts, count)
			res = sol
		}
	}
	return
}
 
func main() {
	items := []item{
		//nama, value, weight, volume
		{"Mangga", 3000, 0.3, 0.025},
		{"Pisang", 1800, 0.2, 0.015},
		{"Jambu", 2500, 2.0, 0.002},
		{"Nanas", 1500, 1.2, 0.012},
	}
	var sumCount, sumValue int
	var sumWeight, sumVolume float64
 
	result := knap(items, 30, 0.25)
	
	fmt.Println("Knapsack Unbound") 
	fmt.Println("Berat Maksimum : 25")
	fmt.Println("Volume Maksimum : 0.25")
	fmt.Println("Item yang dapat dibawa :")
	for i:=0; i<len(result.counts); i++ {
		fmt.Printf("%s   x %d -> Berat: %.3f  Volume: %.3f  Value: %d\n",
			items[i].name, result.counts[i], items[i].weight*float64(result.counts[i]),
			items[i].volume*float64(result.counts[i]), items[i].value*result.counts[i])
 
		sumCount += result.counts[i]
		sumValue += items[i].value * result.counts[i]
		sumWeight += items[i].weight * float64(result.counts[i])
		sumVolume += items[i].volume * float64(result.counts[i])
	}
 
	fmt.Printf("Total %d -> Berat: %.3f  Volume: %.3f  Value: %d\n", sumCount, sumWeight, sumVolume, sumValue)
}