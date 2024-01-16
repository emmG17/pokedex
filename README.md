# PokeREPL - A Command-Line Pokedex in Go

PokeREPL is a lightweight and efficient Read-Eval-Print Loop (REPL) Pokedex written in Go, utilizing the PokeAPI (pokeapi.co) for fetching Pokemon information. The REPL comes with various commands to help you explore the Pokemon world right from your terminal.

## Features

- **help**: Displays a message with available commands.
- **exit**: Exits the program.
- **map**: Moves 20 locations forward in the Pokemon world.
- **mapb**: Moves 20 locations backward in the Pokemon world.
- **explore**: Allows you to explore a location.
- **catch**: Catch a wild Pokemon.
- **inspect**: Inspect the details of a Pokemon.
- **pokedex**: Displays your Pokedex, showcasing captured Pokemon.

## Installation

1. Ensure you have [Go](https://golang.org/doc/install) installed on your system.
2. Clone this repository:

    ```bash
    git clone git@github.com:emmG17/pokedex.git
    ```

3. Navigate to the project directory:

    ```bash
    cd pokedex
    ```

4. Build the REPL:

    ```bash
    go build
    ```

5. Run the REPL:

    ```bash
    ./pokedex
    ```

## Usage

1. Run the REPL using the steps mentioned in the installation section.
2. Use the available commands to interact with the Pokedex.

   ```bash
   pokedex> help
   ```

   This will display the list of commands and their descriptions.

## Caching

PokeREPL features caching to enhance performance and reduce unnecessary API requests. Each request is cached, preventing redundant calls to the PokeAPI.

## Acknowledgments

- Special thanks to [PokeAPI](https://pokeapi.co/) for providing the comprehensive Pokemon data.
- Inspired by the love for Pokemon and command-line interfaces.
- Special thanks to Boots (the cool bot from Boot.dev) for reminding me to finish this project

Feel free to contribute, report issues, or suggest enhancements!
