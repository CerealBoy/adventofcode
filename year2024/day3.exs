defmodule AoC do
  def read_file do
    {:ok, d} = File.read("./year2024/day3-in")
    d
  end

  def part_one(file) do
    ~r/mul\((\d+),(\d+)\)/
    |> Regex.scan(file, capture: :all_but_first)
    |> Enum.map(&mult/1)
    |> Enum.sum
  end

  def part_two(file) do
    ~r/(do)\(\)|(don)\'t\(\)|mul\((\d+),(\d+)\)/
    |> Regex.scan(file, capture: :all_but_first)
    |> Enum.map(fn y -> Enum.filter(y, fn x -> String.length(x) > 0 end) end) # Remove the unmatched portions
    |> switched(true, 0)
  end

  defp mult([left, right]), do: String.to_integer(left) * String.to_integer(right)

  defp switched([], _, acc), do: acc
  defp switched([["do"] | tail], _, acc), do: switched(tail, true, acc)
  defp switched([["don"] | tail], _, acc), do: switched(tail, false, acc)
  defp switched([_ | tail], false, acc), do: switched(tail, false, acc)
  defp switched([head | tail], true, acc), do: switched(tail, true, acc + mult(head))
end

input = AoC.read_file
IO.puts "Part 1: #{AoC.part_one(input)}"
IO.puts "Part 2: #{AoC.part_two(input)}"
