# go-ght

Get Trending repositories on Github.

## Installation

    go get github.com/ramtiga/go-ght

## Usage

    repo := ght.GetRepoInf("go")

    for i, r := range repo {
            fmt.Printf("%d : %s\n", i + 1, r.Name)
    }
    

## License

MIT

## Author

ramtiga


