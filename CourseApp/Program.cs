using System;
using CourseApp.Module2;

namespace CourseApp
{
    public class Program
    {
        public static void Main(string[] args)
        {
            int[] numbers = { 1, 2, 3, 4, -213, 12, -1, 5, 12 };
            var sort = new QuickSort();
            sort.Sort(numbers, 0, numbers.Length - 1);
            foreach (int i in numbers)
            {
                Console.WriteLine(i);
            }
        }
    }
}
