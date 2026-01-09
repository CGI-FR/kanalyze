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

[MIT](https://choosealicense.com/licenses/mit/)
