from typing import List, Dict, Any


def dict_builder(input: List[str]) -> Dict[str, Any]:

    output = dict()

    for idx, v in enumerate(input):

        if v in output.keys():
            output[v].append(idx)
        else:
            output[v] = [idx]

    for k, v in output.items():
        if len(v) == 1:
            output[k] = v[0]

    return output


inputs = ["setName", "Name", "Name"]

o = dict_builder(inputs)
print(o)
