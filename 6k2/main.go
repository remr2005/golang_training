package main

func FindUniq(arr []float32) float32 {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1] != arr[i] && arr[i] != arr[i-1] {
			return arr[i]
		}
	}
	if arr[0] != arr[1] {
		return arr[0]
	}
	return arr[len(arr)-1]
}

func main() {
	// fmt.Println(FindUniq([]float32{1, 1, 2}))
}
