defmodule AoC do
  def read_file do
    {:ok, d} = File.read("./year2025/day2-in")

    d
    |> String.split(",")
    |> make_list([])
  end

  def part_one(list) do
    result =
      list
      |> Enum.map(fn x -> {x, invalid_id?(x, 2)} end)
      |> Enum.reject(fn {_, x} -> x == "0" end)
      |> Enum.map(fn {x, _} -> String.to_integer(x) end)
      |> Enum.reduce(0, fn x, acc -> acc + x end)

    IO.puts("Part #1 #{result}")
  end

  def part_two(list) do
    result =
      list
      |> Enum.map(fn x -> {x, split_based_on_length(x, 1)} end)
      |> Enum.reject(fn {_, x} -> x == nil end)
      |> Enum.map(fn {x, _} -> String.to_integer(x) end)
      |> Enum.reduce(0, fn x, acc -> acc + x end)

    IO.puts("Part #2 #{result}")
  end

  defp split_based_on_length(id, pieces) do
    result =
      id
      |> String.split("")
      |> Enum.reject(fn x -> x == "" end)
      |> Enum.chunk_every(pieces)
      |> Enum.map(&Enum.join/1)
      |> Enum.frequencies()

    case length(Map.keys(result)) do
      1 ->
        case Map.fetch(result, List.first(Map.keys(result))) do
          {:ok, count} ->
            if count > 1 do
              id
            end
        end

      _ ->
        if pieces <= String.length(id) / 2 do
          split_based_on_length(id, pieces + 1)
        end
    end
  end

  defp invalid_id?(id, pieces) when length(id) == 1 or rem(length(id), pieces) == 1, do: "0"

  defp invalid_id?(id, pieces) do
    half = div(String.length(id), pieces)

    if half > 0 do
      first_half = String.slice(id, 0..(half - 1))

      if first_half <> first_half == id do
        id
      else
        "0"
      end
    else
      "0"
    end
  end

  defp make_list([], acc), do: acc

  defp make_list([head | tail], acc) do
    captures = Regex.named_captures(~r/(?<lower>\d+)-(?<upper>\d+)/, head)

    make_list(
      tail,
      acc ++
        expand_numbers(
          String.to_integer(captures["lower"]),
          String.to_integer(captures["upper"])
        )
    )
  end

  defp expand_numbers(head, tail) do
    head..tail
    |> Enum.to_list()
    |> Enum.map(&Integer.to_string/1)
  end
end

list = AoC.read_file()
AoC.part_one(list)
AoC.part_two(list)
