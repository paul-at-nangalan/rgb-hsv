# rgb-hsv
Just a wrapper round HSV-RGB library "github.com/crazy3lf/colorconv"

# Build
go install

# Binary image (Linux)
./bin/rgb-hsv

# Command line
```
rgb-hsv -r=32 -g=112 -b=231 --direction=tohsv
```
output:
h=215.879397, s=0.861000, v=0.906000

```
rgb-hsv -h=215.879397 -s=0.861000 -v=0.906000 --direction=torgb
```
output:
h=215.879397, s=0.861000, v=0.906000