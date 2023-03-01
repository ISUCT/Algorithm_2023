using System;
using CourseApp.Module2;

namespace CourseApp
{
    public class Program
    {
        public static void Main(string[] args)
        {
            // BubbleSort.BubbleSortMethod();
            Console.WriteLine(Plus(10, 12));
            int a = Convert.ToInt32(Console.ReadLine());
            Console.WriteLine(a.GetType());
        }

    public static int Plus(int a, int b)
        {
            /*int a = int.Parse(Console.ReadLine());
            int b = int.Parse(Console.ReadLine());*/

            return a + b;
        }
    }
}
