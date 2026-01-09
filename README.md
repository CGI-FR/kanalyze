# K-Analyze

K-Analyze is a real-time CLI tool designed to visualize the frequency distribution of items in a data stream. It reads lines from standard input and displays a live histogram in your terminal, showing how many unique items have appeared $N$ times.

## Usage

Pipe data into `kanalyze`:

```bash
cat data.txt | ./kanalyze
```

### Example

Analyze the distribution of words in a file:

```bash
cat textfile.txt | tr -s ' ' '\n' | ./kanalyze
```

Analyze IP address frequency from an access log:

```bash
cat access.log | cut -d' ' -f1 | ./kanalyze
```

The output will look like this:

```text
  1 [  3] ███
  2 [  0]
  3 [  1] █
```

In this example:

- 3 items appeared exactly 1 time.
- 0 items appeared exactly 2 times.
- 1 item appeared exactly 3 times.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

Copyright (C) 2026 CGI France

kanalyze is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

kanalyze is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with kanalyze. If not, see http://www.gnu.org/licenses/.
