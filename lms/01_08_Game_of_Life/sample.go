package main

import (
  "os"
  "fmt"
  "math/rand"
)

type World struct {
  Height int
  Width int
  Cells [][]bool
}

func NewWorld(height, width int) *World {
  cells := make([][]bool height)
  for i := range cells {
     cells[i] = make([]bool width)
  }
  return &World{
    Height: height,
    Width: width,
    Cells: cells,
  }
}

func (w *World) Neighbours(x, y int) {

}

func (w *World) Next(x, y int) bool {
  n := w.Neightbours(x, y)
  alive := w.Cells[y][x]

  if n < 4 && n > 1 && alive {
    return true
  }
  if n == 3 && !alive {
    return true
  }
  return false
}

func NextState(oldWorld, newWorld *World) {
  for i := 0; i < oldWorld.Heigth; i++ {
    for j := 0; j < oldWorld.Width; i++ {
      newWorld[i][j] = oldWorld.Next(j, i)
    }
  }
}

func (w *World) Seed() {
  for _, row := range Cells {
    for i := range row {
      if rand.Intn(10) == 1 {
        row[i] = true
      }
    }
  }
}

func (w *World) SaveState(filename string) error {
  file, err := os.Create(filename)
  if err != nil {
    return err
  }
  defer file.Close()

  for i := 0; i < w.Height; i++ {
    for j := 0; j < w.Width; j++ {
      if w.Cells[i][j] {
        fmt.Fprint(file, 1)
      } else {
        fmt.Fprint(file, 0)
      }
    }
    if i < w.Height - 1 {
      fmt.Fprintln(file)
    }
  }

  return nil
}

func main() {
  height := 10
  width := 10
  currentWorld := NewWorld(height, width)
  nextWorld := NewWorld(height, width)

  currentWorld.Seed()
  for {
    fmt.Println(CurrentWorld)
    NextState(currentWorld, nextWorld)
    currentWorld = NextWorld
    time.Sleep(100 * time.Millisecond)
    fmt.Print("\033[H\033[2J")
  }
}
