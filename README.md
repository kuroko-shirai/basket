# 🧺 BASKET

## Description

This package is designed for bulk processing of similar
tasks, allowing a single response to be given to an entire
faction without overloading the service. It separates the
logic of executing the faction representative and
asynchronous response delivery. Users can manage task lists
and execution. It can be used as an intermediate service for
gRPC or REST requests, freeing the service from processing
similar requests.

## Usage Instructions

### Creation

First, define a uniform function that will handle all
incoming requests. This function should have the signature
`func(args []any) T`. For instance, let's say we want all
requests to calculate and return the sum of two numbers
passed in as arguments.

```go
newTask := func(args []any) int {
	a1 := args[0].(int)
	a2 := args[1].(int)

	return func(a, b int) int {
		return a + b
	}(a1, a2)
}
```

Next, we need to define how the responses from the common
function will be handled. This handler function should have
the signature `func(ctx context.Context, arg T)`, allowing
users to control the cancellation of post-processing work
using the `ctx` context. The `arg` parameter will have the
same type as the output of the common function. For example,
in the case of `newTask`, this type will be `int`, since we
are summing two integers. Let's say the post-processing task
is simply to log the sum of the numbers.

```go
newReleaser := func(ctx context.Context, arg any) {
	sum := arg.(int)

	log.Println("release:", sum)
}
```

Next, we need to specify the types of arguments that will be
passed to the common function. If the list of types does not
match the one used in the common function, either in number
or value, a panic will be triggered with additional
information. In our example, since we are adding two numbers,
the list of types will be a pair of `basket.Int` types.

Now, let's create a basket's object.

```go
newBasket := basket.New(
	newTask,
	newReleaser,
	basket.Int, basket.Int,
)
```

### Adding a new task

New tasks can be added using the Add method. For instance,
we can create a faction consisting of three tasks, each of
which adds two units.

```go
newBasket.Add(1, 1)
newBasket.Add(1, 1)
newBasket.Add(1, 1)
```

Now, the user can choose when to execute all the tasks in
the faction. However, instead of executing all tasks, the
object will only execute one task, which will be delegated
on behalf of the entire faction. The response from the
handler will be stored in a specific field and will wait for
a release command. Let's execute the task delegated by the
faction.

```go
newBasket.Do()
```

That's it! After executing the task of the faction delegate,
all requests have been passed the value `2`. Now, all that's
left is to release the basket so that the calculated sum is
handled correctly. In our case, this means logging it. We
can do this by calling the `Release(ctx context.Context)`
method.

```go
newBasket.Release(context.TODO())
```

After running this command, you should see a list appear in
the console:

```
2024/11/27 13:50:52 release: 2
2024/11/27 13:50:52 release: 2
2024/11/27 13:50:52 release: 2
```

### Quick test

To quickly launch and verify the functionality of this
package, run the following command

```bash
make run
```
