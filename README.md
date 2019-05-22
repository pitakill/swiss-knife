<p align="center">
  <h3 align="center">Swiss Knife</h3>
  <p align="center">An effort to bring JS Prototype to Go</p>
  <p align="center">
    <!--<a href="https://godoc.org/github.com/pitakill/swiss-knife">-->
      <!--<img src="https://godoc.org/github.com/pitakill/swiss-knife?status.svg">-->
    <!--</a>-->
    <a href="https://goreportcard.com/report/github.com/pitakill/swiss-knife">
      <img src="https://goreportcard.com/badge/github.com/pitakill/swiss-knife?v=1.0.0">
    </a>
    <a href="https://github.com/pitakill/swiss-knife/blob/master/LICENSE">
      <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
    </a>
<a href="https://app.fossa.io/projects/git%2Bgithub.com%2Fpitakill%2Fswiss-knife?ref=badge_shield" alt="FOSSA Status"><img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Fpitakill%2Fswiss-knife.svg?type=shield"/></a>
    <a href="https://circleci.com/gh/pitakill/swiss-knife">
      <img src="https://circleci.com/gh/pitakill/swiss-knife.svg?style=svg">
    </a>
    <a href="https://codecov.io/gh/pitakill/swiss-knife">
      <img src="https://codecov.io/gh/pitakill/swiss-knife/branch/master/graph/badge.svg" />
    </a>
    <a href="https://cloud.drone.io/pitakill/swiss-knife">
      <img src="https://cloud.drone.io/api/badges/pitakill/swiss-knife/status.svg">
    </a>
  </p>
</p>


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fpitakill%2Fswiss-knife.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fpitakill%2Fswiss-knife?ref=badge_large)

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