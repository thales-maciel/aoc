package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
  scanner := bufio.NewScanner(os.Stdin)

  top_three := []int{0,0,0}
  current_max := 0
  elf_sum := 0

  for scanner.Scan() {
    line := scanner.Text()

    if len(line) == 0 {
      if elf_sum > current_max {
        if elf_sum > top_three[0] {
          top_three[2] = top_three[1]
          top_three[1] = top_three[0]
          top_three[0] = elf_sum
        } else if elf_sum > top_three[1] {
          top_three[2] = top_three[1]
          top_three[1] = elf_sum
        } else {
          top_three[2] = elf_sum
        }
        current_max = top_three[2]
      }
      elf_sum = 0
      continue
    }

    num, _ := strconv.Atoi(line)
    elf_sum += num
  }

  if elf_sum > current_max {
    if elf_sum > top_three[0] {
      top_three[2] = top_three[1]
      top_three[1] = top_three[0]
      top_three[0] = elf_sum
    } else if elf_sum > top_three[1] {
      top_three[2] = top_three[1]
      top_three[1] = elf_sum
    } else {
      top_three[2] = elf_sum
    }
  }

  if scanner.Err() != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", scanner.Err())
  }

  fmt.Println(top_three)
  fmt.Println(top_three[0] + top_three[1] + top_three[2])

}
