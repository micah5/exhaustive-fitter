# exhaustive-fitter
Go library for finding the largest quadrilateral within another convex polygon.

## Installation
```
go get github.com/micah5/exhaustive-fitter
```

![circle](https://github.com/micah5/exhaustive-fitter/assets/40206415/bf48e34c-bae6-45a3-b4d1-cace64e49ecb)
![parallelogram](https://github.com/micah5/exhaustive-fitter/assets/40206415/c0786ee2-ae3e-4a4a-8da1-483d01a5e8a4)

## Usage
```go
// input is a flat array of x,y coordinates
outer := []float64{
  0, 0,
  0, 1,
  1, 1,
  1, 0,
} // This is the shape we want to fit to
inner := []float64{
  0.25, 0.25,
  0.25, 0.75,
  0.75, 0.75,
  0.75, 0.25,
} // This is the shape we want to fit (the "inner" shape)
result, err := fitter.Transform(inner, outer)
if err != nil {
  // do something
}
fitter.Plot("square.png", outer, result) // Helper function to plot the result
```
Make sure your polygon is convex otherwise an error will be returned.
