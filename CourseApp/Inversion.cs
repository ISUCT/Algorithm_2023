using System;
using System.Collections.Generic;
using System.Text;

namespace CourseApp
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

        public static int[] MergeSort(int[] resultArray)
        {
            if (resultArray.Length == 1)
            {
                return resultArray;
            }

            var mid = resultArray.Length / 2;
            var left = new int[mid];
            for (int i = 0; i < mid; i++)
            {
                left[i] = resultArray[i];
            }

            var rigth = new int[resultArray.Length - mid];
            for (int i = 0; i < rigth.Length; i++)
            {
                rigth[i] = resultArray[i + mid];
            }

            left = MergeSort(left);
            rigth = MergeSort(rigth);
            return Merge(left, rigth);
        }
    }
}
