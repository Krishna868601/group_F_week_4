package main

// Function to reverse an array (or slice)
func reverseArray(arr []int) {
	for i, j := 0, len (arr)-1; i < j; i, j = i+1, j-1 { 
		arr[i], arr[j] = arr[j], arr[i] 
		} 
	}