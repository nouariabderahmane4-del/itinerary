# Flight Itinerary Formatter (Go)

This Go program reads an **input text file** containing flight details with airport codes and timestamps, converts them into human-readable formats, and outputs a cleaned, formatted version.  

It uses a CSV file (`airport-lookup.csv`) to replace **IATA** and **ICAO** airport codes with full airport names and cities, and also formats ISO 8601 dates and times (12-hour and 24-hour formats).

---

## Features

- Replace airport codes (IATA `#LAX`, ICAO `##EGLL`) with full airport names.  
- Convert date and time formats:
  - `D(2025-04-05T12:30-02:00)` → `05 Apr 2025`  
  - `T12(2025-04-05T07:30-02:00)` → `07:30AM (-02:00)`  
  - `T24(2025-04-05T14:00Z)` → `14:00 (+00:00)`  
- Cleans up blank lines and normalizes text.  
- Converts vertical whitespace (`\r`, `\v`, `\f`) to newlines.  


---

## EXTRAS:

- Replace codes with corresponding cities (`*#LAX` → *Los Angeles*). 
- Optional colored output in the terminal.  
- Displays a farewell message with ASCII airplane animation. 
- The terminal output...  is "animated" with a slow, character-by-character effect


---

## File Structure

```
project-folder/
├─ main.go
├─ input.txt
├─ output.txt
├─ airport-lookup.csv
├─ Packages/
│  ├─ input_Package/
│  │  └─ input.go        # Handles argument checking, reading input text and CSV
│  └─ processor_Package/
│     └─ processor.go    # Handles airport code replacement, date/time formatting, and final output
```

---

## How It Works

1. **Argument Validation**  
   - `main.go` calls `input.Check_args(os.Args)` to ensure three arguments are passed: input file, output file, and CSV lookup file.  
   - Supports `-h` flag to show usage.

2. **Read Input and CSV**  
   - `input.Read_txt()` reads the full input file into a string.  
   - `input.Read_csv()` reads the CSV file and converts each row into an `Airport` struct for lookup.

3. **Process Input**  
   - `processor.Input_analyzing()`:
     - Replaces airport codes (`#IATA`, `##ICAO`) with full names  
     - Replaces city codes (`*#IATA`, `*##ICAO`) with city/municipality names  
     - Converts ISO 8601 date/time strings (`D()`, `T12()`, `T24()`) into human-readable formats  
     - Normalizes spaces and removes extra blank lines  
     - Produces a version for file output and a colored version for terminal display

4. **Write Output**  
   - `processor.Final_Output()`:
     - Writes formatted text to the output file  
     - Optionally prints colored output in the terminal  
     - Shows a farewell message with an ASCII airplane animation

---

## Example

### Input (`input.txt`)
```
1. D(2022-05-09T08:07Z)
2. T12(2069-04-24T19:18-02:00)
3. T24(2080-05-04T14:54Z)
```

### Output (`output.txt`)
```
1. 09 May 2022
2. 07:18PM (-02:00)
3. 14:54 (+00:00)
```

- Terminal version may include ANSI colors for airport names and times.

---

## How to Run

1. Open the terminal in your project folder.  
2. Run the program:

```bash
go run . ./input.txt ./output.txt ./airport-lookup.csv
```

3. Follow terminal prompts if asked to display colored output.  

---

## License

MIT License

Copyright (c) 2025 

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction...