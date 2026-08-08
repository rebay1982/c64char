# c64char
JPG/PNG to C64 character format converter.

This tool was build to quickly and easily put together custom character data for use in C64 applications/games. The
format of the character ROM/RAM on a C64 is divided in blocs of 8 bytes, where each bit represents an "on/off" state for
each pixel in the 8x8 character (8 pixels (bits) x 8 rows (bytes)).

## Encoding Input

The tool expects a JPG or PNG image with a width and height that are both divisible by 8 with a maximum of size of
320x200 pixels.

## Encoding Output

The tool outputs a format that is compatible with the [ACME](https://sourceforge.net/projects/acme-crossass/) assembler.
Why? that's simply what I use. The code is extensible to allow other formats for other assemblers to be produced when/if
the need comes up.

### Example

The following input image would produce the following output:

<img src="./assets/sample.png" width=100>

```
; CHAR 0
!byte %11111111
!byte %10000001
!byte %10111101
!byte %10100101
!byte %10100101
!byte %10111101
!byte %10000001
!byte %11111111

; CHAR 1
!byte %00000000
!byte %01111110
!byte %01100010
!byte %01010010
!byte %01001010
!byte %01000110
!byte %01111110
!byte %00000000

; CHAR 2
!byte %01111110
!byte %10000001
!byte %10111101
!byte %10001001
!byte %10010001
!byte %10111101
!byte %10000001
!byte %01111110

; CHAR 3
!byte %10000001
!byte %01111110
!byte %01000010
!byte %01000010
!byte %01000010
!byte %01000010
!byte %01111110
!byte %10000001

```
