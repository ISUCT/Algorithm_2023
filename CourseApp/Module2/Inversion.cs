using System;
using System.Collections.Generic;
using System.Text;

namespace CourseApp.Module2
{
    internal class Inversion
    {
        public static int[] Merge(int[] arr1, int[] arr2)
        {
            int i = 0, j = 0, k = 0;
            var resultArray = new int[arr1.Length + arr2.Length];
            while (k < resultArray.Length)
            {
                if ((j == arr2.Length) || (i < arr1.Length && arr1[i] <= arr2[j]))
                {
                    resultArray[k] = arr1[i];
                    i++;
                }
                else
                {
                    resultArray[k] = arr2[j];
                    j++;
                }

                k++;
            }

            return resultArray;
        }

        public static int[] Sort(int[] array)
        {
            var i = 0;
            if (array.Length == 1)
            {
                return array;
            }

            var mid = array.Length / 2;
            var left = new int[mid];
            var right = new int[array.Length - mid];

            Array.Copy(array, 0, left, 0, left.Length);
            Array.Copy(array, left.Length, right, 0, right.Length);

            left = Sort(left);
            right = Sort(right);
            Console.WriteLine(i);
            return Merge(left, right);
        }

        public static int[] Parse()
        {
            var massLen = Convert.ToInt16(Console.ReadLine());
            var str = Console.ReadLine().Split(" ");
            var mass = new int[massLen];
            for (int i = 0; i < mass.Length; i++)
            {
                mass[i] = Convert.ToInt32(str[i]);
            }

            return mass;
        }
    }
}
