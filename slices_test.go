package slicesx

import "fmt"

func ExampleUnique() {
	fmt.Println(Unique([]string{"a", "b", "a", "c", "a", "c", "d", "d"}))
	// Output: [a b c d]
}

func ExampleFind() {
	fmt.Println(Find([]string{"foo", "bar", "baz"}, func(s string) bool {
		return s[0] == 'b'
	}))
	// Output: bar true
}

func ExampleTakeWhile() {
	fmt.Println(TakeWhile([]int{1, 3, 5, 2, 4}, func(x int) bool {
		return x%2 == 1
	}))
	// Output: [1 3 5]
}

func ExampleDropWhile() {
	fmt.Println(DropWhile([]int{1, 3, 5, 2, 4}, func(x int) bool {
		return x%2 == 1
	}))
	// Output: [2 4]
}

func ExampleGroupBy() {
	fmt.Println(GroupBy([]int{1, 3, 5, 2, 4}, func(x int) bool {
		return x%2 == 0
	}))
	// Output: map[false:[1 3 5] true:[2 4]]
}

func ExampleGroup() {
	fmt.Println(Group([]int{1, 2, 3, 4, 5}, 2))
	// Output: [[1 2] [3 4] [5]]
}

func ExamplePartition() {
	fmt.Println(Partition([]int{1, 2, 3, 4, 5}, func(x int) bool {
		return x%2 == 0
	}))
	// Output: [2 4] [1 3 5]
}
