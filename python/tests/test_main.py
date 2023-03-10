import unittest
from src import main


class SummTests(unittest.TestCase):

    def test_positive(self):
        res = main.get_collision_index("ababbababa", "aba")
        self.assertEqual([0, 5, 7], res)


if __name__ == '__main__':
    unittest.main()
