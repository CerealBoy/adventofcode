defmodule AoC do
  def read_file do
    {:ok, d} = File.read("./year2025/day1-in")

    d
    |> String.split("\n")
    |> make_list([])
  end

  def part_one([], _, acc), do: IO.puts("Part #1 #{acc}")

  def part_one(input, position, acc) when position < 0, do: part_one(input, position + 100, acc)

  def part_one(input, position, acc) when position > 99, do: part_one(input, position - 100, acc)

  def part_one([head | tail], 0, acc) do
    case head["direction"] do
      "L" -> part_one(tail, 0 - head["steps"], acc + 1)
      "R" -> part_one(tail, head["steps"], acc + 1)
    end
  end

  def part_one([head | tail], position, acc) do
    case head["direction"] do
      "L" -> part_one(tail, position - head["steps"], acc)
      "R" -> part_one(tail, position + head["steps"], acc)
    end
  end

  def part_two([], _, acc), do: IO.puts("Part #2 #{acc}")

  def part_two(input, position, acc) when position < 0,
    do: part_two(input, position + 100, acc + 1)

  def part_two(input, position, acc) when position > 99,
    do: part_two(input, position - 100, acc + 1)

  def part_two([head | tail], position, acc) do
    case head["direction"] do
      "L" -> part_two(tail, position - head["steps"], acc)
      "R" -> part_two(tail, position + head["steps"], acc)
    end
  end

  defp make_list([""], acc), do: acc

  defp make_list([head | tail], acc) do
    captures = Regex.named_captures(~r/(?<direction>[LR])(?<steps>\d+)/, head)

    make_list(
      tail,
      acc ++
        [%{"direction" => captures["direction"], "steps" => String.to_integer(captures["steps"])}]
    )
  end
end

list = AoC.read_file()
AoC.part_one(list, 50, 0)
AoC.part_two(list, 50, 0)
