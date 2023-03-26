import unittest
from src.module3.string_period import string_period
from src.module3.cycle_shift import cycle_shift
from src.module3.find_substring import get_collision_index


class FindSubstringTest(unittest.TestCase):

    def test_main(self):
        res = get_collision_index("ababbababa", "aba")
        self.assertEqual([0, 5, 7], res)

    def test_zero(self):
        res = get_collision_index("aaaa", "b")
        self.assertEqual([], res)


class CycleShiftTest(unittest.TestCase):

    def test_main(self):
        res = cycle_shift("zabcd", "abcdz")
        self.assertEqual(4, res)

    def test_no_matches(self):
        res = cycle_shift("a", "b")
        self.assertEqual(-1, res)

    def test_zero(self):
        res = cycle_shift("a", "a")
        self.assertEqual(0, res)


class StringPeriodTest(unittest.TestCase):
    def test_main(self):
        res = string_period("abcabcabc")
        self.assertEqual(3, res)

    def test_no_matches(self):
        res = string_period("abc")
        self.assertEqual(-1, res)


if __name__ == '__main__':
    unittest.main()
