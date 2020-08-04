from typing import Dict, Set


AIRPORTS: Set[str] = {
    # name of airport
}
TICKETS: Dict[str, Dict[str, int]] = {
    # src: {dest: count}
}


def k_hops(start: str, destination: str, a: int) -> bool:
    """
    n airport
    m tickets:
    * origin
    * dest

    exactly k flights

    O(n^k) where n is number of airports
    """
    if a == 0:
        return start == destination

    dests = TICKETS[start]
    for dest, count in dests.items():
        if not count:
            continue

        dests[dest] -= 1
        if k_hops(dest, destination, a - 1):
            return True
        dests[dest] += 1

    return False
