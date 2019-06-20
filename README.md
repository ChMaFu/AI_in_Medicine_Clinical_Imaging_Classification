# EyeNet

## Overview

This is the code for [this](https://youtu.be/DCcmFXXAHf4) video on Youtube by Siraj Raval on medical classification and a fork of his original work [here](https://github.com/llSourcell/AI_in_Medicine_Clinical_Imaging_Classification).

The only addition made thus far is the addition of a Go-based image resizing tool that is able to use all cores of a given host machine/VM/container to perform the TTA.

## Some Details

This is performed through a simple approach of splitting the list of images into N-groups, where N is the number of logical cores assigned to the host machine as determined by the Go native library function `runtime.NumCPU()`. (see [NumCPU()](https://golang.org/pkg/runtime/#NumCPU)). Once the image file list is broken up into 'N' groups, 'N' go routines are started to process the images.

The Go routines are all synchronized via a Go [waitgroup](https://golang.org/pkg/sync/#WaitGroup) so that all of the image processing work is completed before the process is terminated.

## Impetus

When attempting to run the original resizing code on the large set of images, I realized how inefficiently the code was using the AWS EC2 instance on which I was running; despite the EC2 instance having 8 cores and sufficinent memory/storage/etc., the Python image processing script was only using a single core, which meant a waste of money for anything more than a single core EC2 with enough memory to sustain the processing.

For the Go version, while not a perfectly linear scaling, the resizing code was able to fully utilize the 8x vCPU EC2 instance on which I was testing vs. only getting 1 vCPU of usage with the original Python implemention.

## Future

I would like to work with other folks in the data science/ML/AI community on similar efforts if anyone is interested.

