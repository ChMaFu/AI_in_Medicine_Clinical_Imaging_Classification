# EyeNet

## Overview

This is the code for [this](https://youtu.be/DCcmFXXAHf4) video on Youtube by Siraj Raval on medical classification and a fork of his original work [here](https://github.com/llSourcell/AI_in_Medicine_Clinical_Imaging_Classification).

The only addition made thus far is the addition of a Go-based image resizing tool that is able to use all cores of a given machine through a simple approach to splitting the list of images into N-groups, where N is the number of logical cores assigned to the host machine as determined by the Go native library function `runtime.NumCPU()`. (see [NumCPU()](https://golang.org/pkg/runtime/#NumCPU))

