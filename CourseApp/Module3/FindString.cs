using System;
using System.Linq;

namespace CourseApp.Module3
{
    public class FindString
    {
        public static void FindStr()
        {
            var a = Console.ReadLine();
            var b = Console.ReadLine();
            string result = null;
            for (int i = 0; i <= a.Length - b.Length + 1; i += b.Length)
            {
                result += a.IndexOf(b, i).ToString();
                result += " ";
            }

            Console.WriteLine(result.Split(" ").Aggregate((d, c) => d + " " + c) + "\r");
        }
    }
}
