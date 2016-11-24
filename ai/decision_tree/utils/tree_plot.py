#!coding: utf-8
from matplotlib import pyplot as plt


class PLTNode(object):
    def __init__(self, tag, area, depth, parent_x, arrow_text=None):
        self.parent_x = parent_x
        self.depth = depth
        self.area = area
        self.tag = tag
        self.arrow_text = arrow_text


def show_tree(plot, tree):
    """ Draw the tree on the plot

    Examples:
        {
          "tag": "root",
          "children": {
              "yes": {"tag": "A"},
              "no": {
                  "tag": "B",
                  "children": {"yes": {"tag": "B1"}, "no": {"tag": "B2"}}
              }
          }
        }

            root
            /   \
           yes  no
          /      \
         A        B
                 / \
               yes  no
              /      \
            B1       B2


    Args:
        plot:
        tree(dict):

    Returns:

    """
    if not tree:
        return
    area = (0, 1)
    depth = 1
    nodes = []
    prt_x = (area[0] + area[1]) / 2.0
    scan_nodes(nodes, '', tree, area, depth, prt_x)
    nodes.sort(key=lambda x: x.depth, reverse=True)
    max_depth_node = max(nodes, key=lambda x: x.depth)
    max_depth = max_depth_node.depth
    for node in nodes:
        draw_node(plot, node, max_depth)


def draw_node(plot, node, max_depth):
    """Draw the node on a subplot
    Args:
        node(PLTNode):

    Returns:

    """
    # calculate x, y of the node text
    area_l, area_r = node.area
    x = float(area_l + area_r) / 2.0
    y_l = 1.0 - 1.0 / max_depth * node.depth
    y_h = 1.0 - 1.0 / max_depth * (node.depth - 1)
    y = (y_h + y_l) / 2.0

    # calculate the arrow start point x y, same with x y of parent's
    p_x = node.parent_x
    if node.depth > 1:
        p_y_l = 1.0 - 1.0 / max_depth * (node.depth - 1)
        p_y_h = 1.0 - 1.0 / max_depth * (node.depth - 2)
        p_y = (p_y_h + p_y_l) / 2.0
    else:
        p_y = 1.0

    sub = plot.subplot()
    # draw the arrow and node text
    sub.annotate(node.tag, (p_x, p_y), xytext=(x, y),
                 bbox=dict(boxstyle="circle", color='#FF6633'),
                 arrowprops=dict(arrowstyle="<-", color='g'),
                 va="center", ha="center")
    # draw the arrow_text if necessary
    if node.arrow_text is not None:
        arrow_mid_x, arrow_mid_y = (p_x + x) / 2.0, (p_y + y) / 2.0
        rotation = None
        sub.text(arrow_mid_x, arrow_mid_y, node.arrow_text,
                 va="center", ha="center", rotation=rotation)


def scan_nodes(nodes, arrow_text, tree, area, depth, parent_x):
    """Scan tree dict to create `PLTNode` nodes list

    Args:
        nodes(list):
        arrow_text(str):
        tree(dict):
        area(tuple):
        depth(int):
        parent_x(float):

    Returns:

    """
    tag = tree["tag"]
    children = tree.get("children") or dict()
    area_lft, area_rgt = area
    node = PLTNode(tag, area, depth, parent_x, arrow_text)
    nodes.append(node)
    if children:
        area_offset = float(area_rgt - area_lft) / (len(children))
        cur = area[0]
        prt_x = (area_rgt + area_lft) / 2.0
        for child_arrow, child in children.items():
            scan_nodes(nodes, child_arrow, child, (cur, cur + area_offset),
                       depth + 1, prt_x)
            cur += area_offset


if __name__ == '__main__':
    tree = {
        "tag": "root",
        "children": {
            "yes": {"tag": "A"},
            "no": {
                "tag": "B",
                "children": {
                    "yes": {
                        "tag": "B1",
                        "children": {"+1": {"tag": "B1+1"}, "-1": {"tag": "B1-"}}
                    },
                    "no": {"tag": "B2"}
                }
            }
        }
    }
    show_tree(plt, tree)
    plt.show()

