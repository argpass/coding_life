#!/bin/bash
sys=$(uname -s)

if [ $sys == "Linux" ]; then
    echo "using Linux"
elif [ $sys = "Darwin" ]; then    # '=' and '==' is ok
    echo "using Mac"
else
    echo "Unkown sys"
fi


