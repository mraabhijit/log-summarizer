# Log Summarizer

A toy project to count the number of occurences of passes word strings in a file.
The tool uses byte chunks and streaming to solve memory problems and handles large files to the size of 1 GB and above gracefully.

## Usage

```bash
>>> go build
>>> log-summarizer <file-path> <word1> <word2> <word2> ... <wordN>
```

## Output

```bash
>>> <num-lines-in-file> <word1-count> <word2-count> <word3-count> ... <wordN-count> 
```
