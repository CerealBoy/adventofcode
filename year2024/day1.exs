defmodule AoC do
  def read_file do
    {:ok, d} = File.read("./year2024/day1-in")
    d
    |> String.split("\n")
    |> make_lists([], [])
  end

  def part_one([], [], diff), do: IO.puts "Part #1 #{diff}"
  def part_one([l | left], [r | right], diff) when l < r, do: part_one(left, right, diff + (r - l))
  def part_one([l | left], [r | right], diff) when l > r, do: part_one(left, right, diff + (l - r))
  def part_one([l | left], [r | right], diff) when l == r, do: part_one(left, right, diff)

  def part_two([], _, acc), do: IO.puts "Part #2 #{acc}"
  def part_two([h | tail], freq, acc) do
    with {:ok, x} <- Map.fetch(freq, h) do
      part_two(tail, freq, acc + (h * x))
    end
    with :error <- Map.fetch(freq, h) do
      part_two(tail, freq, acc)
    end
  end

  defp make_lists([], left, right), do: [Enum.sort(left), Enum.sort(right)]
  defp make_lists(["" | _], left, right), do: make_lists([], left, right)
  defp make_lists([head | tail], left, right) do
    [l, r] = String.split(head, "   ")
    make_lists(tail, Enum.concat(left, [String.to_integer(l)]), Enum.concat(right, [String.to_integer(r)]))
  end
end

[l, r] = AoC.read_file
AoC.part_one(l, r, 0)

freq = Enum.frequencies(r)
AoC.part_two(l, freq, 0)
