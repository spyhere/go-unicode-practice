# Unicode practice (in Go)

## Description

I was curious how runes in Go are working behind the scenes and turned out it's **UTF-8** characters, which I had to dive into.
All-in-all, it's pretty simple.
Current code is showing difference between Latin and Cyrillic symbols, how they are represented and how they can be perceived (as **UTF-8** or raw bytes).

### Usage

Clone the repo, then inside run `go run .` and you will be met with `> `. Input the words from the language you want and put `true` or `false` at the end to tell the program whether you want to read it as **UTF-8** or raw bytes.

**Example:**
```bash
к true
```

Quitting the program
```bash
q
```

### Details

If the symbol needs more than 1 byte to be represented, then this symbol is being encoded with **UTF-8**.
Below is the **UTF-8** mask for 2 bytes character (Cyrillic):
```
110***** 10******
```

When `true` is selected (display as rune) then bytes are showing the Unicode without **UTF-8** mask, when `false` is selected then it's Unicode encoded with **UTF-8**.
That is why the same Cyrillic symbol "к" can give different bytes in this program depending whether you selected `true` or `false` for rune display.
Example:
```bash
> к (as rune: true)
|-----Данные-----|Длина|
        к           1   
 0000010000111010  16   

|-------к--------|
 0000010000111010 

> Ðº (as rune: false)
|-----Данные-----|Длина|
        к           1   
 1101000010111010  16   

|---Ð----|---º----|
 11010000 10111010 
```

By putting `0000010000111010` (binary representation of unicode code point) to the mask `110*****10******` we will have `1101000010111010`.

Putting bits to **UTF-8** mask:
1. **UTF-8** format is `110[5 bits payload] 10[6 bits payload]`.
2. Unicode code point - `00000100 00111010` (binary)
3. Drop leading zeros - `100 00111010`
4. Split it to align **UTF-8** payload - `10000 111010` (5 bits + 6 bits)
5. Add it to the mask - `11010000 10111010`

**Example:**
```text
1. 110***** 10****** - UTF-8 mask (* is a payload space)
2. 00000100 00111010 - unicode code point with 5 leading zeros
3.      100 00111010 - unicode code point without leading zeros
4.    10000   111010 - unicode code point aligned with UTF-8 payload (5 bits + 6 bits)
5. 11010000 10111010 - UTF-8 encoded Cyrillic symbol
```

