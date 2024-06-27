<a name="readme-top"></a>

<br />
<div align="center">
  <h1 align="center">1 billion rows challenge in Golang</h1>
</div>

## About

Learning golang and trying to learn more optimisation techniques by solving the 1 billion rows challenge. original challenge can be found [here](https://github.com/gunnarmorling/1brc)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

[![Go][golang]][golang-url]

## Getting Started

### Prerequisites

- Have go installed

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/NielsDobbelaar/1brc
   ```
2. get the data file by following the steps in the original challenge
3. run the script to run the program or time the program
   ```sh
   ./runSolution.sh
   ```
   or
   ```sh
    ./timeSolution.sh
   ```

## Timings

The timings of the different solutions I tried to solve the challenge as of yet.

### Solution 1: Baseline First Try

- **Execution Time:** 248 seconds (4 minutes 8 seconds)

### Solution 2: Improved parseFloat and Struct Pointer Lookup

- **Enhancements:**
  - Optimized `parseFloat` function.
  - Used pointers to structs for quick lookup.
- **Execution Time:** 116 seconds (1 minute 56 seconds)

### Solution 3: Parallel Processing with Goroutines

- **Enhancements:**
  - Introduced goroutines for parallel processing.
- **Performance:**
  - **4 Workers:** 315 seconds (5 minutes 15 seconds) (Failed to improve)
  - **8 Workers:** 646 seconds (10 minutes 46 seconds) (Failed to improve)

### Solution 4: Goroutines with Blocks of Work

- **Enhancements:**
  - Implemented goroutines with predefined blocks of work.
- **Performance:**
  - **10 Workers:** 75 seconds (1 minute 15 seconds)
  - **20 Workers:** 63 seconds (1 minute 3 seconds)
  - **200 Workers:** 46 seconds
  - **1000 Workers:** 50 seconds
  - **250 Workers:** 45 seconds

### Solution 5: Byte-Based Parsing and Error Handling Optimization

- **Enhancements:**
  - Replaced `strings.Cut` with byte-based parsing.
  - Removed error checking for performance gain.
- **Execution Time:** 45 seconds

### Final Optimization:

- **Additional Environment Optimization:**
  - Plugged in laptop charger.
  - Turned off low power mode.
- **Execution Time:** 37.660 seconds

## Todo

- [ ] remove scanner.Scan()
- [ ] move to integers instead of floats
- [ ] ? use a hashmap
- [ ] improve parallel processing
- [ ] speed up printing of solution
- [ ] tweak time script to run multiple times and get average
- [x] tweak time script to not time building the program

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Credits

- [Niels Dobbelaar](https://github.com/NielsDobbelaar)

[golang]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[golang-url]: https://go.dev/
