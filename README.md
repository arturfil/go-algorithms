# LeetCode Practice Organizer

A Go-based tool to help you randomly select LeetCode problems from your curated list based on various filters like category or problem sets.

## Prerequisites

- Go 1.21 or higher
- A `data.json` file in the root directory containing your problem list

## Project Structure

```
leetcode/
├── data.json           # Problem database
├── main/              # Entry point
├── helpers/           # Core logic for problem selection
├── arrays/            # Solutions organized by category
├── strings/
├── graphs/
├── dynamic_programming/
└── ...                # Other problem categories
```

## Installation

No additional dependencies are required. Simply ensure you have Go installed:

```bash
go version
```

## Building the Project

Build the executable:

```bash
make build
```

Or manually:

```bash
go build -o exec ./main/main.go
```

## Running the Program

### Quick Start

Run with default settings (selects from all undone problems):

```bash
./exec
```

Or using make:

```bash
make run
```

### Using Filters

#### Filter by Category

Select a random problem from a specific category:

```bash
# Run directly
./exec -category=arrays
./exec -category=strings
./exec -category=graphs
./exec -category=dynamic_programming

# Using make
make run category=arrays
make run category=graphs
```

Available categories:
- `arrays`
- `strings`
- `binary-tree`
- `linked-lists`
- `graphs`
- `dynamic_programming`
- `intervals`
- `matrix`
- `heaps`
- `sort-search`
- `binary`
- `two_pointers`
- `sliding_window`
- And more...

#### Filter by Blind 75 List

Select only from the Blind 75 problem list:

```bash
# Run directly
./exec -from75

# Using make
make run from75=true
```

#### Combine Filters

Get a random Blind 75 problem from a specific category:

```bash
# Run directly
./exec -from75 -category=arrays

# Using make
make run from75=true category=arrays
```

## Available Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-from75` | boolean | `false` | Filter only problems from the Blind 75 list |
| `-category` | string | `"all"` | Filter by problem category |

## Example Usage

```bash
# Get any random undone problem
./exec

# Get a random array problem
./exec -category=arrays

# Get a random Blind 75 problem
./exec -from75

# Get a random Blind 75 string problem
./exec -from75 -category=strings

# Get a random dynamic programming problem
./exec -category=dynamic_programming
```

## Output

The program will display:
1. The number of problems matching your filters
2. The randomly selected problem title

Example output:
```
notDone slice length: 45
Problem chosen ---> Two Sum
```

If no problems match your filters:
```
notDone slice length: 0
No problems found matching the specified filters
```

## Managing Your Problem List

Problems are stored in `data.json` with the following structure:

```json
{
  "problems": [
    {
      "isFrom75": true,
      "algorithm": "hashmap",
      "category": "arrays",
      "done": false,
      "title": "Two Sum"
    }
  ]
}
```

To mark a problem as complete, change `"done": false` to `"done": true` in the JSON file.

## Makefile Commands

| Command | Description |
|---------|-------------|
| `make build` | Compile the project into an executable |
| `make run` | Run the executable with default settings |
| `make run category=<name>` | Run with specific category filter |
| `make run from75=true` | Run with Blind 75 filter |

## Tips

- Start your LeetCode practice session by running this tool to get a random problem
- Use category filters when you want to focus on a specific topic
- Use the `-from75` flag when preparing for interviews with common questions
- Keep your `data.json` updated by marking problems as done

## Contributing

When adding new solutions:
1. Place them in the appropriate category folder
2. Update `data.json` if adding new problems
3. Mark problems as done when completed

## License

Personal project for LeetCode practice organization.
