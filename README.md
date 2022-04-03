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

> Testable code is decoupled and single purpose.  
> "Learning Go with tests - os-exec"

> Always make sure your tests reflect how you'd like to use your code.  
> "Learning Go with tests - Error types"

> Small interfaces are good and are easily composed.  
> "Learning Go with tests - context-aware-reader"

> In software engineering, the delegation pattern is an object-oriented design pattern that allows object composition to achieve the same code reuse as inheritance.  
> [Delegation pattern](https://en.wikipedia.org/wiki/Delegation_pattern)

> If your tests are causing you pain, listen to that signal and think about the design of your code.  
> "Learning Go with tests - http-handlers-revisited"

### Meta section

#### Why section

> The promise of software is that it can change. This is why it is called soft ware, it is malleable compared to hardware.

> Go has made a number of choices which one hopes will make it more legacy-proof.
>
> - Go has only 25 keywords and a lot of systems can be built from the standard library and a few other small libraries.
> - The tooling in respect to testing, benchmarking, linting & shipping is first class compared to most alternatives.
> - The standard library is brilliant.
> - Very fast compilation speed for tight feedback loops
> - The Go backward compatibility promise. It looks like Go will get generics and other features in the future but the designers have promised that even Go code you wrote 5 years ago will still build.

> Any software system used in the real-world must change or become less and less useful in the environment.  
> "The Law of Continuous Change"

> As a system evolves, its complexity increases unless work is done to reduce it .
> "The Law of Increasing Complexity"

> When refactoring code you must not be changing behaviour.

> In order to safely refactor you need unit tests because they provide
>
> - Confidence you can reshape code without worrying about changing behaviour
> - Documentation for humans as to how the system should behave
> - Much faster and more reliable feedback than manual testing

> Favour testing behaviour rather than implementation detail.

> Refactoring Gives us signals about our unit tests. If we have to do manual checks, we need more tests. If tests are wrongly failing then our tests are at the wrong abstraction level (or have no value and should be deleted).
>
> - Helps us handle the complexities within and between our units.
>   Unit tests
>   - Give a safety net to refactor.
>   - Verify and document the behaviour of our units.

> (Well designed) units are
>
> - Easy to write meaningful unit tests.
> - Easy to refactor.

> Writing good unit tests is a design problem so think about structuring your code so you have meaningful units that you can integrate together like Lego bricks.

> TDD can help and force you to design well factored software iteratively, backed by tests to help future work as it arrives.
