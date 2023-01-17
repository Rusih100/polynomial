# polynomial
A small library for working with polynomials. Supports big numbers

## Install
```
go get github.com/Rusih100/polynomial
```

## Description
The library supports the following operations on polynomials:
* Addition
* Subtraction
* Multiplication
* Division with remainder
* Taking coefficients modulo
* String representation
* String representation of a polynomial in the form of an array of its coefficients

## Creating a polynomial
Creating a polynomial x^3 + 5x - 4  
```go
polyCoefficient := []*big.Int{ // x^3 + 5x - 4
  big.NewInt(-4),
  big.NewInt(5),
  big.NewInt(0),
  big.NewInt(1),
}

poly := polynomial.NewPolynomial(polyCoefficient) 
```

## Example of addition

```go
polyCoefficientOne := []*big.Int{ // 10x + 5
  big.NewInt(5),
  big.NewInt(10),
}

polyCoefficientTwo := []*big.Int{ // 3x^2 + 2x + 1
  big.NewInt(1),
  big.NewInt(2),
  big.NewInt(3),
}

polyOne := polynomial.NewPolynomial(polyCoefficientOne)
polyTwo := polynomial.NewPolynomial(polyCoefficientTwo)

result := new(polynomial.Polynomial)
result = result.Add(polyOne, polyTwo) // 3x^2 + 12x + 6
```

## Documentation
https://pkg.go.dev/github.com/Rusih100/polynomial
