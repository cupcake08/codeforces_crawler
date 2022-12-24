#!/bin/sh

clear_stuff() {
    rm -f myC
    rm -f *.txt 
    cd ..
    rm -I *.txt
}

# only for testing for now
go run crawler.go
cd test
if g++ test.cpp -o myC; then
    (./myC < ../input0.txt) > output0_test.txt
    if diff -q ../output0.txt output0_test.txt; then
        echo "Test 0 passed"
    else
        echo "Test 0 failed"
    fi

    (./myC < ../input1.txt) > output1_test.txt
    if diff -q ../output1.txt output1_test.txt; then
        echo "Test 1 passed"
    else
        echo "Test 1 failed"
    fi
    clear_stuff
fi
