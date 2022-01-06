#!/bin/sh

runFormula() {
    cmd="rit math sum --number_1=25 --number_2=1"
    $cmd

    sum=".output.sum|tonumber"
    data=$(cat ~/.rit/output.json | jq $sum)
    echo $data

    result=$((data + 24))
    echo $result

    path=~/.rit/output-test.json
    touch $path
    # file=`jq {"result":$result} $path`
    # $file

    json={\"result\":$result}
    echo $json >$path

    test=$(cat $path)
    echo $test
}
