//Package factorial shows the factorial of a given number
package fact

import (
    "math/big"
)

//Function factorial will give the output of number's facotrila
func factorial(x *big.Int) *big.Int {
    n := big.NewInt(1)
    if x.Cmp(big.NewInt(0)) == 0 {
        return n
    }
    return n.Mul(x, factorial(n.Sub(x, n)))
}

