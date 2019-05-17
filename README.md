[![CircleCI](https://circleci.com/gh/pitakill/swiss-knife.svg?style=svg)](https://circleci.com/gh/pitakill/swiss-knife)
[![Build Status](https://cloud.drone.io/api/badges/pitakill/swiss-knife/status.svg)](https://cloud.drone.io/pitakill/swiss-knife)

# Rational

Make the learning curve from js to Go smoother, in the way achieving this
develop a `generic` tool that handles the slice in similar way that js handles
arrays

# Examples

```js
// In js the prototype Array is very powerful, so is desireable to handle slices
// in golang in a similar way. But golang does not have generics. So this
// pretends cover that gap

const animals = ['dog', 'cat', 'rabbit']

// Make plural of the elements of `animals` and store it in a new array
// With prototype we can achieve this with `map` but currently we are going to
// use forEach

const animalsInPlural = []

animals.forEach((animal, i) => animalsInPlural[i] = animal + 's')

// now animalsInPlural is ['dogs', 'cats', 'rabbits']

// Another example is:

const numbers = [1, 2, 3, 5, 8]

numbers.forEach(number => console.log(number * number))

// This outputs:
// 1
// 4
// 9
// 25
// 64
```

```go
// In Go with the lack of generics we need to implement at least two types of
// functions for each one of the examples in js, something like

func stringForEachWithIndex(input []string, iterator func(e string, i int)) {
  for i, e := range input {
    iterator(e, i)
  }
}

animals := []string{"dog", "cat", "rabbit"}

animalsInPlural := make([]string, len(animals))

stringForEachWithIndex(animals, func(e string, i int) {
  animalsInPlural[i] = e + "s"
})

// now animalsInPlural is ["dogs", "cats", "rabbits"]

func intForEach(input []int, iterator func(i int)) {
  for _, in := range input {
    iterator(in)
  }
}

numbers := []int{1, 2, 3, 5, 8}

intForEach(numbers, func(e int) {
  fmt.Println(e * e)
})
```

This snippets are equivalent in each language, but notice how we need to
implement at least two different functions for each slice of type with the
following firms in Go:

`func stringForEach(input []string, iterator func(e element, i int))`

and

`func stringForEach(input []string, iterator func(e element))`

for every builtin type in Go. So its a lot of work

With this library we can make the next thing:

```go
import "github.com/pitakill/swiss-knife/slice"

// ...

numbers := []int{1, 2, 3, 5, 8}

squareNumbers := make([]int, len(numbers))

slice.ForEach(numbers, func(e int, i int) {
  squareNumbers[i] = e * e
})

// now squareNumbers is [1, 4, 9, 25, 64]
// and we can make this

slice.ForEach(numbers, func(e int) {
  fmt.Println(e * e)
})

// and the outputs is:
// 1
// 4
// 9
// 25
// 64

// With a simple use of a `ForEach` that handles both cases of the iterator function
```
