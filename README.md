# CPU Visualization Tool

This tool provides a real-time visualization of CPU cores and threads usage in the terminal. It is built using Go and utilizes the `shirou/gopsutil` library to fetch CPU information. The display is refreshed dynamically to show updated CPU usage.

## Features

- Displays a graphical representation of CPU cores and threads.
- Dynamically updates thread usage with colored bars:
  - **Green**: 0-25% usage.
  - **Yellow**: 26-50% usage.
  - **Orange**: 51-75% usage.
  - **Red**: 76-100% usage.
- Configurable refresh rate between 1 and 10 seconds.

## Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build -o cpu-visualizer
   ```

## Usage

### Run Once

Run the program to display the CPU state once:

```bash
./cpu-visualizer
```

### Watch Mode

Run the program in watch mode to refresh the display periodically:

```bash
./cpu-visualizer --watch
```

### Custom Refresh Rate

Specify a custom refresh rate (in seconds, between 1 and 10):

```bash
./cpu-visualizer --watch --refresh-rate=5
```

## Example Output

```bash
┌───────────────────────────────────────────────────────────┐
│        AMD Ryzen 5 3400G with Radeon Vega Graphics        │
│                          3700MHz                          │
│ ┌──────────────────────────┐ ┌──────────────────────────┐ │
│ │          Core 0          │ │          Core 1          │ │
│ │ ┌──────────────────────┐ │ │ ┌──────────────────────┐ │ │
│ │ │       Thread 0       │ │ │ │       Thread 2       │ │ │
│ │ │█                     │ │ │ │████████              │ │ │
│ │ └──────────────────────┘ │ │ └──────────────────────┘ │ │
│ │ ┌──────────────────────┐ │ │ ┌──────────────────────┐ │ │
│ │ │       Thread 1       │ │ │ │       Thread 3       │ │ │
│ │ │█                     │ │ │ │█                     │ │ │
│ │ └──────────────────────┘ │ │ └──────────────────────┘ │ │
│ └──────────────────────────┘ └──────────────────────────┘ │
│ ┌──────────────────────────┐ ┌──────────────────────────┐ │
│ │          Core 2          │ │          Core 3          │ │
│ │ ┌──────────────────────┐ │ │ ┌──────────────────────┐ │ │
│ │ │       Thread 4       │ │ │ │       Thread 6       │ │ │
│ │ │██████████████████████│ │ │ │█                     │ │ │
│ │ └──────────────────────┘ │ │ └──────────────────────┘ │ │
│ │ ┌──────────────────────┐ │ │ ┌──────────────────────┐ │ │
│ │ │       Thread 5       │ │ │ │       Thread 7       │ │ │
│ │ │█                     │ │ │ │█                     │ │ │
│ │ └──────────────────────┘ │ │ └──────────────────────┘ │ │
│ └──────────────────────────┘ └──────────────────────────┘ │
└───────────────────────────────────────────────────────────┘
```

## Requirements

- Go 1.16+

## License

This project is licensed under the MIT License. See the LICENSE file for details.
