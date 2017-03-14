package utils

func AreInt32MatrixEqual(m1, m2 [][]int32)  bool {
    if len(m1) != len(m2) {
        return false
    }

    if len(m1) == 0 {
        return true
    }

    if len(m1[0]) != len(m2[0]) {
        return false
    }

    for i := 0; i< len(m1); i++ {
        for j := 0; j < len(m1[0]); j++ {
            if m1[i][j] != m2[i][j] {
                return false
            }
        }
    }

    return true
}
