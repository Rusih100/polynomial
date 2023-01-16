package polynomial

import (
	"math/big"
	"strconv"
)

// Реализация многочлена и операций над ним

type Polynomial struct {
	coefficients []*big.Int
}

// Set - Задает массив коэфицентов, начиная со свободного члена многочлена.
func (c *Polynomial) Set(coefficients []*big.Int) *Polynomial {

	c.coefficients = []*big.Int{}

	lenCoefficients := len(coefficients)

	if lenCoefficients == 0 {
		c.coefficients = append(c.coefficients, big.NewInt(0))
		return c
	}

	for i := 0; i < lenCoefficients; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < lenCoefficients; i++ {
		c.coefficients[i].Set(coefficients[i])
	}

	// Убираем лишние нули
	for c.coefficients[len(c.coefficients)-1].Sign() == 0 && len(c.coefficients) > 1 {
		c.coefficients = c.coefficients[:len(c.coefficients)-1]
	}

	return c
}

// Get - Получения коэфицента i-той степени
func (c *Polynomial) Get(i int) *big.Int {
	result := big.NewInt(0)
	return result.Set(c.coefficients[i])
}

// SetPolynomial - Задает массив коэфицентов, относительно другого полинома
func (c *Polynomial) SetPolynomial(p *Polynomial) *Polynomial {

	c.coefficients = []*big.Int{}

	lenCoefficients := len(p.coefficients)

	if lenCoefficients == 0 {
		c.coefficients = append(c.coefficients, big.NewInt(0))
		return c
	}

	for i := 0; i < lenCoefficients; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < lenCoefficients; i++ {
		c.coefficients[i].Set(p.coefficients[i])
	}

	// Убираем лишние нули
	for c.coefficients[len(c.coefficients)-1].Sign() == 0 && len(c.coefficients) > 1 {
		c.coefficients = c.coefficients[:len(c.coefficients)-1]
	}

	return c
}

// Add - Складывает два многочлена a и b, и записывает в c
func (c *Polynomial) Add(_a, _b *Polynomial) *Polynomial {

	a := new(Polynomial)
	b := new(Polynomial)

	aLen := len(_a.coefficients)
	bLen := len(_b.coefficients)

	a.SetPolynomial(_a)
	b.SetPolynomial(_b)

	c.coefficients = []*big.Int{}
	maxLen := 0

	if aLen > bLen {
		maxLen = aLen
	} else {
		maxLen = bLen
	}

	for i := 0; i < maxLen; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < aLen; i++ {
		c.coefficients[i].Add(c.coefficients[i], a.coefficients[i])
	}

	for i := 0; i < bLen; i++ {
		c.coefficients[i].Add(c.coefficients[i], b.coefficients[i])
	}

	// Убираем лишние нули
	for c.coefficients[len(c.coefficients)-1].Sign() == 0 && len(c.coefficients) > 1 {
		c.coefficients = c.coefficients[:len(c.coefficients)-1]
	}

	return c
}

// Sub - вычитает из многочлена a многочлен b, и записывает в c
func (c *Polynomial) Sub(_a, _b *Polynomial) *Polynomial {

	a := new(Polynomial)
	b := new(Polynomial)

	aLen := len(_a.coefficients)
	bLen := len(_b.coefficients)

	a.SetPolynomial(_a)
	b.SetPolynomial(_b)

	c.coefficients = []*big.Int{}
	maxLen := 0

	if aLen > bLen {
		maxLen = aLen
	} else {
		maxLen = bLen
	}

	for i := 0; i < maxLen; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < aLen; i++ {
		c.coefficients[i].Add(c.coefficients[i], a.coefficients[i])
	}

	for i := 0; i < bLen; i++ {
		c.coefficients[i].Sub(c.coefficients[i], b.coefficients[i])
	}

	// Убираем лишние нули
	for c.coefficients[len(c.coefficients)-1].Sign() == 0 && len(c.coefficients) > 1 {
		c.coefficients = c.coefficients[:len(c.coefficients)-1]
	}

	return c
}

// Mul - Умножает два многочлена a и b, и записывает в c
func (c *Polynomial) Mul(_a, _b *Polynomial) *Polynomial {

	a := new(Polynomial)
	b := new(Polynomial)

	aLen := len(_a.coefficients)
	bLen := len(_b.coefficients)

	a.SetPolynomial(_a)
	b.SetPolynomial(_b)

	c.coefficients = []*big.Int{}

	maxLen := aLen + bLen - 1

	for i := 0; i < maxLen; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			c.coefficients[i+j].Add(c.coefficients[i+j], new(big.Int).Mul(a.coefficients[i], b.coefficients[j]))
		}
	}

	// Убираем лишние нули
	for c.coefficients[len(c.coefficients)-1].Sign() == 0 && len(c.coefficients) > 1 {
		c.coefficients = c.coefficients[:len(c.coefficients)-1]
	}

	return c
}

// Mod - Берет модуль для каждого коэфицента многочлена a
func (c *Polynomial) Mod(_a *Polynomial, _mod *big.Int) *Polynomial {

	a := new(Polynomial)
	aLen := len(_a.coefficients)

	mod := new(big.Int)
	mod.Set(_mod)

	a.SetPolynomial(_a)

	// Взятие модуля

	c.coefficients = []*big.Int{}

	for i := 0; i < aLen; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := 0; i < aLen; i++ {
		c.coefficients[i].Set(new(big.Int).Mod(a.coefficients[i], mod))
	}

	// Убираем лишние нули
	for c.coefficients[len(c.coefficients)-1].Sign() == 0 && len(c.coefficients) > 1 {
		c.coefficients = c.coefficients[:len(c.coefficients)-1]
	}

	return c

}

