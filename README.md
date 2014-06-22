# go-ght

Get Trending repositories on Github.

## Installation

    go get github.com/ramtiga/go-ght

## Usage

    repo := ght.GetRepoInf("go")

    for i, r := range repo {
		fmt.Printf("%3d  %-40s  %4s %4s\n", i+1, r.Name, r.Star, r.Fork)
    }
    

## License

MIT

## Author

ramtiga


