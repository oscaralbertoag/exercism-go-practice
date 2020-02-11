package raindrops

import (
  "strconv"
)


func Convert(n int) string {
  result := ""

  result += DetermineRaindropSound(n, 3, "Pling")
  result += DetermineRaindropSound(n, 5, "Plang")
  result += DetermineRaindropSound(n, 7, "Plong")

  if result == "" {
    result = strconv.Itoa(n)
  }

  return result
}

func DetermineRaindropSound(n int, factor int, sound string) string {
  raindropSound := ""

  if n % factor == 0 {
    raindropSound = sound
  }

  return raindropSound
}
