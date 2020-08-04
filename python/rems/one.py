#!/usr/bin/env python3
import copy
from typing import (
    List,
)


def possibleWords(dictionary: List[str], hand: List[str]) -> List[str]:
    letters = copy.copy(hand)
    possibleWords = []
    discardHand = []
    currWord = ''
    for word in dictionary:
        currWord = ''
        letters = copy.copy(hand)
        for letter in letters:
            if letter in word:
                currWord += letter
                discardHand.append(letter)
            elif letter not in word:
                discardHand.append(letter)

            if currWord == word:
                possibleWords += [word]
                break
    return possibleWords


def validWords(possibleWords, characters):
    for word in possibleWords


def testPossibleWords():
  assert possibleWords(["xyz", "xz", "zz"], ["x", "z", "k", "r"]) == ['xz']
  assert possibleWords(["xyz", "xz", "zz"], ["y", "z", "k", "x"]) == ['xyz', 'xz']


if __name__ == '__main__':
    testPossibleWords()
