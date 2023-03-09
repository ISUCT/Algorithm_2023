using System;
using CourseApp.Module2;

namespace CourseApp
{
    public class Program
    {
        public static void Main(string[] args)
        {
            int[] numbers = { 1, 4, 7, 22, 5 };
            var mergeSort = new MergeSort();
            mergeSort.Sort(numbers);
            foreach (int i in numbers)
            {
                Console.WriteLine(i);
            }
        }
    }
}
