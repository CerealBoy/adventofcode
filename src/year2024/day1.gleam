import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import simplifile as file

pub fn main() {
  let lists = gather_input("./year2024/day1-in")

  use inp <- result.map(lists)

  let #(left, right) = list.unzip(inp)
  let left_sorted = list.sort(left, by: int.compare)
  let right_sorted = list.sort(right, by: int.compare)

  let p1 = list.zip(left_sorted, right_sorted)
    |> list.map(fn(x) { x.1 - x.0 })
    |> list.map(int.absolute_value)
    |> int.sum

  io.println("Part 1: " <> int.to_string(p1))

  let finders =
    list.unique(left)
    |> list.fold(from: dict.new(), with: fn(count, value) {
      dict.insert(
        count,
        value,
        list.count(right, where: fn(x) { x == value }),
      )
    })

  let p2 = list.fold(over: left, from: 0, with: fn(total, value) {
    let count =
      dict.get(finders, value)
      |> result.unwrap(0)
    total + { value * count }
  })

  io.println("Part 2: " <> int.to_string(p2))
}

fn gather_input(path) -> Result(List(#(Int, Int)), String) {
  let in_file =
    file.read(from: path)
    |> result.map(string.trim_end)
    |> result.replace_error("Could not read file: " <> path)
  use contents <- result.try(in_file)

  string.split(contents, on: "\n")
  |> list.map(make)
  |> result.all
}

fn make(line: String) -> Result(#(Int, Int), String) {
  let res =
    string.split_once(line, "   ")
    |> result.replace_error("Could not split: " <> line)

  use #(left, right) <- result.try(res)

  let left_res =
    int.parse(left)
    |> result.replace_error("Could not parse int: " <> left)
  let right_res =
    int.parse(right)
    |> result.replace_error("Could not parse int: " <> right)

  use left_num <- result.try(left_res)
  use right_num <- result.try(right_res)
  Ok(#(left_num, right_num))
}
