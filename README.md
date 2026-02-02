![Static Badge](https://img.shields.io/badge/Everybody_Codes-2024-blue) ![Static Badge](https://img.shields.io/badge/Go-1.25-blue)

# Everybody Codes

Solutions for [Everybody Codes](https://everybody.codes/events) challenges in Go.

## Progress - 2024

| **Quest** | **Progress** |
| --------- | ------------ |
| [Quest 1](https://everybody.codes/event/2024/quests/1) | Part 1 ⭐<br>Part 2 ⭐<br>Part 3 ⭐ |

## Usage
Within the year you are working with...

### Create a new quest
```bash
make new01
```

### Run a quest
```bash
make quest01
```

### Delete a quest
```bash
make delete01
```

### Help
```bash
make help
```

### Starting a new year
To set up for a new year (e.g., 2025):
```bash
mkdir -p 2025
cp 2024/Makefile 2025/
```

The Makefile automatically detects the year from the directory name, so no modifications needed!
It should also correctly build the year's directory structure as each new quest is added.

## Quest Structure

Each quest has 3 parts, with separate input files for each part:
- `part1()` - reads from `2024/inputs/quest01/part1.txt`
- `part2()` - reads from `2024/inputs/quest01/part2.txt`
- `part3()` - reads from `2024/inputs/quest01/part3.txt`

Solutions are automatically timed and results are displayed with execution time.