// QuoRem - Деление с остатком многочлена a на многочлен b
func (c *Polynomial) QuoRem(a, b *Polynomial) (quo, rem *Polynomial) {

	A := new(Polynomial)
	B := new(Polynomial)

	aLen := len(a.coefficients)
	bLen := len(b.coefficients)

	A.SetPolynomial(a)
	B.SetPolynomial(b)

	if bLen == 1 && B.coefficients[0].Sign() == 0 {
		panic("Division by zero")
	}

	// Очистка исходного масива
	c.coefficients = []*big.Int{}

	// Проверка длин
	if aLen < bLen {
		return c, A
	}

	aLen = aLen - 1
	bLen = bLen - 1

	// Деление
	QLen := aLen - bLen + 1

	for i := 0; i < QLen; i++ {
		c.coefficients = append(c.coefficients, big.NewInt(0))
	}

	for i := aLen; i >= bLen; i-- {
		c.coefficients[i-bLen].Set(
			new(big.Int).Div(A.coefficients[i], B.coefficients[bLen]),
		)

		for j := bLen; j >= 0; j-- {
			A.coefficients[i-bLen+j].Set(
				new(big.Int).Sub(A.coefficients[i-bLen+j], new(big.Int).Mul(B.coefficients[j], c.coefficients[i-bLen])),
			)
		}
	}

	// Убираем лишние нули
	for c.coefficients[len(c.coefficients)-1].Sign() == 0 && len(c.coefficients) > 1 {
		c.coefficients = c.coefficients[:len(c.coefficients)-1]
	}

	// Убираем лишние нули
	for A.coefficients[len(A.coefficients)-1].Sign() == 0 && len(A.coefficients) > 1 {
		A.coefficients = A.coefficients[:len(A.coefficients)-1]
	}

	return c, A
}

// Представление полинома в виде строки вида x^3 + 2x - 1
func (c *Polynomial) String() string {

	lenCoefficients := len(c.coefficients)
	result := ""

	temp := big.NewInt(0)

	// Случай длины 1
	if lenCoefficients == 1 && c.coefficients[0].Sign() == 0 {
		return "0 "
	}

	// Общий случай
	for i := lenCoefficients - 1; i >= 0; i-- {

		temp.Set(c.coefficients[i])

		// Печатаем число если оно не ноль
		if temp.Sign() != 0 {

			// Смотрим на знак
			// Добавляем плюс, если число положительное и не самая большая степень
			if temp.Sign() > 0 && i != lenCoefficients-1 {
				result = result + "+ "
			}
			// Добавляем минус, если число отрицательное
			if temp.Sign() < 0 {
				result = result + "-"
				if i != lenCoefficients-1 {
					result = result + " "
				}

			}

			// Смотрим на на коэфицент по модулю
			temp = temp.Abs(temp)

			// Печатаем если коэфицент не равен 1
			if temp.Cmp(big.NewInt(1)) == 0 && i == 0 {
				result = result + temp.String()

			} else if temp.Cmp(big.NewInt(1)) != 0 {
				result = result + temp.String()
			}

			// Смотрим на степень и ставим x
			if i != 0 {
				result = result + "x"
			}

			// Печатем степень
			if i != 0 && i != 1 {
				result = result + "^" + strconv.Itoa(i) + " "
			} else {
				result = result + " "
			}

		}

	}

	return result
}

// StringCoefficients - Представление полиноса в виде строки вектора X(3 3 2)
func (c *Polynomial) StringCoefficients() string {

	lenCoefficients := len(c.coefficients)
	result := "P:("

	for i := lenCoefficients - 1; i >= 0; i-- {
		result = result + c.coefficients[i].String()
		if i != 0 {
			result = result + " "
		} else {
			result = result + ")"
		}
	}
	return result
}

// Value - Вычисляет значение многочлена при конкретном x
func (c *Polynomial) Value(_x *big.Int) *big.Int {

	// Копируем значения, чтобы не менять значения по указателю
	x := new(big.Int)
	x.Set(_x)

	result := big.NewInt(0)
	temp := new(big.Int)

	cLen := len(c.coefficients)

	for i := 0; i < cLen; i++ {
		temp = Pow(x, big.NewInt(int64(i)))
		temp = temp.Mul(temp, c.coefficients[i])
		result = result.Add(result, temp)
	}

	return result
}

// NewPolynomial - Создает полином и задает ему массив значений при создании
func NewPolynomial(coefficients []*big.Int) *Polynomial {
	return new(Polynomial).Set(coefficients)
}

// Pow - Алгоритм быстрого возведения в степень.
func Pow(_a *big.Int, _n *big.Int) (result *big.Int) {

	// Копируем значения, чтобы не менять значения по указателю
	a := new(big.Int)
	n := new(big.Int)

	a.Set(_a)
	n.Set(_n)

	// n < 0
	if n.Sign() == -1 {
		panic("n must be greater than or equal to zero")
	}

	result = big.NewInt(1)

	for i := 0; i < n.BitLen(); i++ {
		if n.Bit(i) == 1 {
			result = result.Mul(result, a)
		}
		a = a.Mul(a, a)
	}
	return result
}
