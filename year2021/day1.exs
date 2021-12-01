defmodule AoC do
  def read_file do
    {:ok, d} = File.read("./day1-in")
    d
    |> String.split
    |> Enum.map(&String.to_integer/1)
  end

  def first([], _, count), do: IO.puts "#1: " <> Integer.to_string(count)
  def first([head | tail], prev, count) do
    with :true <- head > prev do
      first(tail, head, count+1)
    end
    with :false <- head > prev do
      first(tail, head, count)
    end
  end

  def second([], _, _, _, count), do: IO.puts "#2: " <> Integer.to_string(count)
  def second([head | tail], first, second, prev, count) do
    with :true <- first+second+head > prev do
      second(tail, second, head, first+second+head, count+1)
    end
    with :false <- first+second+head > prev do
      second(tail, second, head, first+second+head, count)
    end
  end
end

[head | list] = AoC.read_file
AoC.first(list, head, 0)

[first | [second | [third | list]]] = AoC.read_file
AoC.second(list, second, third, first+second+third, 0)
