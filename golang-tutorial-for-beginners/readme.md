```go

a := []int{1,2,3} // slice
a := [...]int{1,2,3} // array

// slice
a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// copy a
// NOT reference
b = a
// But slice a like array[:] (which means sliced)
// all data will be refercens
// that is both in slice and in array 
b := a[:]


append function is posssible only slice! NOT array.
because array dont need to that function, just insert data into index where you want 
```