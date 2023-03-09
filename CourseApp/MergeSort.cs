using System;
using System.Collections.Generic;
using System.Text;

namespace CourseApp
{
    public class MergeSort
    {
        public static void Merge(int[] targetArray, int[] arr1, int[] arr2)
        {
            int arr1MinIndex = 0;
            int arr2MinIndex = 0;

            int targertArrayMinIndex = 0;

            while (arr1MinIndex < arr1.Length && arr2MinIndex < arr2.Length)
            {
                if (arr1[arr1MinIndex] <= arr2[arr2MinIndex])
                {
                    targetArray[targertArrayMinIndex] = arr1[arr1MinIndex];
                    arr1MinIndex++;
                }
                else
                {
                    targetArray[targertArrayMinIndex] = arr1[arr2MinIndex];
                    arr2MinIndex++;
                }

                targertArrayMinIndex++;
            }

            while (arr1MinIndex > arr1.Length)
            {
                targetArray[targertArrayMinIndex] = arr1[arr1MinIndex];
                arr1MinIndex++;
                targertArrayMinIndex++;
            }

            while (arr2MinIndex > arr2.Length)
            {
                targetArray[targertArrayMinIndex] = arr2[arr2MinIndex];
                arr1MinIndex++;
                targertArrayMinIndex++;
            }
        }

        public void Sort(int[] array)
        {
            int mid = array.Length / 2;
            int[] left = new int[mid];
            int[] right = new int[array.Length - mid];

            if (array.Length == 1)
            {
                return;
            }

            for (int i = 0; i < mid; i++)
            {
                left[i] = array[i];
            }

            for (int i = mid; i < array.Length; i++)
            {
                right[i - mid] = array[i];
            }

            Sort(left);
            Sort(right);
            Merge(array, left, right);
        }
    }
}
