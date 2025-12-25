# Unicode practice (in Go)

## Description

I was curious how runes in Go are working behind the scenes and turned out it's **UTF-8** characters, which I had to dive into.
All-in-all, it's pretty simple.
Current code is showing difference between Latin and Cyrillic symbols, how they are represented and how they can be perceived (as **UTF-8** or raw bytes).

### Usage

Clone the repo, then inside run `go run .` and you will be met with `> `. Input the words from the language you want and put `true` or `false` at the end to tell the program whether you want to read it as **UTF-8** or raw bytes.
Example:
```bash
к true
```

### Details

If the symbol needs more than 1 byte to be represented, then this symbol is being encoded with **UTF-8**.
Below is the **UTF-8** mask for 2 bytes character (Cyrillic):
```
110***** 10******
```

When `true` is selected (display as rune) then bytes are showing the Unicode after decoded with **UTF-8**, when `false` is selected then raw bytes are shown.
That is why the same Cyrillic symbol "к" can give different bytes in this program.
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

By adding `0000010000111010` (decoded as **UTF-8**) to the mask `110*****10******` we will have `1101000010111010` (raw bytes).

Addition:
```text
00000100 00111010
110***** 10******
11010000 10111010
```

