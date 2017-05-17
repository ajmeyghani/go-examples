package closure

func IncMaker() func() int {
  i := 0
  return func() int {
    i = i + 1
    return i
  }
}
