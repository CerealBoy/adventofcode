defmodule AoC do
  def read_file do
    {:ok, d} = File.read("./day2-in")
    d
    |> String.split("\n")
  end

  def first([], horizontal, depth), do: IO.puts "#1: " <> Integer.to_string(horizontal*depth)
  def first([head | tail], horizontal, depth) do
    case String.split(head) do
      ["forward", amount] -> first(tail, horizontal+String.to_integer(amount), depth)
      ["down", amount] -> first(tail, horizontal, depth+String.to_integer(amount))
      ["up", amount] -> first(tail, horizontal, depth-String.to_integer(amount))
      [] -> first([], horizontal, depth) # wtf
    end
  end

  def second([], horizontal, depth, _), do: IO.puts "#2: " <> Integer.to_string(horizontal*depth)
  def second([head | tail], horizontal, depth, aim) do
    case String.split(head) do
      ["forward", amount] -> second(tail, horizontal+String.to_integer(amount), depth+(aim*String.to_integer(amount)), aim)
      ["down", amount] -> second(tail, horizontal, depth, aim+String.to_integer(amount))
      ["up", amount] -> second(tail, horizontal, depth, aim-String.to_integer(amount))
      [] -> second(tail, horizontal, depth, aim)
    end
  end
end

list = AoC.read_file
AoC.first(list, 0, 0)
AoC.second(list, 0, 0, 0)
