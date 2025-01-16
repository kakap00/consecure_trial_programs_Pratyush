module github.com/kakap00/package_trial

go 1.23.4

replace github.com/kakap00/custompackage v0.0.0 => ../custompackage
//simply telling go that instead of searching in github, we can replace with local location

require (
    github.com/kakap00/custompackage v0.0.0
)