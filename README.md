# Practice of [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

### Code along with short explanations inline comments

## Used commands

| Description                                               | Command                               |
| --------------------------------------------------------- | ------------------------------------- |
| Run tests recursively                                     | go test -v ./...                      |
| Run Benchmarks                                            | go test -bench=.                      |
| Run coverage                                              | go test -cover                        |
| Install cli to check unchecked errors                     | go get -u github.com/kisielk/errcheck |
| Run coverage                                              | go test -cover                        |
| Run errcheck inside directory to get unchecked line error | errcheck                              |
| Go race detector                                          | go test -race                         |

## Top tips

> Optimism is an occupational hazard of programming. Feedback is the treatment.  
> Kent Beck

> Be as close to delivering real consumer value as quickly as possible (often called a "happy path"). Once we have delivered a small amount of consumer value end-to-end, further iteration of the rest of the requirements is usually straightforward.  
> Learning Go with tests - Reading files
