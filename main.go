package main

import bMdl "github.com/vivalapanda/branch_analysis/branching_model"

func main() {
	model := bMdl.New(3)
	model.ParseFile("dataset/trace_1.txt.bds")
	model.ParseFile("dataset/trace_3.txt.bds")
	model.ParseFile("dataset/trace_4.txt.bds")
	model.ParseFile("dataset/trace_5.txt.bds")
	model.ParseFile("dataset/trace_6.txt.bds")
	model.ParseFile("dataset/trace_7.txt.bds")
	model.ParseFile("dataset/trace_8.txt.bds")
	model.ParseFile("dataset/trace_9.txt.bds")
	model.ParseFile("dataset/trace_10.txt.bds")
	model.ParseFile("dataset/trace_11.txt.bds")
	model.ParseFile("dataset/trace_12.txt.bds")
	model.ParseFile("dataset/trace_13.txt.bds")
	model.ParseFile("dataset/trace_14.txt.bds")
	model.ParseFile("dataset/trace_15.txt.bds")
	model.ParseFile("dataset/trace_16.txt.bds")
	model.ParseFile("dataset/trace_17.txt.bds")
	model.ParseFile("dataset/trace_18.txt.bds")
	model.ParseFile("dataset/trace_19.txt.bds")
	model.ParseFile("dataset/trace_20.txt.bds")
	model.DisplayStatistics()
}
