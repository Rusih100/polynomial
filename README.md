# polynomial
A small library for working with polynomials. Supports big numbers

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
polyCoefficient := []*big.Int{
  big.NewInt(-4),
  big.NewInt(5),
  big.NewInt(0),
  big.NewInt(1),
}

poly := polynomial.NewPolynomial(polyCoefficient)
```
