defmodule AoC do
  def read_file do
    {:ok, d} = File.read("./year2024/day2-in")
    d
    |> String.split("\n", trim: true)
    |> Enum.map(&line_to_ints/1)
  end

  def part_one(list) do
    list
    |> Enum.map(&is_valid?/1)
    |> Enum.sum
  end

  def part_two(list) do
    list
    |> Enum.map(&valid_chunked?/1)
    |> Enum.sum
  end

  defp line_to_ints(line) do
    line
    |> String.split
    |> Enum.map(&String.to_integer/1)
  end

  defp is_valid?(list) do
    grouped = Enum.zip(list, Enum.drop(list, 1))

    inc =
      grouped
      |> Enum.map(fn {a, b} ->
        a < b and Kernel.abs(b - a) <= 3 and Kernel.abs(b - a) >= 1
      end)

    dec =
      grouped
      |> Enum.map(fn {a, b} ->
        a > b and Kernel.abs(b - a) <= 3 and Kernel.abs(b - a) >= 1
      end)

    if Enum.all?(inc) or Enum.all?(dec) do
      1
    else
      0
    end
  end

  defp valid_chunked?(line) do
    Enum.map(0..length(line), fn idx ->
      Enum.take(line, idx) ++ Enum.drop(line, idx+1)
    end)
    |> Enum.map(&is_valid?/1)
    |> Enum.any?(fn x -> x == 1 end)
    |> if(do: 1, else: 0)
  end
end

input = AoC.read_file
IO.puts "Part 1: #{AoC.part_one(input)}"
IO.puts "Part 2: #{AoC.part_two(input)}"
