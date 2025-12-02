# Calculate Pi using Monte Carlo.
Estimating the value of pi () by finding the area of a circle within a square. 
1. Draw a square and inscribe a circle within it.
2. Scatter a large number of random points uniformly across the square.
3. Count the number of points that fall inside the circle (e.g., by checking if the distance from the point to the center is less than the radius).
4. The ratio of points inside the circle to the total number of points is an estimate of the ratio of the circle's area to the square's area: PR^2 / 2R ^2  = P/4
5. Multiplying this ratio by 4 gives an estimate of P

## Dependencies
Requires [Ebitengine](https://ebitengine.org/en/documents/install.html) for Go 2D graphics.

## Usage
```
./pi <mode> <number of samples> [<number of samples increment>]
```
Run in text mode with 10,000,000 samples:
```
./pi t 10000000
```
Run in graphical mode starting with 10 samples and increasing by 10 samples with each iteration:
```
./pi g 10 10
```
