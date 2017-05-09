## Summary

This is a go program which parses a dataset of branch predictor traces in order
to build a Markov chain representing the likelyhood that a certain branch
is taken.

## Dataset

The dataset was taken from [Championship Branch Prediction](https://www.jilp.org/cbp/).

### Parsing

We took the branch traces provied as well as the framework for reading them.
We then modified the framework by adding

```C++
printf("%x(:%x\n", cbp_inst.taken, cbp_inst.instruction_addr);
```

at the bottom of the `get_branch_record` function in tread.cc of the framework.
This caused the framework to print out the address and whether the branch was
taken for every branch instruction. We piped this to file for each trace
producing the text files in the dataset directory of this repository.

### Format

Our files are named `trace_$n.txt.bds` and contain one row per branch instruction
encountered in the trace. The file is formatted as:
`wasBranchTaken:instuctionAddress`

### Output

Running this Go program will result in it printing each state and the probability
that the state led to a branch. Currently one has to edit main to change
how many bits of history are stored.
