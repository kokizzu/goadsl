package main
import "fmt"
func binarySearch(a []float64, x float64, l, r int) int { //O(log n)
    if r < l {		//kondisi jika data yang dicari tidak ditemukan (karena dalam rekursi r terus di increment maka saat r<l berarti data tidak ada)
        return -1	//mengembalikan -1
    }
    mid := (l + r) / 2 //mencari nilai tengah
    if a[mid] > x { //jika nilai dari posisi tengah array lebih besar dari nilai yang dicari 
        return binarySearch(a, x, l, mid-1) //maka akan dipanggil rekursi dengan perubahan parameter, sekarang rentang pencarian berkurang (berubah kearah kiri)
    } else if a[mid] < x { //jika nilai dari posisi tengah array lebih kecil dari nilai yang dicari 
        return binarySearch(a, x, mid+1, r) //maka akan dipanggil rekursi dengan perubahan parameter, sekarang rentang pencarian berkurang (berubah kearah kanan)
    }
    return mid //jika pencarian sudah selesai maka nilai mid yang dikembalikan sebagai posisi nilai yang dicari dalam array
}
func main(){
	list := []float64{1, 2, 3, 4, 5}
	var yang_dicari float64
	fmt.Println("nilai yang dicari?")
	fmt.Scan(&yang_dicari)
	posisi:=(binarySearch(list, yang_dicari, 0,4))
	//rentang pencarian dari 0 sampai n-1 (karena array dimulai dengan index 0)
	if posisi == -1 {
		fmt.Println("nilai yang anda cari tidak ditemukan")
	}else {
	fmt.Println("nilai yang anda cari ada di array elemen ke", posisi)
	}
}