# LRU

## Description
	
	Simple LRU implementation of container/list

## Installation

	$ go get github.com/wayne666/LRU

## Usage

	l := LRU.New(capacity)

	l.Set(key, value) // Set key value

	l.Get(key) // Get value of key

	l.Purge() // Purge all cache
