# adventofcode

Solutions for the Advent of Code series

## build

    make build

To keep this simple, building is summed up with the single command. There's not much to it.

## execution

    adventofcode [-h] [-d] [year-20${YEAR_NUMBER}] [-t] [--day ${DAY_NUMBER}]

Help is included within the binary where comprehensive explanation can be found for all switches and actions.

## convention

The ``year-20${YEAR_NUMBER}`` maps to a directory here that represents a specific year.

The ``--day ${DAY_NUMBER}`` will map to the input file and the algorithm applied to that input. Combined with ``-t`` to toggle between the input or test data specifically. For example, if the day of **6** was selected, in the below command, the listed files would be used.

    $ adventofcode year-2020 --day 6
      year2020/day6-in
    $ adventofcode year-2019 --day 3 -t
      year2019/day3-test

