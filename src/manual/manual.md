---
title: "IVS Calculator Manual"
subtitle: "v1.0.0"
author: "Tomáš Martykán, Filip Štolfa, Andrei Meleca, Nichita Gutu"
date: "29/04/2021"
titlepage: true
titlepage-color: "2BB9F1"
titlepage-text-color: "FFFFFF"
titlepage-rule-color: "FFFFFF"
footer-left: "Faculty of Information Technology, Brno University of Technology"
papersize: a4
geometry:
- margin=1in
---
IVS Calculator is an easy to use calculation tool with advanced features.

<!-- HIDEBUILTIN
## Install

IVS Calculator is provided as a Debian package. 
You can install this package by double clicking it in the file manager. You need to provide your password and your account needs to have sudo privileges. 

Or it can be installed using the following command: 

```
$ sudo dpkg -i ivs-calculator.deb
```

## Uninstall

To uninstall the application, you can use the Software Center. Go to the Installed section, find IVS Calculator and click uninstall.

Or in the terminal use the following command:

```
$ sudo dpkg -r ivs-calculator
```
HIDEBUILTIN -->

## Usage

The input is in standard mathematical form. To get a result, press <kbd>Enter</kbd> or click the <kbd>=</kbd> button. 
Calculations are done in mathematical order - multiplication and division are performed before addition and subtraction. 

The result of the calculation is persisted in the history for later. The history is cleared when the window is closed. 

The **C/CE** button operates in two ways. By clicking the button normally, it clears the last character. By clicking for a longer period, the whole input is cleared.

## Functions

* Addition
  * Example: 1+5 
* Subtraction
  * Example: 4-2
* Multiplication
  * Example: 5*3
* Division
  * Example: 12/4
* Modulo (division remainder)
  * Example: 5%3
* Abs
  * Example: |-4|
* Power of
  * Example: 6^2
  * Alternate syntax: 6p2
* Root
  * Example: 2√25
  * Alternate syntax: 2r25
* Factorial
  * Example: 4!

## About

© Copyright 2021 Tomáš Martykán, Filip Štolfa, Andrei Meleca, Nichita Gutu

The program is provided under the [GNU General Public License v3.0](https://github.com/martykan/ivs-calculator/blob/main/LICENSE)

Source code is available on [GitHub](https://github.com/martykan/ivs-calculator/)

Created as a part of IVS (Practical Aspects of Software Design) course at the Faculty of Information Technology at Brno University of Technology.
