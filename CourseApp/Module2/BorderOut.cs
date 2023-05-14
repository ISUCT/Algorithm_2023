using System;

namespace CourseApp.Module2
{
    public class BorderOut
    {
        public static int[] Merge(int[] arr1, int[] arr2)
        {
            var res = new int[arr1.Length + arr2.Length];
            int k = 0, i = 0, j = 0;
            while (k < res.Length)
            {
                if ((j == arr2.Length) || (i < arr1.Length && arr1[i] <= arr2[j]))
                {
                    res[k] = arr1[i];
                    i++;
                }
                else
                {
                    res[k] = arr2[j];
                    j++;
                }

               k++;
            }

           return res;
        }

        public static int[] MergeSort(int[] arr)
        {
            if (arr.Length == 1)
            {
                return arr;
            }

            var mid = arr.Length / 2;
            var left = new int[mid];
            var right = new int[arr.Length - mid];

            Array.Copy(arr, 0, left, 0, left.Length);
            Array.Copy(arr, left.Length, right, 0, right.Length);

            left = MergeSort(left);
            right = MergeSort(right);

            return Merge(left, right);
        }
    }
}