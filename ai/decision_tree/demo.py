#!coding: utf-8
"""
This is a demo implement for <<统计学习方法>>.5.3 ID3 example
"""
import numpy as np
from decision_tree import DecisionTree


if __name__ == "__main__":
    dat = np.loadtxt("./dat.csv", dtype="i", delimiter=",",
                     usecols=range(1, 5))
    # AGE,WORK,HOUSE,PASS
    dt = DecisionTree(column_names=("age", "work", "house", "pass"))
    tree = dt.training(dat)
    tree_dict = tree.to_dict()["children"]["root"]
    from utils import tree_plot
    from matplotlib import pyplot as plt
    print tree_dict
    tree_plot.show_tree(plt, tree_dict)
    plt.show()
