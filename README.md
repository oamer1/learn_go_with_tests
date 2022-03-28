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

> When embedding types, really think about what impact that has on your public API.
> "Learning Go with tests - json"

> Breaking rules
>
> - Most rules in software engineering aren't really rules, just best practices that work 80% of the time.
> - We discovered a scenario where one of our previous "rules" of not testing internal functions was not helpful for us so we broke the rule.
> - It's important when breaking rules to understand the trade-off you are making. In our case, we were ok with it because it was just one test and would've been very difficult to exercise the scenario otherwise.
> - In order to be able to break the rules you must understand them first. An analogy is with learning guitar. It doesn't matter how creative you think you are, you must understand and practice the fundamentals.  
>   "Learning Go with tests - io"

> Whenever you have a lot of dependencies for a thing in your system, it implies it's doing too much.
> "Learning Go with tests - time"

> Testable code is decoupled and single purpose,
> "Learning Go with tests - os-exec"
