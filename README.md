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
> "Kent Beck"

> Make the test work quickly, committing whatever sins necessary in process.  
> "Kent Beck"

> Be as close to delivering real consumer value as quickly as possible (often called a "happy path"). Once we have delivered a small amount of consumer value end-to-end, further iteration of the rest of the requirements is usually straightforward.  
> "Learning Go with tests - Reading files"

> By taking TDD approach we
>
> - Keep the problem space small at any given time
> - Don't go down rabbit holes
> - If we ever get stuck/lost, doing a revert wouldn't lose loads of work.  
>   "Learning Go with tests - HTTP server"

> interface{} makes your code:
>
> - Less safe (mix apples and oranges), requires more error handling
> - Less expressive, interface{} tells you nothing about the data
> - More likely to rely on reflection, type-assertions etc which makes your code more difficult to work with and more error prone as it pushes checks from compile-time to runtime
>   "Learning Go with tests - Generics"

> Duplication is better than bad abstraction.
