# Monte Carlo.
## Calculate Pi 
Estimating the value of pi () by finding the area of a circle within a square. 
1. Draw a square and inscribe a circle within it.
2. Scatter a large number of random points uniformly across the square.
3. Count the number of points that fall inside the circle (e.g., by checking if the distance from the point to the center is less than the radius).
4. The ratio of points inside the circle to the total number of points is an estimate of the ratio of the circle's area to the square's area: PR^2 / 2R ^2  = P/4
5. Multiplying this ratio by 4 gives an estimate of P
 
## Calculate an Integral of any function
1. Draw a rectangle that fits a sergment of a functional graph in a given X interval
2. Generated random dots within it
3. Count how many falls inside (between the function graph and X axis)
4. The ratio of the count of inside points to total points multiplied by the area of the rectangle is the value of the integral
 
## Dependencies
Requires [Ebitengine](https://ebitengine.org/en/documents/install.html) for Go 2D graphics.

## Usage
```
./pi --help
Usage of ./pi:
  -c int
    	Number of samples (default 1000)
  -f string
    	[sqrt(x) x^2 1/(x+1)(sqrt(x)) x 10 e^x 1/x] (default "x")
  -i int
    	Number of samples increment (default 1000)
  -m string
    	Mode: pi, pi-text, func (default "pi")
  -n float
    	Min x of the integration range
  -x float
    	Max x of the integration range (default 10)
```

Calculate Pi text mode with 100,000,000 samples:
```
 ./pi -m pi-text -c 100000000
Starting Monte Carlo Pi estimation with 100000000 points...
Estimated Pi: 3.141819
Actual math.Pi: 3.141593
Error: 0.0072%
```
Calculate Pi in graphical mode starting with 10 samples and increasing by 10 samples with each iteration:
```
./pi -m pi -c 10 -i 10
```

Square root of x from 1 to 10:
```
./pi -c 100000000 -m func -n 1 -x 10 -f "sqrt(x)"
Starting Monte Carlo for function sqrt(x) with 100000000 points...
Estimate: 20.414715
```

Square of x from 1 to 10:
```
./pi -c 100000000 -m func -n 1 -x 10 -f x^2      
Starting Monte Carlo for function x^2 with 100000000 points...
Estimate: 332.945685
```
