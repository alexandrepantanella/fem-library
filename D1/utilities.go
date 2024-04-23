package D1

import (
	"fmt"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func VecToString(vec *mat.VecDense) string {
    var str strings.Builder
    size := vec.Len()
    str.WriteString("[")
    for i := 0; i < size; i++ {
        str.WriteString(fmt.Sprintf("%.8f", vec.AtVec(i)))
        if i != size-1 {
            str.WriteString(", ")
        }
    }
    str.WriteString("]")
    return str.String()
}