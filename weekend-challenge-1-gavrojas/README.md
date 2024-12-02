[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/Cnet9Ax4)
# Weekend challenge 1

Each project should be a package of your module. When you finish you should upload it to your Github repo.

The program can be tested either with just a `go run` or ideally with an executable. So we do something like:

```shell
go run main.go evenodd
```

or

```shell
./w1challenge evenodd
```

The program will know which package to execute using a CLI argument:

- temperature
- bmi
- mario

## Temperature Converter

Create a function that converts a temperature from Celsius to Fahrenheit or vice versa, depending on a second argument that indicates the desired conversion.

```shell
$ ./w1challenge temperature
What's the temperature?
0
To which system?(C/F)
F
0 Celsius equals 32 Fahrenheit
```

## BMI Calculator

This program should asks the user for their weight and height and reply with they BMI and a message based on their underweight, normal weight or overweight.

In this example the user inputs `80` for their weight and `1.85` for their height.

```shell
$ ./w1challenge bmi
How much do you weigh? (don't lie)
80
How tall are you? (barefoot)
1.85
Right now your BMI is 23.37
You have a normal weight, I have healthy envy of you
```

If you don't know how to calculate the BMI you can click [here](http://letmegooglethat.com/?q=bmi+formula)

Use the following ranges for the final message:

If BMI is less than 18.5
-> 'You are underweight, add more potato to the broth'

If BMI is greater than or equal to 18.5 but less than 25
-> 'You have a normal weight, I have healthy envy of you'

If BMI is greater than or equal to 25
-> 'You are overweight, I know, the pandemic has affected us all'

Assume that the user will input valid numbers.

## Mario

> This exercise is borrowed from [CS50](https://cs50.harvard.edu/x/2020/), the Harvard introductory course to computer science 😁

The exercise is called Mario because toward the end of World 1-1 in Nintendo’s Super Mario Brothers, Mario must ascend right-aligned pyramid of blocks like this one:

![](https://cs50.harvard.edu/x/2020/psets/1/mario/less/pyramid.png)

This program should asks the user for `height` and print a right-aligned pyramid made of hashes (`#`) of that height.

Here some examples:

```
$ ./w1challenge mario
Pyramid height: 8
       #
      ##
     ###
    ####
   #####
  ######
 #######
########
```

```
$ ./w1challenge mario
Pyramid height: 4
   #
  ##
 ###
####
```

```
$ ./w1challenge mario
Pyramid height: 2
 #
##
```

```
$ ./w1challenge mario
Pyramid height: 1
#
```

The height should be between 1 and 8. If the user input is something different, the program should ask again and again until get a valid number:

```
$ ./w1challenge mario
Pyramid height: 0
Pyramid height: -5
Pyramid height: 5
    #
   ##
  ###
 ####
#####
```
