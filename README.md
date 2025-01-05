# EnvList

EnvList is a simple Go package to list environment variables.

## Installation

```sh
go get github.com/divakarz/envlist

## Example Usage
 
package main

import "github.com/divakarz/envlist"

func main() {
    envlist.ListEnvVars("all")  // or "keys"
}
