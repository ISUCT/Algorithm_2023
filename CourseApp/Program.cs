using System;
using CourseApp.Module2;
using CourseApp.Module3;

namespace CourseApp
{
    public class Program
    {
        public static void Main()
        {
            var str = Console.ReadLine();
            var subStr = Console.ReadLine();
            var res = CycleOfStr.Find(str, subStr);
            Console.WriteLine(res);
        }
    }
}