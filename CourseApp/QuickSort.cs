using System;
using System.Collections.Generic;
using System.Globalization;
using System.Text;

namespace CourseApp
{
    public class QuickSort
    {
        // private int[] numbers = { 4, -1, 3421, 1234, -32, 12 }; - массив, чтобы скопировать его в Main() при вызове данной функции
        public static int FindPivot(int[] numbers, int minIndex, int maxIndex) // поиск опроного члена
        {
            int pivot = minIndex - 1;
            for ( int i = minIndex; i < maxIndex; i++)
            {
                if (numbers[i] < numbers[maxIndex])
                {
                    pivot++;
                    (numbers[pivot], numbers[i]) = (numbers[i], numbers[pivot]);
                }
            }

            pivot++;
            (numbers[pivot], numbers[maxIndex]) = (numbers[maxIndex], numbers[pivot]);

            return pivot;
        }

        public static int[] Sort(int[] numbers, int minIndex, int maxIndex) // сама сортировка (рекурсивно для подмассивов)
        {
            if (minIndex >= maxIndex)
            {
                return numbers; // если длина подмассива == 1 => выйти из цикла
            }

            int pivot = FindPivot(numbers, minIndex, maxIndex);

            Sort(numbers, minIndex, pivot - 1);
            Sort(numbers, pivot + 1, maxIndex);

            return numbers;
        }
    }
}
