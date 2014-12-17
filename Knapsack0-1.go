package main
 
import "fmt"
 
type item struct {
    name string
    weight float64
    value int
}
 
//O(2^i)
func knap(wants []item, i int, max float64) ([]string, float64, int) {
    if i < 0 || max == 0 {
        return nil, 0, 0
    } else if wants[i].weight > max {
        return knap(wants, i-1, max)
    }
    i0, w0, v0 := knap(wants, i-1, max)
    i1, w1, v1 := knap(wants, i-1, max-wants[i].weight)
    v1 += wants[i].value
    if v1 > v0 {
        return append(i1, wants[i].name), w1 + wants[i].weight, v1
    }
    return i0, w0, v0
}

func main() {
    var wants = []item{
        //item name, item weight, item value
        {"map", 9.4, 150},
        {"compass", 13.2, 35},
        {"water", 153.12, 200},
        {"sandwich", 50.4, 160},
        {"glucose", 15.5, 60},
        {"tin", 68.2, 45},
        {"banana", 27.6, 60},
        {"apple", 39.7, 40},
        {"cheese", 23.2, 30},
        {"beer", 52.1, 10},
        {"suntan cream", 11.02, 70},
        {"camera", 32.23, 30},
        {"T-shirt", 24.44, 15},
        {"trousers", 48.54, 10},
        {"umbrella", 73.3, 40},
    }

    fmt.Println("Knapsack 1")
    items, w, v := knap(wants, len(wants)-1, 200)
    fmt.Println("Berat maksimum : 200")
    fmt.Println("Item yang dapat dibawa :", items)
    fmt.Println("Berat total:", w)
    fmt.Println("Value total:", v)

    var wants2 = []item{
        //item name, item weight, item value
        {"beef", 3.8, 36},
        {"pork", 5.4, 43},
        {"ham", 3.6, 90},
        {"greaves", 2.4, 45},
        {"flitch", 4.0, 30},
        {"brawn", 2.5, 56},
        {"welt", 3.7, 67},
        {"salami", 3.0, 95},
        {"sausage", 5.9, 98},
    }

    fmt.Println("\nKnapsack 2")
    items2, w2, v2 := knap(wants2, len(wants2)-1, 20)
    fmt.Println("Berat maksimum : 20")
    fmt.Println("Item yang dapat dibawa :", items2)
    fmt.Println("Berat total:", w2)
    fmt.Println("Value total:", v2)    
}
