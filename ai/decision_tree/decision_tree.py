#!coding: utf-8
from math import log
import numpy as np


def lg(v):
    return log(v, 2)


class Node(object):
    def __init__(self, tag):
        self.tag = tag
        self.children = dict()

    def is_leaf(self):
        return len(self.children) == 0

    def __setitem__(self, key, node):
        self.children[key] = node

    def __getitem__(self, key):
        return self.children[key]

    def to_dict(self):
        d = dict(tag=self.tag, children=dict())
        for k, c in self.children.items():
            d["children"][k] = c.to_dict()
        return d

    def __str__(self):
        return str(self.to_dict())


class DecisionTree(object):
    def __init__(self, column_names, data=None):
        self.column_names = column_names
        self.training_data = None
        self.tree = None
        if data is not None:
            self.init(data)

    def merge(self, data):
        """
        Args:
            data(np.ndarray):

        Returns:

        """
        rows = []
        if len(data.shape) == 1:
            rows.append(data)
        else:
            for r in data:
                rows.append(r)
        for r in self.training_data:
            rows.append(r)
        self.training_data = np.array(rows)

    def init(self, data):
        if len(data.shape) == 1:
            data = np.array([data])
        if len(data.shape) != 2:
            raise ValueError(u"invalid data %s", repr(data))
        if data.shape[1] != len(self.column_names):
            raise ValueError(u"expect %s cols , got %s",
                             len(self.column_names), data.shape[1])
        self.training_data = data

    def training(self, training_data):
        """
        Examples:
            A, B, Y
            0, 1, 1
            0, 1, 0
            0, 0, 1
            1, 1, 1

            Y is the classification, A and B are trait
        Args:
            training_data(np.ndarray):

        Returns:

        """
        if self.training_data is None:
            self.init(training_data)

        training_data = self.training_data
        root = Node(None)
        key = "root"
        r_idx = range(len(training_data))
        c_idx = range(len(training_data[0]))
        y_index = c_idx[-1]
        build_tree(root, key, training_data, c_idx, r_idx, y_index,
                   self.column_names)
        return root


def build_tree(tree, key, training_data, c_idx, r_idx, y_index, column_names):
    # start training
    row_cnt = len(r_idx)
    # col_cnt = len(training_data[0])
    V = [dict(index=i, ref_y=dict(), ref_self=dict(), info_gain=None)
         for i in c_idx]
    _v_y = None
    # scan rows to prepare basic info
    for r_i in r_idx:
        row = training_data[r_i]
        y_value = row[y_index]
        for _v in V:
            i = _v["index"]
            if i == y_index:
                _v_y = _v
            value = row[i]
            ref_self = _v["ref_self"]
            ref_y = _v["ref_y"]
            if value not in ref_self:
                ref_self[value] = 1
            else:
                ref_self[value] += 1
            if value not in ref_y:
                ref_y[value] = dict()
            if y_value not in ref_y[value]:
                ref_y[value][y_value] = 1
            else:
                ref_y[value][y_value] += 1
    if len(_v_y["ref_self"]) == 1:
        # all Y is the same value
        y = _v_y["ref_self"].keys()[0]
        tree[key] = Node(y)
        return tree
    # H(Y)
    h_y = sum([p * lg(p**(-1))
               for p in [float(n)/row_cnt
                         for y, n in _v_y["ref_self"].items()]])

    # H(Y|A), H(Y|B), ...;
    # g(Y, A), g(Y, B), ...
    t_list = [_v for _v in V if _v['index'] != y_index]
    for T in t_list:
        ref_self = T["ref_self"]
        ref_y = T["ref_y"]
        # H(Y|T)
        h_y_t = 0
        for ti, n in ref_self.items():
            y_cnt_map = ref_y[ti]
            # P{T=ti}
            p_ti = float(n)/row_cnt
            # H(Y|T=ti)
            h_ti = sum([p * lg(p**(-1))
                        for p in [float(y_n)/n
                                  for y, y_n in y_cnt_map.items()]])
            h_y_t += p_ti * h_ti
        # g(Y,T) = H(Y) - H(Y|T)
        g_y_t = h_y - h_y_t
        T['info_gain'] = g_y_t

    # choose the maximum info gain one
    choose_t = max(t_list, key=lambda x: x["info_gain"])
    idx = choose_t['index']
    tag = column_names[idx]
    tree[key] = Node(tag)
    # split the sample by the T's values, build child nodes
    keys = choose_t["ref_self"].keys()
    new_c_idx = [i for i in c_idx if i != idx]
    for k in keys:
        new_r_idx = [i for i in r_idx if training_data[i][idx] == k]
        build_tree(tree[key], k, training_data, new_c_idx, new_r_idx, y_index,
                   column_names)

