# go-stateful

This is a simple state machine that simulates the following workflow:

1. Initiate multiple test runs against a given merge commit
2. Wait for all of those test runs to be completed
3. Calculate the aggregate test results


This would be used to analyze the results of a single test suite against a single merge commit. If we wanted to search multiple merge commits, we would want to create multiple instances of this state machine and incorporate them into an even larger state machine to manage the search process.
