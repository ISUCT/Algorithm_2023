using System;
using System.Collections.Generic;
using System.Globalization;
using System.Text;

namespace CourseApp
{
    public class QuickSort
    {
        // private int[] numbers = { 4, -1, 3421, 1234, -32, 12 };
        public static int FindPivot(int[] numbers, int minIndex, int maxIndex) // поиск опроного члена
        {
            int pivot = minIndex - 1;
            int temp;
            for ( int i = minIndex; i < maxIndex; i++)
            {
                if (numbers[i] < numbers[maxIndex])
                {
                    pivot++;
                    temp = numbers[pivot];
                    numbers[pivot] = numbers[i];
                    numbers[i] = temp;
                }
            }

            pivot++;
            temp = numbers[pivot];
            numbers[pivot] = numbers[maxIndex];
            numbers[maxIndex] = temp;

            return pivot;
        }

        public int[] Sort(int[] numbers, int minINdex, int maxIndex)
        {
            if (minINdex >= maxIndex)
            {
                return numbers;
            }

            int pivot = FindPivot(numbers, minINdex, maxIndex);

            Sort(numbers, minINdex, pivot - 1);
            Sort(nmbers, pivot + 1, maxIndex);

            return numbers;
        } // сама сортировка
    }
}
