import argparse
import sys


def first_sim_pt1(s1: list[str], s2: list[str]) -> str | None:
    m: dict[str, bool] = {}
    for ks1 in s1:
        m[ks1] = True
    for ks2 in s2:
        if m.get(ks2):
            return ks2

    return None

def first_sim_pt2(s1: list[str], s2: list[str], s3: list[str]) -> str | None:
    m: dict[str, int] = {}
    for ks1 in s1:
        m[ks1] = 1

    for ks2 in s2:
        if m.get(ks2) == 1:
            m[ks2] += 1
    for ks3 in s3:
        if m.get(ks3) == 2:
            return ks3

    return None

def calc_score(s: str) -> int | None:
    c = s[0]
    char_code = ord(c)

    if 65 <= char_code <= 90:
        return char_code - 38

    if 97 <= char_code <= 122:
        return char_code - 96

    return None

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-f",
                        help="Relative or absolute path to input file",
                        default="./input.txt")
    args = parser.parse_args()

    item_pri_total, badge_pri_total = 0, 0
    group: list[list[str]] = []

    with open(args.f, encoding="utf-8") as f:
        COUNT = 1
        while True:
            line = f.readline()
            if not line:
                break
            sack = [*line]
            half = len(sack) // 2
            c1 = sack[:half]
            c2 = sack[half:]

            # Part 1
            fsi = first_sim_pt1(c1, c2)
            if fsi is None:
                print("failed to find match between compartments")
                sys.exit(1)
            score = calc_score(fsi)
            if score is None:
                print(f"could not parse char from {fsi}")
                sys.exit(1)
            item_pri_total += score

            # Part 2
            group.append(sack)
            if len(group) == 3:
                fsb = first_sim_pt2(group[0], group[1], group[2])
                if fsb is None:
                    print("failed to find group badge")
                    sys.exit(1)
                score = calc_score(fsb)
                if score is None:
                    print(f"could not parse char from {fsb}")
                    sys.exit(1)
                badge_pri_total += score
                group = []

            COUNT += 1

    # Part 1
    print(f"Total item priorities: {item_pri_total}")

    # Part 2
    print(f"Total badge priorities: {badge_pri_total}")

